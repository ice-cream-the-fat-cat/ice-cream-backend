package rules_router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	rules_controllers "github.com/ice-cream-backend/controllers/v1/rules"
	rules_models "github.com/ice-cream-backend/models/v1/rules"
	"github.com/ice-cream-backend/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateRule(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: create rules")
	utils.EnableCors(&w)

	if r.Method == "POST" {
		var rulesPost rules_models.Rules
		_ = json.NewDecoder(r.Body).Decode(&rulesPost)
	
		res, err := rules_controllers.CreateRule(rulesPost)
		
		if err != nil {
			fmt.Fprintf(w, "Error creating rules!")
		} else {
			newRule := rules_controllers.GetRulesById(res.InsertedID)
	
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newRule)
		}
	}
}

func CreateRules(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: create multiple rules")
	utils.EnableCors(&w)

	if r.Method == "POST" {
		var multipleRulesPost []rules_models.Rules
		_ = json.NewDecoder(r.Body).Decode(&multipleRulesPost)
	
		res, err := rules_controllers.CreateRules(multipleRulesPost)
	
		if err != nil {
			fmt.Fprintf(w, "Error creating multiple rules!")
		} else {
			newRules := rules_controllers.GetRulesByRuleIds(res.InsertedIDs)
	
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newRules)
		}
	}
}

func EditRuleByRuleId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: create edit rule by id")
	vars := mux.Vars(r)
	utils.EnableCors(&w)

	var rule rules_models.Rules
	_ = json.NewDecoder(r.Body).Decode(&rule)

	paramsRuleId := vars["ruleId"]

	oid, err := primitive.ObjectIDFromHex(paramsRuleId)

	if err != nil {
		log.Println("Error converting params ruleId to ObjectId:", err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Invalid ruleId provided")
	} else {
		res, err := rules_controllers.UpdateRuleByRuleId(oid, rule)

		if err != nil {
			log.Println("Error updating rule:", err)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Error updating rule")
		}

		if res.MatchedCount != 0 {
			updatedRule := rules_controllers.GetRulesByRuleId(oid)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedRule)
		} else {
			log.Println("could not find matching rule ID:", oid)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Error updating rule: no matching oid")
		}
	}

	
}