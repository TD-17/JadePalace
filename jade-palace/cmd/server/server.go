package server

import "github.com/TD17/jade-palace/internal/router"

func Server() {
	r := router.SetupRouter()
	r.Run()
}
