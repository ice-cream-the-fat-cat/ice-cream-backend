package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ice-cream-backend/routes/v1"
	completed_tasks_router "github.com/ice-cream-backend/routes/v1/completed_tasks"
	gardens_router "github.com/ice-cream-backend/routes/v1/gardens"
	rules_router "github.com/ice-cream-backend/routes/v1/rules"
	"github.com/ice-cream-backend/utils"
)

func createServer() {
	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomePage).Methods("GET")

	router.HandleFunc("/icecream-test", routes.TestMongoDB).Methods("GET")

	// gardens
	router.HandleFunc("/api/v1/gardens/", gardens_router.CreateGardens).Methods("POST")
	router.HandleFunc("/api/v1/gardens/{gardenId}", gardens_router.GetGardenByGardenId).Methods("GET")
	router.HandleFunc("/api/v1/gardens/userid/{userFireBaseId}", gardens_router.GetGardensByUserId).Methods("GET")

	// rules
	router.HandleFunc("/api/v1/rules", rules_router.CreateRule).Methods("POST")
	router.HandleFunc("/api/v1/rules/bulk", rules_router.CreateRules).Methods("POST")

	// completedTasks
	router.HandleFunc("/api/v1/completedTasks", completed_tasks_router.CreateCompletedTasks).Methods("POST")

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func init() {
	utils.LoadEnv()
}

func main() {
	createServer()
}
