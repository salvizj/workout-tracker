package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"workout_tracker/global"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(config global.Config) (*sql.DB, error) {
	if _, err := os.Stat(config.DatabasePath); os.IsNotExist(err) {
		log.Printf("SQLite database file does not exist, creating %s\n", config.DatabasePath)
	}

	db, err := sql.Open("sqlite3", config.DatabasePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("failed to create tables: %v", err)
	}

	return db, nil
}
