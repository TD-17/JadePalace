package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/TD17/jade-palace/internal/router"
)

func Server() {
	log.Println("✅ Server Initialization Completed! Now starting Gin server...")

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

	// **Explicitly print the port for Render**
	fmt.Printf("🚀 Server running on port %s\n", port)
	log.Printf("Listening on 0.0.0.0:%s", port)

	// **Ensure Gin binds to 0.0.0.0 instead of localhost**
	err := r.Run("0.0.0.0:" + port)
	if err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
