package main

import (
	"html/template"

	"github.com/aadi-1024/notes/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func SetupRouter(mux *chi.Mux) {
	files := []string{"./templates/home.page.gohtml", "./templates/base.layout.gohtml"}
	t := template.Must(template.ParseFiles(files...))
	mux.Get("/", handlers.IndexPageHandler(t))
}
