package services

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"

	"github.com/TD17/jade-palace/pkg"
)

func OpenAIChat(message string) (*openai.ChatCompletionResponse, error) {
	resp, err := pkg.OpenAiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{

			Model:     "gpt-3.5-turbo",
			MaxTokens: 100,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "system",
					Content: "You are a helpful assistant.",
				},
				{
					Role:    "user",
					Content: message,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("Completion error %v", err)
		//nil can't be variable, it can be memory location. So openai.ChatCompletionResponse will be a pointer
		return nil, err
	}
	fmt.Println(resp.Choices[0].Message.Content)
	return &resp, nil
}
