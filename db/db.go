package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/library_db?parseTime=true")
	if err != nil {
		panic("Could not connect to database: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = DB.Ping()

	if err != nil {
		panic(err)
	}

	createTables()
}

func createTables() {
	createBooksTable := `CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		title VARCHAR(150) NOT NULL,
		genre VARCHAR(150) NOT NULL,
		description TINYTEXT,
		published DATETIME NOT NULL,
		user_id INTEGER
	)`

	_, err := DB.Exec(createBooksTable)
	if err != nil {
		panic(err.Error())
	}
}
