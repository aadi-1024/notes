package main

import (
	"html/template"

	"github.com/aadi-1024/notes/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
)

func SetupRouter(mux *chi.Mux) {
	mux.Use(csrf.Protect([]byte("VERY_SECRET_KEY"), csrf.Path("/"), csrf.Secure(false)))
	mux.Use(app.session.LoadAndSave)
	mux.Use(AuthMiddleware)
	mux.Get("/login", handlers.LoginPageHandler(template.Must(template.New("login.page.gohtml").ParseFiles("/home/aaditya/Projects/notes/templates/login.page.gohtml", "/home/aaditya/Projects/notes/templates/base.layout.gohtml"))))
	mux.Post("/login/new", handlers.LoginPostHandler(app.db, app.session))
	mux.Get("/register", handlers.RegisterPageHandler(template.Must(template.New("register.page.gohtml").ParseFiles("/home/aaditya/Projects/notes/templates/register.page.gohtml", "/home/aaditya/Projects/notes/templates/base.layout.gohtml"))))
	mux.Post("/register/new", handlers.RegisterPostHandler(app.db, app.validator))
	mux.Get("/notes", handlers.GetAllNotes(app.db, app.session, template.Must(template.New("notes.partial.gohtml").ParseFiles("/home/aaditya/Projects/notes/templates/notes.partial.gohtml"))))
	mux.Post("/notes/new", handlers.NotePostHandler(app.db, app.session))
	mux.Delete("/delete/{id}", handlers.NoteDeleteHandler(app.db, app.session))
	mux.Get("/", handlers.IndexPageHandler(template.Must(template.New("home.page.gohtml").ParseFiles("/home/aaditya/Projects/notes/templates/home.page.gohtml", "/home/aaditya/Projects/notes/templates/base.layout.gohtml")), app.session))
	mux.Get("/logout", handlers.LogoutHandler(app.session))
}
