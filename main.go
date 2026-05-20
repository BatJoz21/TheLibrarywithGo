package main

import (
	"net/http"

	"example.com/the-library-with-go/db"
	"example.com/the-library-with-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/books", getBooks)
	server.POST("/books", enterNewBook)

	server.Run(":8080") // localhost:8080
}

func getBooks(context *gin.Context) {
	books := models.GetAllBooks()
	context.JSON(http.StatusOK, books)
}

func enterNewBook(context *gin.Context) {
	var book models.Book
	err := context.ShouldBindJSON(&book)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	book.ID = 1
	book.UserID = 1

	book.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "book": book})
}
