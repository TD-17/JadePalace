package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TD17/jade-palace/internal/models"
)

func SampleOpenAI(message string) (*models.Response, error) {
	jsonFile, err := os.Open("sample_response.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return nil, err
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)

	var response models.Response
	err = decoder.Decode(&response)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	// Access and use the parsed data
	fmt.Println("ID:", response.ID)
	fmt.Println("Content:", response.Choices[0].Message.Content)
	fmt.Println("The question was:", message)
	return &response, nil

}
