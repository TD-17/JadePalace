package server

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/TD17/jade-palace/internal/router"
)

func Server() {
	// Add a logger
	r := gin.Default()

	// Load the HTML templates
	r.LoadHTMLGlob("templates/*")

	// Group routes under specific prefixes and add routes from different files
	authGroup := r.Group("/auth")
	router.AuthRouter(authGroup)

	userGroup := r.Group("/user")
	router.SetupRouter(userGroup)

	// Get the PORT from the environment (required for Render)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 for local testing
	}

	// Start the server on the assigned port
	fmt.Printf("Server running on port %s\n", port)
	r.Run(":" + port)
}
