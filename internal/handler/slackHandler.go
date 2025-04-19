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
            log.Println("❌ Failed to bind JSON:", err)
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        log.Println("✅ Received install data:", installData)
        collection, err := database.OpenCollection("slack_installations")
        if err != nil {
            log.Println("❌ Failed to open collection:", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open collection"})
            return
        }

        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

        _, err = collection.InsertOne(ctx, payload)
        if err != nil {
            log.Println("❌ Failed to save data to DB:", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save"})
            return
        }
        log.Println("✅ Successfully stored install data")
        c.JSON(http.StatusOK, gin.H{"message": "Slack installation stored successfully"})
    }
}
