package server

import "github.com/TD17/jade-palace/internal/router"

func Server() {
	// TODO: Add a logger
	r := router.SetupRouter()
	r.Run()
}
