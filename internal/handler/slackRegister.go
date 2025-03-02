package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/TD17/jade-palace/database"
	"github.com/TD17/jade-palace/internal/models"
)

// func Register() gin.HandlerFunc {
// 	userCollection, err := database.OpenCollection("slack_user")
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
// 		defer cancel() // Defer once, so it will apply to the entire context use
// 		var user models.Slackuser

// 		err := c.BindJSON(&user)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		validationError := validate.Struct(user)
// 		if validationError != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": validationError.Error()})
// 			return
// 		}

// 		// Return the count of the users
// 		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
// 		if err != nil {
// 			log.Panic(err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
// 		}

// 		if count > 0 {
// 			c.JSON(http.StatusConflict, gin.H{"error": "This email already exists"})
// 			return
// 		}

// 		user.ID = primitive.NewObjectID()
// 		user.User_id = user.ID.Hex()
// 		token, refreshToken, _ := pkg.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
// 		user.Token = &token
// 		user.Refresh_token = &refreshToken

// 		// Add data to the database
// 		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
// 		if insertErr != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"err": insertErr.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, resultInsertionNumber)
// 	}

// }

func Register(c *gin.Context) {
	userCollection, err := database.OpenCollection("slack_user")
	if err != nil {
		log.Panic(err)
	}
	// Parse the incoming JSON request
	var user models.Slackuser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Check if user already exists
	var existingUser models.Slackuser
	err = userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	// Generate a token (placeholder logic; replace with real token generation)
	token := "generated_token_12345"

	// Insert the new user into the database
	_, err = userCollection.InsertOne(ctx, bson.M{
		"first_name": user.First_name,
		"last_name":  user.Last_name,
		"password":   user.Password, // Note: Hash passwords in production
		"email":      user.Email,
		"status":     user.Status,
		"token":      token,
	})

	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Return success response with token
	c.JSON(http.StatusOK, gin.H{
		"message": "Registration successful",
		"token":   token,
	})
}
