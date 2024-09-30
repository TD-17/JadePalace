package main

import (
	"log"
	"os"

	"github.com/TD17/jade-palace/cmd/server"
	"github.com/TD17/jade-palace/config"
	"github.com/TD17/jade-palace/pkg"
)

func main() {
	config.InitConfig()
	err := pkg.OpenaiInit()

	if err != nil {
		log.Fatal("chatgpt couldn't be initialised")
		os.Exit(0)
	}

	server.Server()
}
