package models

import "example.com/the-library-with-go/db"

type User struct {
	ID        int64
	FirstName string `json:"first-name" binding:"required"`
	LastName  string `json:"last-name" binding:"required"`
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
