package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Process User Verification
func VerifyUser(c *gin.Context) {
	userID := c.PostForm("user_id") // Get user ID from form submission

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}
	// Store user registration in the database (Dummy response for now)
	c.JSON(http.StatusOK, gin.H{"message": "User " + userID + " registered successfully!"})
}
