package discordbot

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// get .env parameters
func getBotParams(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
