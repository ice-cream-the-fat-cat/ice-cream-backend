package routes

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the API Gateway!")
	fmt.Println("Endpoint hit: homepage")
}
