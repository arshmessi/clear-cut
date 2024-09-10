package config

import (
	"log"
	"os"
)

// Config holds the configuration settings for the application.
type Config struct {
	ServerPort string // Port on which the server will run
	DbPath     string // Path to the SQLite database file
	JWTSecret  string // Secret key for JWT authentication
}

// AppConfig is a global variable to hold the application's configuration
var AppConfig Config

// LoadConfig initializes the configuration with environment variables or default values.
func LoadConfig() {
	// Get server port, default to 8080 if not set
	port := getEnv("SERVER_PORT", "8080")

	// Get database path, default to ./db/sqlite.db if not set
	dbPath := getEnv("DB_PATH", "./db/sqlite.db")

	// Get JWT secret, it's required and should not be empty
	jwtSecret := getEnv("JWT_SECRET", "")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is required but not set")
	}

	AppConfig = Config{
		ServerPort: port,
		DbPath:     dbPath,
		JWTSecret:  jwtSecret,
	}

	log.Println("Configuration loaded successfully")
}

// getEnv retrieves environment variables or returns a default value
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
