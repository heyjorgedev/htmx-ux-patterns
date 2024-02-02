package ecommerce_oob

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
var currentItems = 0

func init() {
	t = template.Must(template.ParseFS(fs, "template.html"))
}

func RegisterRoutes(r chi.Router) {
	r.Get("/", HandleHomepage)
	r.Post("/add-to-cart", HandleAddToCart)
}

type Page struct {
	CartItems int
}

func HandleHomepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, Page{CartItems: currentItems})
}

func HandleAddToCart(w http.ResponseWriter, r *http.Request) {
	// simulate adding to cart
	time.Sleep(1 * time.Second)

	currentItems = currentItems + 1

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Will confetti be displayed?
	w.Header().Set("HX-Trigger", "confetti")

	w.WriteHeader(http.StatusCreated)

	// The form doesn't need to be rendered if you have hx-swap="none"
	// but here I decided to render it so that you can see several html snippets being sent
	t.ExecuteTemplate(w, "form", nil)

	t.ExecuteTemplate(w, "cart", currentItems)
}
