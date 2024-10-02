package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/TD17/jade-palace/internal/services"
)

func RunPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func ChatHandler(c *gin.Context) {
	//services.OpenAIChat("hi there")
	services.SampleOpenAI()
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
