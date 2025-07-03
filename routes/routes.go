package routes

import (
	"telu-event-apps/middlewares"

	"github.com/gin-gonic/gin"
)

// *gin.Engine is type gin.Default() return
func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// Create a new group for authenticated routes
	// This group will contain routes that require authentication
	// The routes in this group will be prefixed with "/"
	// and will use the Authenticate middleware
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancleRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
