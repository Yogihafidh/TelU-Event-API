package routes

import (
	"net/http"
	"telu-event-apps/models"

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
