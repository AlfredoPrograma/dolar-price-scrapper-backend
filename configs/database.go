package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoClient() *mongo.Client {
	mongoUri := fmt.Sprintf(
		"mongodb://%v:%v@database:27017/%v",
		GetEnvVar("DB_USER"),
		GetEnvVar("DB_PASSWORD"),
		GetEnvVar("DB_NAME"),
	)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func ConnectMongoDB(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected to MongoDB \nENV: %v", GetEnvVar("CURRENT_ENV"))
}

var DB *mongo.Client
