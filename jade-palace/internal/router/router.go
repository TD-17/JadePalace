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
	r.GET("/pang", handler.RunPang)
	r.GET("/prng", handler.RunPrng)
	r.GET("/plng", handler.RunPlng)
	r.GET("/chatgpt", handler.GetGPT)

	return r
}
