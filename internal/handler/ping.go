package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/TD17/jade-palace/internal/models"
	"github.com/TD17/jade-palace/internal/services"
)

func RunPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func ChatHandler(c *gin.Context) {
	//services.OpenAIChat("hi there")
	var requestData models.RequestData
	err := json.NewDecoder(c.Request.Body).Decode(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding JSON body"})
		return
	}
	response, err := services.SampleOpenAI(requestData.Value) //This is a sampple
	// response, err := services.OpenAIChat(requestData.Value) //This is a real deal. Carefull!! It will cost you.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failure",
		})
		return
	}
	// contentArray := strings.Split(response.Choices[0].Message.Content, "\n")
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    response.Choices[0].Message.Content,
	})

}
