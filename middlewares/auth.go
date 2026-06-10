package middlewares

import (
	"net/http"
	"strings"

	"example.com/the-library-with-go/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Please login first"})
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")

	userID, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	context.Set("userID", userID)

	context.Next()
}
