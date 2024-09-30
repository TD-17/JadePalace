package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("error reading config file: %v\n", err)
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	} else {
		fmt.Println("Config file loaded successfully")
	}

	viper.AutomaticEnv()
	viper.BindEnv("openai_api_key")

	fmt.Println("OpenAI key", viper.GetString("openai_api_key"))
	fmt.Println("Server port", viper.GetInt("server.port"))
}
