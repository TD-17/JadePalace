package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Struct to parse user input
type ChatRequest struct {
	Message string `json:"message"`
}

// Struct to interact with OpenAI API
type OpenAIRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

// ChatGPT message structure
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Struct for OpenAI API response
type OpenAIResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
}

// ChatGPT interaction handler
func ChatWithGPT(c *gin.Context) {
	var userInput ChatRequest

	// Bind JSON request body
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Call OpenAI API
	response, err := getChatGPTResponse(userInput.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get response from ChatGPT"})
		return
	}

	// Return ChatGPT's response to user
	c.JSON(http.StatusOK, gin.H{"response": response})
}

// Function to call OpenAI API
func getChatGPTResponse(userMessage string) (string, error) {
	// Get API key from environment variable
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OpenAI API key not set")
	}

	// Prepare OpenAI API request payload
	payload := OpenAIRequest{
		Model: "gpt-3.5-turbo",
		Messages: []ChatMessage{
			{Role: "system", Content: "You are a helpful assistant."},
			{Role: "user", Content: userMessage},
		},
	}

	// Convert payload to JSON
	payloadBytes, _ := json.Marshal(payload)

	// Make request to OpenAI API
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse JSON response
	var openAIResp OpenAIResponse
	err = json.Unmarshal(body, &openAIResp)
	if err != nil {
		return "", err
	}

	// Extract and return ChatGPT's reply
	if len(openAIResp.Choices) > 0 {
		return openAIResp.Choices[0].Message.Content, nil
	}

	return "No response from ChatGPT", nil
}
