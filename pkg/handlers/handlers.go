package handlers

import (
	"html/template"
	"net/http"

	"github.com/aadi-1024/notes/pkg/database"
	"github.com/aadi-1024/notes/pkg/models"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/validator/v10"
)

func IndexPageHandler(t *template.Template, s *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		// session := s.Get(r.Context(), "sessionData").(models.Session)
		t.ExecuteTemplate(w, "home.page.gohtml", map[string]any{
			"LoggedIn": true,
		})
	}
}

func LoginPageHandler(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		t.ExecuteTemplate(w, "login.page.gohtml", nil)
	}
}

func RegisterPageHandler(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		t.ExecuteTemplate(w, "register.page.gohtml", nil)
	}
}

func RegisterPostHandler(db *database.Database, v *validator.Validate) http.HandlerFunc {
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

		if err := v.Struct(user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = db.RegisterUser(user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("HX-Redirect", "/login")
		w.Header().Add("Content-Type", "text/html")
		w.WriteHeader(http.StatusSeeOther)
	}
}

func LoginPostHandler(db *database.Database, s *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")

		user := models.User{
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
			User: &user,
		}

		s.Put(r.Context(), "sessionData", session)
		w.Header().Set("Content-Type", "text/html")
		w.Header().Add("HX-Redirect", "/")
		w.WriteHeader(http.StatusOK)
	}
}

func LogoutPostHandler(s *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := s.RenewToken(r.Context()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := s.Destroy(r.Context()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "text/html")
		w.Header().Add("HX-Redirect", "/login")
		w.WriteHeader(http.StatusOK)
	}
}

func GetAllNotes(db *database.Database, s *scs.SessionManager, t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userid := s.Get(r.Context(), "sessionData").(models.Session).User.Id

		notes, err := db.GetAll(userid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "text/html")
		t.ExecuteTemplate(w, "notes.partial.gohtml", notes)
	}
}

func NotePostHandler(db *database.Database, s *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Get(r.Context(), "sessionData").(models.Session)

		if err := r.ParseForm(); err != nil {
			w.Header().Add("Content-Type", "text/plain")
			w.Write([]byte(err.Error()))
			return
		}

		title := r.PostForm.Get("title")
		text := r.PostForm.Get("text")

		_, err := db.CreateNote(models.Note{
			UserId: session.User.Id,
			Title:  title,
			Text:   text,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte("successful"))
	}
}
