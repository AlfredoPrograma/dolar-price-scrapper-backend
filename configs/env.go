package configs

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const (
	LOCAL = "LOCAL"
	DEV   = "DEV"
	PROD  = "PROD"
)

type EnvVarsMap map[string]string

var envVarsMap = EnvVarsMap{}

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	loadEnvVarsMap()
}

func GetEnvVar(e string) string {
	v := envVarsMap[e]

	return v
}

func loadEnvVarsMap() {
	b, err := os.ReadFile(".env")

	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(b), "\n")

	for _, l := range lines {
		if string(l[0]) == "#" {
			continue
		}

		key, value := strings.Split(l, "=")[0], strings.Split(l, "=")[1]
		envVarsMap[key] = value
	}
}
