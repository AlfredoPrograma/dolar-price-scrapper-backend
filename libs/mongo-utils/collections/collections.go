package collections

import (
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

func StringifyCollectionName(s string) string {
	_, base := strings.Split(s, ".")[0], strings.Split(s, ".")[1]
	upperIndexes := []int{}
	collectionNameArr := strings.Split(strings.ToLower(base), "")
	collectionName := ""

	for i, char := range base {
		if i == 0 {
			continue
		}

		if string(char) != strings.ToLower(string(char)) {
			upperIndexes = append(upperIndexes, i)
		}
	}

	for _, v := range upperIndexes {
		// dollarprices
		firstPart, secondPart := collectionNameArr[:v], collectionNameArr[v:]

		collectionName = strings.Join(firstPart, "") + "-" + strings.Join(secondPart, "")
	}

	return collectionName
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	DB_NAME := os.Getenv("DB_NAME")

	collection := client.Database(DB_NAME).Collection(collectionName)
	return collection
}
