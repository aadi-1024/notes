package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aadi-1024/notes/pkg/database"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	_ "net/http/pprof"
)

var app = &Config{}

func main() {
	mux := chi.NewMux()
	s := scs.New()
	s.Lifetime = 30 * time.Minute
	s.Cookie.Name = "sessionCookie"
	app.session = s
	
	app.debug = true
	db, err := database.InitDatabase("postgres://postgres:password@localhost:5432/notes", 2*time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	
	app.db = db

	if app.debug {
		go func() {
			fmt.Println(http.ListenAndServe("localhost:6060", nil))
		}()		
	}

	validator := validator.New(validator.WithRequiredStructEnabled())
	app.validator = validator

	SetupRouter(mux)
	srv := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	if err = srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
