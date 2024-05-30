package main

import (
	"html/template"

	"github.com/aadi-1024/notes/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func SetupRouter(mux *chi.Mux) {
	mux.Use(app.session.LoadAndSave)
	mux.Get("/login", handlers.LoginPageHandler(template.Must(template.New("login.page.gohtml").ParseFiles("./templates/login.page.gohtml", "./templates/base.layout.gohtml"))))
	mux.Post("/login", handlers.LoginPostHandler(app.db, app.session))
	mux.Get("/register", handlers.RegisterPageHandler(template.Must(template.New("register.page.gohtml").ParseFiles("./templates/register.page.gohtml", "./templates/base.layout.gohtml"))))
	mux.Post("/register", handlers.RegisterPostHandler(app.db))
	mux.Mount("/", mux.Group(func(r chi.Router) {
		r.Use(AuthMiddleware)
		r.Get("/", handlers.IndexPageHandler(template.Must(template.New("home.page.gohtml").ParseFiles("./templates/home.page.gohtml", "./templates/base.layout.gohtml")), app.session))
		r.Post("/logout", handlers.LogoutPostHandler(app.session))
	}))
}
