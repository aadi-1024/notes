package handlers

import (
	"html/template"
	"net/http"

	"github.com/aadi-1024/notes/pkg/database"
	"github.com/aadi-1024/notes/pkg/models"
	"github.com/alexedwards/scs/v2"
)

func IndexPageHandler(t *template.Template, s *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		session := s.Get(r.Context(), "sessionData").(models.Session)
		t.ExecuteTemplate(w, "home.page.gohtml", map[string]any{
			"LoggedIn": session.LoggedIn,
			"Notes": session.Notes,
		})
	}
}

func LoginPageHandler(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		t.ExecuteTemplate(w, "login.page.gohtml", nil)
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

		user.Id = id
		session := models.Session{
			LoggedIn: true,
			User: &user,
			Notes: make([]*models.Note, 0),
		}
	
		s.Put(r.Context(), "sessionData", session)
		w.Header().Set("Content-Type", "text/html")
		w.Header().Add("HX-Redirect", "/")
		w.WriteHeader(http.StatusOK)
	}
}
