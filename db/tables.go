package db

import (
	"database/sql"
	"fmt"
	"log"
)

func createTables(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);
	`)
	if err != nil {
		return fmt.Errorf("error creating tables: %v", err)
	}

	log.Println("Tables created or already exist.")
	return nil
}
