package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectDB() *mongo.Client {
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

var DB = connectDB()
