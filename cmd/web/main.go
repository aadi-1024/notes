package main

import (
	"log"
	"net/http"
	"time"

	"github.com/aadi-1024/notes/pkg/database"
	"github.com/go-chi/chi/v5"
)

var app = &Config{}

func main() {
	mux := chi.NewMux()

	app.debug = true
	db, err := database.InitDatabase("postgres://postgres:password@localhost:5432/notes", 2*time.Second)
	if err != nil {
		log.Fatalln(err)
	}

	app.db = db

	SetupRouter(mux)
	srv := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	if err = srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
