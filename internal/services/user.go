package services

import (
	"clear-cut/internal/models"
	"sync"
)

var (
	users []models.User
	mu    sync.Mutex
)

// GetUsers returns all users
func GetUsers() []models.User {
	mu.Lock()
	defer mu.Unlock()
	return users
}

// CreateUser adds a new user to the slice
func CreateUser(user models.User) {
	mu.Lock()
	defer mu.Unlock()
	users = append(users, user)
}
