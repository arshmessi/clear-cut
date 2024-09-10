package handlers

import (
	"clear-cut/internal/auth"
	"clear-cut/internal/models"
	"clear-cut/internal/storage"
	"encoding/json"
	"net/http"
)

// CreateGroupHandler creates a new group
func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	var group models.Group

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Insert group into the database
	db := storage.GetDB()
	result, err := db.Exec("INSERT INTO groups (name, description) VALUES (?, ?)", group.Name, group.Description)
	if err != nil {
		http.Error(w, "Failed to create group", http.StatusInternalServerError)
		return
	}

	// Retrieve the last inserted ID
	groupID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve group ID", http.StatusInternalServerError)
		return
	}
	group.ID = int(groupID)

	// Send the response
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(group); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// GetGroupsHandler returns all groups for the authenticated user
func GetGroupsHandler(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
		return
	}

	claims, err := auth.ValidateJWT(token.Value)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	var groups []models.Group
	db := storage.GetDB()
	rows, err := db.Query("SELECT id, name, description FROM groups WHERE user_email = ?", claims.Email)
	if err != nil {
		http.Error(w, "Failed to fetch groups", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Read groups from the rows
	for rows.Next() {
		var group models.Group
		if err := rows.Scan(&group.ID, &group.Name, &group.Description); err != nil {
			http.Error(w, "Failed to scan group", http.StatusInternalServerError)
			return
		}
		groups = append(groups, group)
	}

	// Handle any error encountered while iterating over rows
	if err := rows.Err(); err != nil {
		http.Error(w, "Failed to fetch groups", http.StatusInternalServerError)
		return
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(groups); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
