package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/heyjorgedev/htmx-ux-patterns/skeleton_loading"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the home page"))
	})

	r.Route("/skeleton", skeleton_loading.RegisterRoutes)

	http.ListenAndServe(":8080", r)
}