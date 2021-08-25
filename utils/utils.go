package utils

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv()  {
	err := godotenv.Load(".env")

	if err != nil && os.Getenv("GO_ENV") != "production" {
		log.Fatal("Error loading .env file")
	}
}

func EnableCors(w *http.ResponseWriter) {
	header := (*w).Header()
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	header.Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
}
