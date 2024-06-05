package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/aadi-1024/notes/pkg/database"
	"github.com/aadi-1024/notes/pkg/models"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/csrf"
)

func IndexPageHandler(t *template.Template, s *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		// session := s.Get(r.Context(), "sessionData").(models.Session)
		t.ExecuteTemplate(w, "home.page.gohtml", map[string]any{
			"LoggedIn":       true,
			csrf.TemplateTag: csrf.TemplateField(r),
		})
	}
}

func LoginPageHandler(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		t.ExecuteTemplate(w, "login.page.gohtml", map[string]any{
			csrf.TemplateTag: csrf.TemplateField(r),
		})
	}
}

func RegisterPageHandler(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Header().Add("Content-Type", "text/html")
		t.ExecuteTemplate(w, "register.page.gohtml", map[string]any{
			csrf.TemplateTag: csrf.TemplateField(r),
		})
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

func LogoutHandler(s *scs.SessionManager) http.HandlerFunc {
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
		t.ExecuteTemplate(w, "notes.partial.gohtml", map[string]any{
			"notes": notes,
		})
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

		w.Header().Add("HX-Refresh", "true")
		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte("successful"))
	}
}

func NoteDeleteHandler(d *database.Database, s *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userId := s.Get(r.Context(), "sessionData").(models.Session).User.Id
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = d.DeleteNote(id, userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("HX-Refresh", "true")
		w.WriteHeader(http.StatusOK)
	}
}

func NoteUpdateGetHandler(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		t.ExecuteTemplate(w, "update.partial.gohtml", map[string]any{
			"Id":             id,
			csrf.TemplateTag: csrf.TemplateField(r),
		})
	}
}

func NoteUpdatePutHandler(d *database.Database, s *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := s.Get(r.Context(), "sessionData").(models.Session).User.Id

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		note := models.Note{
			Id:     id,
			UserId: uid,
			Title:  r.FormValue("title"),
			Text:   r.FormValue("text"),
		}

		err = d.UpdateNote(note)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("HX-Refresh", "true")
		w.Write([]byte("successful"))
	}
}
