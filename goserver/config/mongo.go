package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoUri() string {
	return os.Getenv("MONGO_URI")
}

func DefaultCtx() context.Context {
	return context.Background()
}

func MongoClient() *mongo.Database {
	otps := options.Client().ApplyURI(MongoUri())

	client, err := mongo.Connect(context.Background(), otps)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Cannot connect to mongodb")
	}
	return client.Database("notes")
}
