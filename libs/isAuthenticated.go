package libs

import "net/http"

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for a session or token here. For simplicity, we'll use a cookie.
		cookie, err := r.Cookie("session_token")
		if err != nil || cookie.Value == "" {
			if r.URL.Path == "/login-view" {
				next.ServeHTTP(w, r)
				return
			}
			http.Redirect(w, r, "/login-view", http.StatusSeeOther)
			return
		}
		if r.URL.Path == "/login-view" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		// If authenticated, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
