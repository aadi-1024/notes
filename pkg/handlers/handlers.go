package handlers

import (
	"html/template"
	"net/http"

	"github.com/aadi-1024/notes/pkg/database"
	"github.com/aadi-1024/notes/pkg/models"
	"github.com/alexedwards/scs/v2"
)

func IndexPageHandler(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		t.ExecuteTemplate(w, "home.page.gohtml", map[string]any{
			"LoggedIn": true,
		})
	}
}

func LoginPostHandler(db *database.Database, s *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		username := r.PostFormValue("username")
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")

		user := models.User{
			Username: username,
			Email:    email,
			Password: password,
		}

		id, err := db.LoginUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if err = s.RenewToken(r.Context()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s.Put(r.Context(), "userId", id)
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte("successful"))
	}
}
