package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv()  {
	err := godotenv.Load(".env")

	if err != nil {
		// TODO: Figure out how to load .env in Heroku
		log.Println("Error loading .env file")
	}
}
