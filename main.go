package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ice-cream-backend/routes/v1"
	completed_tasks_router "github.com/ice-cream-backend/routes/v1/completed_tasks"
	flowers_router "github.com/ice-cream-backend/routes/v1/flowers"
	gardens_router "github.com/ice-cream-backend/routes/v1/gardens"
	rules_router "github.com/ice-cream-backend/routes/v1/rules"
	users_router "github.com/ice-cream-backend/routes/v1/users"
	"github.com/ice-cream-backend/utils"
)

func createServer() {
	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomePage).Methods("GET")

	router.HandleFunc("/icecream-test", routes.TestMongoDB).Methods("GET")

	// gardens
	router.HandleFunc("/api/v1/gardens", gardens_router.CreateGardens).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/gardens/{gardenId}", gardens_router.GetGardenByGardenId).Methods("GET")
	router.HandleFunc("/api/v1/gardens/userid/{userFireBaseId}", gardens_router.GetGardensByUserId).Methods("GET")
	router.HandleFunc("/api/v1/gardens/{gardenId}", gardens_router.UpdateGardenById).Methods("PUT")
	router.HandleFunc("/api/v1/gardens/{gardenId}", gardens_router.DeleteGardenByGardenId).Methods("DELETE")

	// rules
	router.HandleFunc("/api/v1/rules", rules_router.CreateRule).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/rules/bulk", rules_router.CreateRules).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/rules/{ruleId}", rules_router.UpdateRuleByRuleId).Methods("PUT")

	// completedTasks
	router.HandleFunc("/api/v1/completedTasks", completed_tasks_router.CreateCompletedTasks).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/completedTasks/{completedTaskId}", completed_tasks_router.DeleteCompletedTaskByCompletedTaskId).Methods("DELETE")

	// flowers
	router.HandleFunc("/api/v1/flowers", flowers_router.GetFlowers).Methods("GET")

	// users
	router.HandleFunc("/api/v1/users/{userId}", users_router.GetUserByUserId).Methods("GET")

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func init() {
	utils.LoadEnv()
}

func main() {
	createServer()
}
