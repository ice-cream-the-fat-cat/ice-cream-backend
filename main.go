package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ice-cream-backend/routes/v1"
	completedTasks_router "github.com/ice-cream-backend/routes/v1/completedTasks"
	gardens_router "github.com/ice-cream-backend/routes/v1/gardens"
	rules_router "github.com/ice-cream-backend/routes/v1/rules"
	"github.com/ice-cream-backend/utils"
)

func createServer() {
	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomePage).Methods("GET")

	router.HandleFunc("/icecream-test", routes.TestMongoDB).Methods("GET")

	router.HandleFunc("/{version}/gardens/{id}", gardens_router.GardensIdGet).Methods("GET")
	router.HandleFunc("/{version}/gardens/user/{userid}", gardens_router.GardensUserIdGet).Methods("GET")
	router.HandleFunc("/{version}/gardens/", gardens_router.GardensPost).Methods("POST")

	// rules
	router.HandleFunc("/api/v1/rules", rules_router.CreateRules).Methods("POST")

	// completedTasks
	router.HandleFunc("/api/v1/completedTasks", completedTasks_router.CreateCompletedTasks).Methods("POST")

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func init() {
	utils.LoadEnv()
}

func main() {
	createServer()
}
