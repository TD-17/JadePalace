package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func dbInstance() *mongo.Client {
	MongoDb := viper.GetString("mongodb_url")
	if MongoDb == "" {
		log.Fatal("mongo db url not found")
	}
	//here ctx represents a "deadline" context that expires after 10 seconds.
	//here cancel ensures that after the connection is attempted, the resources tied to this context (including the timeout) are freed.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Database!")
	return client
}

func DbInit() {
	Client = dbInstance()
}

func OpenCollection(collectionName string) (*mongo.Collection, error) {
	var collection *mongo.Collection = Client.Database("jade_palace").Collection(collectionName)
	if collection == nil {
		return nil, errors.New("no Collection was found")
	}
	return collection, nil
}
