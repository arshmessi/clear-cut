package main

import (
	"clear-cut/config"
	"clear-cut/internal/handlers"
	"clear-cut/internal/storage"
	"log"
	"net/http"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize the database
	storage.InitDB(false)

	// Define the routes
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/profile", handlers.ProfileHandler)
	http.HandleFunc("/groups", handlers.CreateGroupHandler) // POST
	http.HandleFunc("/groups/list", handlers.GetGroupsHandler) // GET
	http.HandleFunc("/expenses", handlers.AddExpenseHandler) // POST
	http.HandleFunc("/expenses/list", handlers.GetExpensesHandler) // GET

	// Start the server using the port from configuration
	log.Printf("Server starting on port %s...", config.AppConfig.ServerPort)
	if err := http.ListenAndServe(":"+config.AppConfig.ServerPort, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
