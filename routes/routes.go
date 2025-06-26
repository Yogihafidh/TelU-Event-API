package routes

import "github.com/gin-gonic/gin"

// *gin.Engine is type gin.Default() return
func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvents)
	server.PUT("/events/:id", updateEvent)
}
