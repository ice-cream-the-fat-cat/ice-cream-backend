package rules_routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	rules_controllers "github.com/ice-cream-backend/rules/controllers"
	rules_models "github.com/ice-cream-backend/rules/models"
	"github.com/ice-cream-backend/utils"
)

func CreateRules(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	var rulesPost rules_models.RulesPost
	_ = json.NewDecoder(r.Body).Decode(&rulesPost)

	rules_controllers.CreateRules(rulesPost)
	fmt.Fprintf(w, "Welcome to create rules!")
	fmt.Println("Endpoint hit: create rules")
}