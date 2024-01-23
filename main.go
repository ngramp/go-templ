package main

import (
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"globolist/components"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	component := components.Home("World")
	r.Get("/", templ.Handler(component).ServeHTTP)
	http.ListenAndServe(":3000", r)
}
