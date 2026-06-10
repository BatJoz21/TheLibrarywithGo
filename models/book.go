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
	LendStatus  bool
	UserID      int64
}

func (b *Book) Save() error {
	query := `INSERT INTO books(title, genre, description, published, lend_status, user_id) VALUES
	(?, ?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(b.Title, b.Genre, b.Description, b.Published, b.LendStatus, b.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	b.ID = id

	return err
}

func (book Book) Update() error {
	query := `UPDATE books
	SET title = ?, genre = ?, description = ?, published = ?
	WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(book.Title, book.Genre, book.Description, book.Published, book.ID)

	return err
}

func (book Book) Delete() error {
	query := `DELETE FROM books WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(book.ID)

	return err
}

func UpdateLendStatus(id int64) error {
	book, err := GetBookByID(id)
	if err != nil {
		return err
	}
	status := !book.LendStatus

	query := `UPDATE books SET lend_status = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(status, id)

	return err
}

func GetAllBooks() ([]Book, error) {
	query := `SELECT * FROM books`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err = rows.Scan(&book.ID, &book.Title, &book.Genre, &book.Description, &book.Published, &book.LendStatus, &book.UserID)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func GetBookByID(id int64) (*Book, error) {
	query := `SELECT * FROM books WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var book Book
	err := row.Scan(&book.ID, &book.Title, &book.Genre, &book.Description, &book.Published, &book.LendStatus, &book.UserID)
	if err != nil {
		return nil, err
	}

	return &book, nil
}
