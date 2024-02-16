package websockets

import (
	"embed"
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
)

//go:embed *.html
var fs embed.FS
var t *template.Template

func init() {
	t = template.Must(template.ParseFS(fs, "template.html"))
}

func RegisterRoutes(r chi.Router) {
	r.Get("/", HandleHomepage)
}

func HandleHomepage(w http.ResponseWriter, r *http.Request) {
	t.Execute(w, nil)
}
