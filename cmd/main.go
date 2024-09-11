package main

import (
	"clear-cut/config"
	"clear-cut/internal/handlers"
	"clear-cut/internal/storage"
	"log"
	"net/http"
)

// CORS middleware to handle CORS preflight requests and set headers
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Adjust this if needed to restrict allowed origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")


		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize the database
	storage.InitDB(false)

	// Create a new ServeMux and define the routes
	mux := http.NewServeMux()
	mux.HandleFunc("/register", handlers.RegisterHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/profile", handlers.ProfileHandler)
	mux.HandleFunc("/groups", handlers.CreateGroupHandler) // POST
	mux.HandleFunc("/groups/list", handlers.GetGroupsHandler) // GET
	mux.HandleFunc("/expenses", handlers.AddExpenseHandler) // POST
	mux.HandleFunc("/expenses/list", handlers.GetExpensesHandler) // GET

	// Wrap the mux with CORS middleware
	corsMux := CORS(mux)

	// Start the server using the port from configuration
	log.Printf("Server starting on port %s...", config.AppConfig.ServerPort)
	if err := http.ListenAndServe(":"+config.AppConfig.ServerPort, corsMux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
