package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/books", getBooks)
	server.GET("/books/:id", getBook)
	server.POST("/books", enterNewBook)
	server.PUT("/books/:id", updateBook)
	server.DELETE("/books/:id", deleteBook)
}
