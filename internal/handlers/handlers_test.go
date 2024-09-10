package handlers

import (
	"bytes"
	"clear-cut/internal/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterHandler(t *testing.T) {
	// Initialize in-memory database
	storage.InitDB(true)

	// Define the test request
	userJSON := `{"name":"Test User","email":"test@example.com","password":"password123"}`
	req := httptest.NewRequest("POST", "/register", bytes.NewBufferString(userJSON))
	rec := httptest.NewRecorder()

	// Call the handler
	RegisterHandler(rec, req)

	// Check the response
	res := rec.Result()
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, res.StatusCode)
	}
}
