package router

import (
	"github.com/gin-gonic/gin"

	"github.com/TD17/jade-palace/internal/handler"
)

func AuthRouter(r *gin.RouterGroup) {
	// r := gin.Default()

	r.POST("/users/signup", handler.Signup())
	// r.POST("/users/login", handler.Login())
	//This request will come from slack bot
	r.POST("/users/register", handler.Register)

}
