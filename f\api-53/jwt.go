package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"go-training/handler"
)

func main() {
	r := chi.NewRouter()

	// define endpoints and associate with handler
	r.Get("/CSV", handler.CSVHandler)
	r.Get("/auth/signin", handler.SigninHandler)
	r.Get("/guest", handler.GuestHandler)

	// setup http server on port 8080
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}