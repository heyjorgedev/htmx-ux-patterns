package websockets

import (
	"bytes"
	"embed"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"golang.org/x/net/websocket"
	"html/template"
	"io"
	"net/http"
	"sync"
)

type VotingRequest struct {
	Vote string `json:"vote"`
}

type VotingState struct {
	FireEmojiVotes          int
	ExplodingHeadEmojiVotes int
	CoolEmojiVotes          int
}

//go:embed *.html
var fs embed.FS
var t *template.Template
var conns = make(map[*websocket.Conn]bool)
var m sync.Mutex
var votes = &VotingState{}

func init() {
	t = template.Must(template.ParseFS(fs, "template.html"))
	conns = map[*websocket.Conn]bool{}
}

func RegisterRoutes(r chi.Router) {
	r.Get("/", HandleHomepage)
	r.Handle("/feed", HandleWebsocket())
}

func HandleHomepage(w http.ResponseWriter, r *http.Request) {
	t.Execute(w, votes)
}

func HandleWebsocket() http.Handler {
	connectionsChanged := make(chan struct{})
	go func() {
		for {
			select {
			case <-connectionsChanged:
				m.Lock()
				buff := bytes.NewBuffer(nil)
				t.ExecuteTemplate(buff, "connected-clients", len(conns))
				for conn := range conns {
					conn.Write(buff.Bytes())
				}
				m.Unlock()
			}
		}
	}()
	return websocket.Handler(func(ws *websocket.Conn) {
		// conns[ws] is not concurrent-safe so we need to lock it before reading or writing
		m.Lock()
		conns[ws] = true
		// after the change we can unlock it
		m.Unlock()

		connectionsChanged <- struct{}{}

		// defer is used to ensure that the connection is removed from the map when the function returns
		defer func() {
			m.Lock()
			delete(conns, ws)
			m.Unlock()
			connectionsChanged <- struct{}{}
		}()

		for {
			b := make([]byte, 1024)
			n, err := ws.Read(b)
			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}
				continue
			}
			vote := &VotingRequest{}
			json.NewDecoder(bytes.NewReader(b[:n])).Decode(vote)

			// we create a buffer to store the html we are going to render in a few lines
			buff := bytes.NewBuffer(nil)

			m.Lock()
			switch vote.Vote {
			case "fire":
				votes.FireEmojiVotes++
				t.ExecuteTemplate(buff, "fire-emoji", votes.FireEmojiVotes)
			case "exploding":
				votes.ExplodingHeadEmojiVotes++
				t.ExecuteTemplate(buff, "exploding-head-emoji", votes.ExplodingHeadEmojiVotes)
			case "cool":
				votes.CoolEmojiVotes++
				t.ExecuteTemplate(buff, "cool-emoji", votes.CoolEmojiVotes)
			}
			for conn := range conns {
				// after rending the templates to the buffer we now broadcast the html response to the websocket connections
				_, _ = conn.Write(buff.Bytes())
			}
			m.Unlock()
		}
	})
}
