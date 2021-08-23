package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ice-cream-backend/routes"
	"github.com/ice-cream-backend/utils"
)

func createServer()  {
	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomePage)

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":" + port, router))
}

func init()  {
	utils.LoadEnv()	
}

func main()  {
	createServer()
}