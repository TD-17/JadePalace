package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var OpenAIKey, MongoDBURL, DBPassword, PublicKey, PrivateKey string

func InitConfig() {
	// Load config.json file for local development
	viper.SetConfigName("config")   // Name of the file (without extension)
	viper.SetConfigType("json")     // JSON format
	viper.AddConfigPath(".")        // Look in the current directory
	viper.AddConfigPath("./config") // Also look in the config/ directory
	viper.AutomaticEnv()            // Override with environment variables if set

	// Try reading config file (if running locally)
	if err := viper.ReadInConfig(); err == nil {
		log.Println("✅ Loaded config.json for local development")
	} else {
		log.Println("⚠️ No config file found, using environment variables")
	}

	// Read values (prioritizing environment variables)
	OpenAIKey = viper.GetString("OPENAI_API_KEY")
	MongoDBURL = viper.GetString("MONGODB_URL")
	DBPassword = viper.GetString("DB_PASSWORD")
	PublicKey = viper.GetString("PUBLIC_KEY")
	PrivateKey = viper.GetString("PRIVATE_KEY")

	// Print for debugging
	fmt.Println("✅ Configuration Loaded:")
	fmt.Println("🔹 OpenAI Key:", OpenAIKey)
	fmt.Println("🔹 MongoDB URL:", MongoDBURL)

	// Ensure required variables exist
	if OpenAIKey == "" || MongoDBURL == "" || DBPassword == "" {
		panic("❌ ERROR: Missing required environment variables or config values")
	}
}
