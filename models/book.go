package models

import (
	"time"

	"example.com/the-library-with-go/db"
)

type Book struct {
	ID          int64
	Title       string    `binding:"required"`
	Genre       string    `binding:"required"`
	Description string    `binding:"required"`
	Published   time.Time `binding:"required"`
	UserID      int
}

var books = []Book{}

func (b Book) Save() error {
	query := `INSERT INTO books(title, genre, description, published, user_id) VALUES
	(?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(b.Title, b.Genre, b.Description, b.Published, b.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	b.ID = id

	return err
}

func GetAllBooks() []Book {
	return books
}
