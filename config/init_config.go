package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var OpenAIKey, MongoDBURL, DBPassword, PublicKey, PrivateKey string

func InitConfig() {
	viper.AutomaticEnv() // Automatically read from environment variables

	// Read from environment variables
	OpenAIKey = os.Getenv("OPENAI_API_KEY")
	MongoDBURL = os.Getenv("MONGODB_URL")
	DBPassword = os.Getenv("DB_PASSWORD")
	PublicKey = os.Getenv("PUBLIC_KEY")
	PrivateKey = os.Getenv("PRIVATE_KEY")

	// Print configuration for debugging
	fmt.Println("✅ Configuration Loaded:")
	fmt.Println("OpenAI API Key:", OpenAIKey)
	fmt.Println("MongoDB URL:", MongoDBURL)
	fmt.Println("Database Password:", DBPassword)
	fmt.Println("Public Key:", PublicKey)
	fmt.Println("Private Key:", PrivateKey)

	// Check if required variables are missing
	if OpenAIKey == "" || MongoDBURL == "" || DBPassword == "" {
		panic("❌ ERROR: Missing required environment variables")
	}
}
