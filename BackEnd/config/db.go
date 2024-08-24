package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Client

func ConnectDB() {
    uri := os.Getenv("MONGODB_URI")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()


    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal("MongoDB connection error: ", err)
    }


    if err := client.Ping(ctx, readpref.Primary()); err != nil {
        log.Fatal("Could not ping MongoDB: ", err)
    }

    DB = client
    log.Println("Connected to MongoDB")
}

func DisconnectDB() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := DB.Disconnect(ctx); err != nil {
        log.Fatal("Error disconnecting from MongoDB: ", err)
    }
    log.Println("Disconnected from MongoDB")
}
