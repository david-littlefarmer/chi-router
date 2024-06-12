package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// r.Get("/first/second", GetHandler)

	r.Route("/first", func(r chi.Router) {
		r.Get("/second", RouterHandler)
	})

	r.Get("/first/second", GetHandler)

	http.ListenAndServe(":3333", r)
}

func RouterHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("router\n"))
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get\n"))
}
