package handlers

import (
	"clear-cut/internal/auth"
	"clear-cut/internal/models"
	"clear-cut/internal/storage"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// RegisterHandler handles user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Insert user into the database
	db := storage.GetDB()
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// LoginHandler handles user login and token generation
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.User

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Retrieve the user from the database
	var storedUser models.User
	db := storage.GetDB()
	err := db.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", creds.Email).Scan(&storedUser.ID, &storedUser.Name, &storedUser.Email, &storedUser.Password)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := auth.GenerateJWT(storedUser.Email)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Set the token in the cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true, // Set HttpOnly to prevent JavaScript access
		SameSite: http.SameSiteStrictMode, // Add SameSite attribute for CSRF protection
	})

	w.WriteHeader(http.StatusOK)
}

// ProfileHandler returns the user's profile
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
		return
	}

	// Validate JWT token
	claims, err := auth.ValidateJWT(token.Value)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Retrieve user profile from the database
	var user models.User
	db := storage.GetDB()
	err = db.QueryRow("SELECT id, name, email FROM users WHERE email = ?", claims.Email).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Send user profile response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
