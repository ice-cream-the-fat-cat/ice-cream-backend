package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	errors_models "github.com/ice-cream-backend/models/v1/errors"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil && os.Getenv("GO_ENV") != "production" {
		log.Fatal("Error loading .env file")
	}
}

func EnableCors(w *http.ResponseWriter) {
	header := (*w).Header()
	// TODO: Limit access to just frontend domains
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Access-Control-Allow-Methods", "DELETE, POST, PUT, GET, OPTIONS")
	header.Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
}

func StartPerformanceTest() time.Time {
	return time.Now()
}

func StopPerformanceTest(start time.Time, message string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", message, elapsed)
}

func ConvertAPIStringToDate(date string) time.Time {
	splitDate := strings.Split(date, "-")
	year, _ := strconv.Atoi(splitDate[0])
	month, _ := strconv.Atoi(splitDate[1])
	day, _ := strconv.Atoi(splitDate[2])

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func SendErrorBack(w http.ResponseWriter, err error, info string) {
	log.Println(info + ":", err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest);
	var iceCreamError errors_models.IceCreamErrors
	iceCreamError.Error = err.Error()
	iceCreamError.Info = info
	json.NewEncoder(w).Encode(iceCreamError)
}
