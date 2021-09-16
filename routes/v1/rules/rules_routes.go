package rules_router

import (
	"encoding/json"
	"fmt"
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
		var newRule rules_models.Rules
		_ = json.NewDecoder(r.Body).Decode(&newRule)

		if rules_models.RuleValidation(newRule) {
			res, err := rules_controllers.CreateRule(newRule)

			if err != nil {
				utils.SendErrorBack(w, err, "Error creating rules!")
			} else {
				newRule := rules_controllers.GetRulesByRuleId(res.InsertedID)

				utils.SendResponseBack(w, newRule, http.StatusCreated)
			}
		} else {
			utils.SendErrorBack(w, fmt.Errorf("missing required fields for creating rule: %+v", newRule), "Missing required fields to create rule")
		}
	}
}

func CreateRules(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: create multiple rules")
	utils.EnableCors(&w)

	if r.Method == "POST" {
		start := utils.StartPerformanceTest()
		var multipleRulesPost []rules_models.Rules
		_ = json.NewDecoder(r.Body).Decode(&multipleRulesPost)

		var invalidRule bool
		for _, rule := range multipleRulesPost {
			if !rules_models.RuleValidation(rule) {
				invalidRule = true
				break;
			}
		}

		if !invalidRule {
			res, err := rules_controllers.CreateRules(multipleRulesPost)

			if err != nil {
				utils.SendErrorBack(w, err, "Error creating multiple rules!")
			} else {
				newRules := rules_controllers.GetRulesByRuleIds(res.InsertedIDs)

				utils.SendResponseBack(w, newRules, http.StatusCreated)
				utils.StopPerformanceTest(start, "Successful create rules (routes)")
			}
		} else {
			utils.SendErrorBack(w, fmt.Errorf("missiing required fields for creating bulk rules: %+v", multipleRulesPost), "Missing required fields to create bulk rules")
		}
	}
}

func UpdateRuleByRuleId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: update rule by id")
	vars := mux.Vars(r)
	utils.EnableCors(&w)

	if r.Method == "PUT" {
		var rule rules_models.Rules
		_ = json.NewDecoder(r.Body).Decode(&rule)

		paramsRuleId := vars["ruleId"]

		if rules_models.RuleValidation(rule) {
			oid, err := primitive.ObjectIDFromHex(paramsRuleId)
	
			if err != nil {
				utils.SendErrorBack(w, err, "Invalid ruleId provided")
			} else {
				res, err := rules_controllers.UpdateRuleByRuleId(oid, rule)
	
				if err != nil {
					utils.SendErrorBack(w, err, "Error updating rule")
				}
	
				if res.MatchedCount != 0 {
					updatedRule := rules_controllers.GetRulesByRuleId(oid)
	
					utils.SendResponseBack(w, updatedRule, http.StatusOK)
				} else {
					utils.SendErrorBack(w, fmt.Errorf("could not find matching rule ID: %s", oid), "Error updating rule: no matching oid")
				}
			}
		} else {
			utils.SendErrorBack(w, fmt.Errorf("missing required fields to update rule: %+v", rule), "Missing required fields to update rule")
		}
	}
}
