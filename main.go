package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ice-cream-backend/routes"
	rules_routes "github.com/ice-cream-backend/rules/routes"
	"github.com/ice-cream-backend/utils"
)

func createServer()  {
	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomePage).Methods("GET")

	router.HandleFunc("/icecream-test", routes.TestMongoDB).Methods("GET")

	
	
	router.HandleFunc("/api/v1/rules", rules_routes.CreateRules).Methods("POST")
	
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":" + port, router))
}

func init()  {
	utils.LoadEnv()	
}

func main()  {
	createServer()
}