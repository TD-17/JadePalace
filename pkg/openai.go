package pkg

import (
	"errors"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

var OpenAiClient *openai.Client

func OpenaiInit() error {
	apiKey := viper.GetString("openai_api_key")
	if apiKey == "" {
		return errors.New("open ai api key not found")
	}

	OpenAiClient = openai.NewClient(apiKey)
	return nil
}
