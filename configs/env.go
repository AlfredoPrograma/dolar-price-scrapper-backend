package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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
