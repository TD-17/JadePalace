package router

import (
	"github.com/gin-gonic/gin"

	"github.com/TD17/jade-palace/internal/handler"
)

func SetupRouter() *gin.Engine {
	// var r *gin.Engine
	// r = gin.Default()
	r := gin.Default()

	r.GET("/ping", handler.RunPing)
	r.POST("/chat", handler.ChatHandler)
	// r.POST("/value", handler.UseValue)

	return r
}
