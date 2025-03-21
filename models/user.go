package models

import (
	"errors"
	"example.com/sample-api/database"
	"example.com/sample-api/utilities"
)

type User struct {
	ID           int64
	EmailAddress string `binding:"required"`
	Password     string `binding:"required"`
}

func (user *User) Create() error {
	query := `
		INSERT INTO users (emailAddress, password)
		VALUES ($1, $2)
	`
	statement, err := database.DATABASE.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	hashedPassword, err := utilities.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := statement.Exec(user.EmailAddress, hashedPassword)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	user.ID = userID
	return err
}

func (user *User) ValidateCredentials() error {
	query := `
		SELECT password 
		FROM users 
		WHERE emailAddress = ?
	`
	row := database.DATABASE.QueryRow(query, user.EmailAddress)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)
	if err != nil {
		return errors.New("user-not-found")
	}

	passwordIsValid := utilities.ComparePassword(user.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("invalid-password")
	}

	return nil
}
