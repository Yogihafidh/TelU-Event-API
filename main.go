package main

import (
	"net/http"
	"strconv"
	"telu-event-apps/db"
	"telu-event-apps/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	db.InitDB()

	// Create Gin engine that is equipped with two basic middleware, which are Logger and Recovery.
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvents)

	// Listen and serve on ":8080"
	server.Run(":9090")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}

	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later."})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvents(context *gin.Context) {
	// Extracting incoming JSON data into an Event struct
	var event models.Event
	err := context.ShouldBindJSON(&event)
	event.ID = 1
	event.UserID = 1

	// Error handling
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create events. Try again later."})
		return
	}

	// Save data
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event"})
		return
	}
	// Return success response
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully.", "event": event})
}
