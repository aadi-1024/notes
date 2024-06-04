package main

import (
	"html/template"

	"github.com/aadi-1024/notes/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
)

func SetupRouter(mux *chi.Mux) {
	csrf := csrf.Protect([]byte("A_HUGE_SECRET"), csrf.Path("/"), csrf.Secure(false))
	mux.Use(app.session.LoadAndSave)
	mux.Use(csrf)
	mux.Use(AuthMiddleware)
	mux.Get("/login", handlers.LoginPageHandler(template.Must(template.New("login.page.gohtml").ParseFiles("/home/aaditya/Projects/notes/templates/login.page.gohtml", "/home/aaditya/Projects/notes/templates/base.layout.gohtml"))))
	mux.Post("/login", handlers.LoginPostHandler(app.db, app.session))
	mux.Get("/register", handlers.RegisterPageHandler(template.Must(template.New("register.page.gohtml").ParseFiles("/home/aaditya/Projects/notes/templates/register.page.gohtml", "/home/aaditya/Projects/notes/templates/base.layout.gohtml"))))
	mux.Post("/register", handlers.RegisterPostHandler(app.db, app.validator))
	mux.Get("/notes", handlers.GetAllNotes(app.db, app.session, template.Must(template.New("notes.partial.gohtml").ParseFiles("/home/aaditya/Projects/notes/templates/notes.partial.gohtml"))))
	mux.Post("/notes", handlers.NotePostHandler(app.db, app.session))
	mux.Put("/notes", nil)
	mux.Get("/delete/{id}", handlers.NoteDeleteHandler(app.db, app.session))
	mux.Get("/", handlers.IndexPageHandler(template.Must(template.New("home.page.gohtml").ParseFiles("/home/aaditya/Projects/notes/templates/home.page.gohtml", "/home/aaditya/Projects/notes/templates/base.layout.gohtml")), app.session))
	mux.Get("/logout", handlers.LogoutHandler(app.session))
}
