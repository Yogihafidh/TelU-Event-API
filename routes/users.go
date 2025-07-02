package routes

import (
	"net/http"
	"telu-event-apps/models"
	"telu-event-apps/utils"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	// Extracting incoming JSON data into an Event struct
	var user models.User
	err := context.ShouldBindJSON(&user)

	// Error handling
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create users. Try again later."})
		return
	}

	// Cek apakah email sudah dipakai sebelum insert user baru.
	exists, err := user.ExistsByEmail()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	if exists {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Email already registered"})
		return
	}

	// Save data base on the Event struct
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	// Return success response
	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully.", "user": gin.H{
		"id":    user.ID,
		"email": user.Email,
	}})
}

func login(context *gin.Context) {
	// 1. Extracting incoming JSON data into an Event struct
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not login. Try again later."})
		return
	}

	// 2. Validate user credentials
	isValid, err := user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	if !isValid {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Email or password incorrect."})
		return
	}

	// 3. if credentials are valid, generate a JWT token
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token. Try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successfully.", "token": token})

}
