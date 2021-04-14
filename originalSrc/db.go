package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func ConnectDB() (collection *mongo.Collection) {
	mongoConnectionString, exists := os.LookupEnv("MONGODB_CONNECTION_STRING")

	if !exists {
		log.Fatal("MongoDB Connection String is not defined")
	}

	// Set MongoDB Client Options
	clientOptions := options.Client().ApplyURI(mongoConnectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("incforge").Collection("notes")

	return collection
}
