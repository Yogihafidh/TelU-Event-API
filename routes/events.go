package routes

import (
	"net/http"
	"strconv"
	"telu-event-apps/models"

	"github.com/gin-gonic/gin"
)

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
	// 1. Extracting incoming JSON data into an Event struct
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create events. Try again later."})
		return
	}

	// 2. Get userId and email from the context, set by the authenticate middleware
	userId := context.GetInt64("userId")
	email := context.GetString("email")

	// 3. Set the UserID field of the event to the authenticated user's ID
	event.UserID = userId

	// 4. Save data base on the Event struct
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event"})
		return
	}

	// 5. Return success response
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully.", "event": event, "userId": userId, "email": email})
}

func updateEvent(context *gin.Context) {
	// Extracting the event ID from the URL parameter and converting it to int64
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	// Get userId from the context, set by the authenticate middleware
	userId := context.GetInt64("userId")

	// Check if the event exists. If it does not exist, return an error and if it does, continue with the update process
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event. Try again later."})
		return
	}

	// Check if the authenticated user is the owner of the event
	if event.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "You are not authorized to update"})
		return
	}

	// Extracting incoming JSON data into an Event struct
	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent) // Bind the incoming JSON to the updateEvent variable
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event. Try again later."})
		return
	}
	updateEvent.ID = eventId // Set the ID of the event to be updated

	// Update event based on the updateEvent variable
	err = updateEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event. Try again later."})
		return
	}

	// Return success response
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully."})
}

func deleteEvent(context *gin.Context) {
	// Extracting the event ID from the URL parameter and converting it to int64
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	// Get userId from the context, set by the authenticate middleware
	userId := context.GetInt64("userId")

	// Check if the event exists. If it does not exist, return an error and if it does, continue with the update process
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event. Try again later."})
		return
	}

	if userId != event.UserID {
		context.JSON(http.StatusForbidden, gin.H{"message": "You are not authorized to delete this event."})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event. Try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})
}
