package handlers

import (
	"html/template"
	"net/http"
)

func IndexPageHandler(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		t.ExecuteTemplate(w, "home.page.gohtml", map[string]any{
			"LoggedIn": true,
		})
	}
}
