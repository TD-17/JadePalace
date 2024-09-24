package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
func RunPang(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "pang",
	})
}
func RunPrng(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"message": "prng",
	})
}
func RunPlng(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "plng",
	})
}
