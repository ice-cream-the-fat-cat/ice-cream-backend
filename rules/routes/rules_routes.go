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
	fmt.Println("Endpoint hit: create rules")
	utils.EnableCors(&w)
	var rulesPost rules_models.Rules
	_ = json.NewDecoder(r.Body).Decode(&rulesPost)

	res, err := rules_controllers.CreateRules(rulesPost)
	
	if err != nil {
		fmt.Fprintf(w, "Error creating rules!")
	} else {
		fmt.Fprintln(w, "Successfully created rules:", res.InsertedID)
	}
}