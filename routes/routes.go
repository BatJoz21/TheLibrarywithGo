package routes

import (
	"example.com/the-library-with-go/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)

	server.GET("/books", getBooks)
	server.GET("/books/:id", getBook)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/books", enterNewBook)
	authenticated.PUT("/books/:id", updateBook)
	authenticated.DELETE("/books/:id", deleteBook)

	server.PUT("/books/:id/lend", updateLendStatus)
}
