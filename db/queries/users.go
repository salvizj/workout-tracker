package db

import (
	"fmt"
	"workout_tracker/global"
)

func InsertUser(email, password string) error {
	query := `INSERT INTO Users (email, password) VALUES (?, ?)`

	_, err := global.DB.Exec(query, email, password)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}
	return nil
}
