package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register page handler
func RegisterHandler(c *gin.Context) {
	userID := c.Query("user_id") //Get user id from slack

	// Send response as JSON so Slack bot can send the link
	c.JSON(http.StatusOK, gin.H{
		"url": "http://localhost:8080/user/register-page?user_id=" + userID,
	})

}

// HTML Registration page
func ShowRegisterPage(c *gin.Context) {
	userID := c.Query("user_id")
	c.HTML(http.StatusOK, "register.html", gin.H{"user_id": userID})
}
