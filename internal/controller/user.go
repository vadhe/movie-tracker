package controller

import (
	"database/sql"
	"movie-tracker/internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	dbService := database.New()
	errt := r.ParseForm()
	if errt != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	// Query the database for the stored password hash
	var storedHash string
	err := dbService.DB().QueryRow("SELECT password_hash FROM users WHERE username = $1", username).Scan(&storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Compare the provided password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	sessionToken := uuid.New().String()
	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour), // Set cookie to expire in 24 hours
		HttpOnly: true,                           // Make the cookie accessible only to the server
		Secure:   false,                          // Set to true if serving over HTTPS
		Path:     "/",                            // Cookie will be sent for all paths in the domain
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
