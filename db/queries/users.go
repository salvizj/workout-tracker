package db

import (
	"fmt"
	types "workout_tracker"
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
func GetUser(email string) (user types.User, err error) {
	query := `SELECT email, password FROM Users WHERE email = ?`
	err = global.DB.QueryRow(query, email).Scan(&user.Email, &user.Password)
	if err != nil {
		return user, fmt.Errorf("failed to get user: %v", err)
	}

	return user, nil
}
