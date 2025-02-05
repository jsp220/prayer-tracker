package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// MongoDB Client
var Client *mongo.Client

// ConnectMongoDB initializes the database connection
func ConnectMongoDB() {
	// Load MongoDB URI from environment variables
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI is not set in .env file")
	}

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not ping MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")

	// Assign client to global variable
	Client = client
}

// GetCollection returns a MongoDB collection
func GetCollection(dbName, collectionName string) *mongo.Collection {
	return Client.Database(dbName).Collection(collectionName)
}

// DisconnectMongoDB closes the MongoDB connection
func DisconnectMongoDB() {
	if Client != nil {
		err := Client.Disconnect(context.Background())
		if err != nil {
			log.Fatal("Error disconnecting MongoDB:", err)
		}
		fmt.Println("Disconnected from MongoDB!")
	}
}
