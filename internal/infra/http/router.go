package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(ErrorMiddleware)

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	r.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	})

	return r
}
