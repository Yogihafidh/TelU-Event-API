package middlewares

import (
	"net/http"
	"telu-event-apps/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	// Get the token from the request header to check if the user is authorized
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized. Please provide a valid token."})
		return
	}

	email, userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token. Please provide a valid token."})
		return
	}

	context.Set("email", email)
	context.Set("userId", userId)
	context.Next()
}
