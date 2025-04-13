package handler

import (
	"context"
	"time"
	"net/http"
	"github.com/TD17/jade-palace/database"
	"github.com/gin-gonic/gin"
)

func StoreSlackInstallation() gin.HandlerFunc {
    return func(c *gin.Context) {
        var payload map[string]interface{}
        if err := c.ShouldBindJSON(&payload); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        collection, err := database.OpenCollection("slack_installations")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open collection"})
            return
        }

        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

        _, err = collection.InsertOne(ctx, payload)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Slack installation stored successfully"})
    }
}
