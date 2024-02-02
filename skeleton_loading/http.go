package skeleton_loading

import (
	"embed"
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
	"time"
)

//go:embed *.html
var fs embed.FS
var t *template.Template

func init() {
	t = template.Must(template.ParseFS(fs, "template.html"))
}

func RegisterRoutes(r chi.Router) {
	r.Get("/", HandleHomepage)
	r.Get("/stats/subscribers", HandleSubscribers)
	r.Get("/stats/average-click", HandleAverageClicks)
	r.Get("/stats/open-rate", HandleOpenRate)
}

func HandleHomepage(w http.ResponseWriter, r *http.Request) {
	t.Execute(w, nil)
}

type Stat struct {
	Label string
	Value string
}

func HandleSubscribers(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	s := Stat{
		Label: "Subscribers",
		Value: "10,897",
	}

	t.ExecuteTemplate(w, "stats", s)
}

func HandleAverageClicks(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	s := Stat{
		Label: "Avg. Click Rate",
		Value: "25.10%",
	}

	t.ExecuteTemplate(w, "stats", s)
}

func HandleOpenRate(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	s := Stat{
		Label: "Avg. Open Rate",
		Value: "32.17%",
	}

	t.ExecuteTemplate(w, "stats", s)
}
