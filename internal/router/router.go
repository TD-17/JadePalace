package router

import (
	"github.com/gin-gonic/gin"

	"github.com/TD17/jade-palace/internal/handler"
)

func SetupRouter(r *gin.RouterGroup) {
	// var r *gin.Engine
	// r = gin.Default()
	//r := gin.Default()

	r.POST("/ping", handler.RunPing)
	r.POST("/chat", handler.ChatHandler)
	r.GET("/token", handler.HandlerToken)
	r.GET("/welcome", handler.Welcome)
	r.GET("/register", handler.RegisterHandler)       // API returning URL
	r.GET("/register-page", handler.ShowRegisterPage) // Serves HTML Page
	r.POST("/verify", handler.VerifyUser)             // Processes verification

}
