package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

// InitDB initializes a global DB connection
func InitDB() *sql.DB {
	var err error
	dsn := "postgres://dev1:dev1pg@localhost:5432/biomassx" // your connection string
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	return db
}

// GetDB returns the existing DB instance
func GetDB() *sql.DB {
	if db == nil {
		log.Fatal("Database connection is not initialized")
	}
	return db
}
