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

	// Verify the token and extract the email and user ID
	email, userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token. Please provide a valid token."})
		return
	}

	// Set the email and user ID in the context for use in subsequent handlers
	context.Set("email", email)
	context.Set("userId", userId)
	
	// Call the next handler in the chain
	// This is important to allow the request to proceed to the next middleware or handler
	context.Next()
}
