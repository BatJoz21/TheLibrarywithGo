package routes

import (
	"net/http"
	"strconv"

	"example.com/the-library-with-go/models"
	"github.com/gin-gonic/gin"
)

func getBooks(context *gin.Context) {
	books, err := models.GetAllBooks()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": err.Error()})
		return
	}
	context.JSON(http.StatusOK, books)
}

func getBook(context *gin.Context) {
	bookId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message:": err.Error()})
		return
	}

	book, err := models.GetBookByID(bookId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": err.Error()})
		return
	}

	context.JSON(http.StatusOK, book)
}

func enterNewBook(context *gin.Context) {
	var book models.Book
	err := context.ShouldBindJSON(&book)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	book.UserID = 1
	book.LendStatus = false

	err = book.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "book": book})
}

func updateBook(context *gin.Context) {
	bookID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message:": err.Error()})
		return
	}

	_, err = models.GetBookByID(bookID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": err.Error()})
		return
	}

	var updatedBook models.Book
	err = context.ShouldBindJSON(&updatedBook)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updatedBook.ID = bookID
	err = updatedBook.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message:": "Book update successfully"})
}

func deleteBook(context *gin.Context) {
	bookID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message:": err.Error()})
		return
	}

	book, err := models.GetBookByID(bookID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": err.Error()})
		return
	}

	err = book.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Deleting book successfully"})
}
