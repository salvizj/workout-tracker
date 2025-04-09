package services

import (
	db "workout_tracker/db/queries"
	"workout_tracker/utils"
)

func SignUp(email, password string) (error error) {
	err := utils.ValidateUserForm(email, password)
	if err != nil {
		return err
	}

	err = db.InsertUser(email, password)
	if err != nil {
		return err
	}

	return nil
}
