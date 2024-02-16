package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/heyjorgedev/htmx-ux-patterns/ecommerce_oob"
	"github.com/heyjorgedev/htmx-ux-patterns/skeleton_loading"
	"github.com/heyjorgedev/htmx-ux-patterns/websockets"
	"net/http"
	"os"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/public/*", http.StripPrefix("/public", http.FileServer(http.Dir("public"))).ServeHTTP)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the home page"))
	})

	r.Route("/ecommerce-oob", ecommerce_oob.RegisterRoutes)
	r.Route("/skeleton", skeleton_loading.RegisterRoutes)
	r.Route("/websockets", websockets.RegisterRoutes)

	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = "127.0.0.1:8080"
	}

	http.ListenAndServe(listenAddr, r)
}
