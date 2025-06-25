package main

import (
	"telu-event-apps/db"
	"telu-event-apps/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	db.InitDB()

	// Create Gin engine that is equipped with two basic middleware, which are Logger and Recovery.
	server := gin.Default()

	routes.RegisterRoutes(server)

	// Listen and serve on ":8080"
	server.Run(":9090")
}
