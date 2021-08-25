package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the API Gateway!")
	fmt.Println("Endpoint hit: homepage")
}

func Gardens(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Gardens API  version:%v id:%v\n", vars["version"], vars["id"])
	fmt.Println("Endpoint hit: gardens")
}
