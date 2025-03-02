package server

import (
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

	// Start the server
	r.Run(":8080")
}
