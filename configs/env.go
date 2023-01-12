package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: Make package to get all .ENV vars easily
// TODO: Move some functions to the correct files (why is ConnectDB in env.go file ??)
func EnvMongoURI(e string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_URI_ENV_MAP := map[string]string{
		"PROD": os.Getenv("MONGO_URI"),
		"TEST": os.Getenv("MONGO_URI_TEST"),
	}

	return DB_URI_ENV_MAP[e]
}

func ConnectDB() *mongo.Client {
	DB_URI := EnvMongoURI("TEST")
	client, err := mongo.NewClient(options.Client().ApplyURI(DB_URI))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected to MongoDB \nENV: %v", DB_URI)

	defer cancel()
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	DB_NAME := os.Getenv("DB_NAME")

	collection := client.Database(DB_NAME).Collection(collectionName)
	return collection
}

var M = ConnectDB()
