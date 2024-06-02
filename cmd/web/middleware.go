package main

import (
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.session.Exists(r.Context(), "sessionData") && !strings.HasPrefix(r.URL.String(), "/login") && !strings.HasPrefix(r.URL.String(), "/register") {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
