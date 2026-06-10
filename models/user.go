package models

import (
	"errors"

	"example.com/the-library-with-go/db"
	"example.com/the-library-with-go/utils"
)

type User struct {
	ID        int64
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Email     string `binding:"required"`
	Password  string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users(first_name, last_name, email, password)
	VALUES (?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	u.ID = id
	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&u.ID, &retrivedPassword)
	if err != nil {
		return err
	}

	if !utils.CheckPasswordHash(u.Password, retrivedPassword) {
		return errors.New("Failed to login, credentials invalid")
	}

	return nil
}
