package main

import (
	"log"
	"test_lamarria/be/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	r := gin.Default()

	// Setup the routes
	router.SetupRoutes(r)

	// Start the server on port 8080
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Unable to start the server: ", err)
	}
}
