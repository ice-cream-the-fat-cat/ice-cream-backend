package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv()  {
	err := godotenv.Load(".env")

	if err != nil && os.Getenv("GO_ENV") != "production" {
		log.Fatal("Error loading .env file")
	}
}
