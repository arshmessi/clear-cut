package models

// User represents a user in the system
type User struct {
	ID       int    `json:"id"`       // Changed to int to match the database schema
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"` // Exclude from JSON responses (omit empty)
}
