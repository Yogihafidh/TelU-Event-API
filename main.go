package main

import (
	"net/http"
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
	server.POST("/events", createEvents)

	// Listen and serve on ":8080"
	server.Run(":9090")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	// Extracting incoming JSON data into an Event struct
	var event models.Event
	err := context.ShouldBindJSON(&event)
	event.ID = 1
	event.UserID = 1

	// Error handling
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	// Save data
	event.Save()

	// Return success response
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully.", "event": event})
}
