package rules_routes

import (
	"fmt"
	"net/http"
)

func CreateRules(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to create rules!")
	fmt.Println("Endpoint hit: create rules")
}