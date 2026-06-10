package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)

	server.GET("/books", getBooks)
	server.GET("/books/:id", getBook)
	server.POST("/books", enterNewBook)
	server.PUT("/books/:id", updateBook)
	server.PUT("/books/:id/lend", updateLendStatus)
	server.DELETE("/books/:id", deleteBook)
}
