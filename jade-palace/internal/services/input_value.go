package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TD17/jade-palace/internal/models"
	"github.com/gin-gonic/gin"
)

// func UserInput(w http.ResponseWriter, r *http.Request) error {
// 	var requestData RequestData
// 	err := json.NewDecoder(r.Body).Decode(&requestData)
// 	if err != nil {
// 		http.Error(w, "Error decoding JSON body", http.StatusBadRequest)
// 		return err
// 	}

// 	// Use the parsed value as a function parameter
// 	processValue(requestData.Value)

// 	// Send a response
// 	fmt.Fprintf(w, "JSON data received and processed successfully")
// 	return nil
// }

// Using Gin
func UserInput(c *gin.Context) {
	// Decode the JSON request body
	var requestData models.RequestData
	err := json.NewDecoder(c.Request.Body).Decode(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding JSON body"})
		return
	}

	// Use the parsed value as a function parameter
	processValue(requestData.Value)

	// Send a response
	c.JSON(http.StatusOK, gin.H{"message": "JSON data received and processed successfully"})
}

func processValue(value string) {
	// Process the value here
	fmt.Println("Received value:", value)
}
