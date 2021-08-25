package rules_controllers

import (
	"fmt"
	"log"

	rules_models "github.com/ice-cream-backend/rules/models"
)

func CreateRules(rulesPost rules_models.RulesPost)  {
	log.Println("came into create rules controller with post:", rulesPost)
	fmt.Printf("%+v\n", rulesPost)
}