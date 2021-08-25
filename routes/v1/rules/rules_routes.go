package rules_router

import (
	"encoding/json"
	"fmt"
	"net/http"

	rules_controllers "github.com/ice-cream-backend/controllers/v1/rules"
	rules_models "github.com/ice-cream-backend/models/v1/rules"
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
		newRule := rules_controllers.GetRules(res.InsertedID)

		// fmt.Fprintln(w, "Successfully created rules:", res.InsertedID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newRule)
	}
}