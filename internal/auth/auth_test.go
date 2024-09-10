package auth

import (
	"clear-cut/internal/storage"
	"os"
	"testing"
)

func TestGenerateAndValidateJWT(t *testing.T) {
	email := "test@example.com"

	// Generate token
	token, err := GenerateJWT(email)
	if err != nil {
		t.Fatalf("Failed to generate JWT: %v", err)
	}

	// Validate token
	claims, err := ValidateJWT(token)
	if err != nil {
		t.Fatalf("Failed to validate JWT: %v", err)
	}

	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}
}

func TestMain(m *testing.M) {
	// Initialize the database
	storage.InitDB(true)

	// Run the tests
	exitCode := m.Run()

	// Clean up (e.g., close the database connection if needed)
	// No specific cleanup is needed for SQLite as it handles this internally

	// Exit with the proper code
	os.Exit(exitCode)
}
