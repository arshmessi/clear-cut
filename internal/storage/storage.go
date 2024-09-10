package storage

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db   *sql.DB
	once sync.Once
)

// InitDB initializes the SQLite database and creates tables if they do not exist
func InitDB(useTestDB bool) {
	once.Do(func() {
		var err error
		dbPath := "./db/sqlite.db"
		if useTestDB {
			dbPath = ":memory:" // Use an in-memory database for tests
		}
		db, err = sql.Open("sqlite3", dbPath)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}

		createTables()
	})
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}

func createTables() {
	// Create tables
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS groups (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create groups table: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS expenses (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			description TEXT NOT NULL,
			amount REAL NOT NULL,
			paid_by TEXT NOT NULL,
			group_id INTEGER,
			FOREIGN KEY (group_id) REFERENCES groups(id)
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create expenses table: %v", err)
	}
}
