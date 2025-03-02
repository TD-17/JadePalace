package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/TD17/jade-palace/internal/services"
)

func HandlerToken(c *gin.Context) {
	token := services.GetTokenTemp()
	c.JSON(http.StatusOK, gin.H{
		"message": token,
	})
}
