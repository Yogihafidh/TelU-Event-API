package routes

import (
	"log"
	"net/http"
	"strconv"
	"telu-event-apps/models"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	// Extracting user ID from the context, set by the authenticate middleware
	userId := context.GetInt64("userId")
	log.Println("User ID from context:", userId)

	// Extracting the event ID from the URL parameter and converting it to int64
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	// Check if the event exists
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later."})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Successfully registered for event", "event": event, "userID": userId})

}
func cancleRegistration(context *gin.Context) {}
