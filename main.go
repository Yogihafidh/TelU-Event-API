package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create Gin engine that is equipped with two basic middleware, which are Logger and Recovery.
	server := gin.Default()

	// Listen and serve on ":8080"
	server.Run(":8080") 
}
