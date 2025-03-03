package main

import (
	"log"
	"os"

	"github.com/TD17/jade-palace/cmd/server"
	"github.com/TD17/jade-palace/config"
	"github.com/TD17/jade-palace/database"
	"github.com/TD17/jade-palace/pkg"
)

func main() {
	log.Println("🔍 Starting config initialization...")
	config.InitConfig()

	log.Println("🔍 Connecting to database...")
	database.DbInit()
	log.Println("✅ Database initialized!")

	log.Println("🔍 Initializing OpenAI...")
	err := pkg.OpenaiInit()
	if err != nil {
		log.Fatal("❌ ChatGPT couldn't be initialized")
		os.Exit(0)
	}
	log.Println("✅ OpenAI initialized!")

	log.Println("🚀 Starting server...")
	server.Server()
}
