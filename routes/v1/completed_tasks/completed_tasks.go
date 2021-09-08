package completed_tasks_router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	completed_tasks_controllers "github.com/ice-cream-backend/controllers/v1/completed_tasks"
	users_controllers "github.com/ice-cream-backend/controllers/v1/users"
	completed_tasks_models "github.com/ice-cream-backend/models/v1/completed_tasks"
	errors_models "github.com/ice-cream-backend/models/v1/errors"
	"github.com/ice-cream-backend/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var COIN_AFTER_COMPLETED_TASK = 1

func CreateCompletedTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: create completedTasks")
	utils.EnableCors(&w)

	if r.Method == "POST" {
		var completedTasksPost completed_tasks_models.CompletedTasks
		_ = json.NewDecoder(r.Body).Decode(&completedTasksPost)

		ruleIds := []interface{}{completedTasksPost.RuleId}
		checkExistingCompletedTasks := completed_tasks_controllers.GetCompletedTasksByRuleIdWithDate(ruleIds, completedTasksPost.Date)

		if len(checkExistingCompletedTasks) > 0 {
			log.Println("Already has existing completedTask for this rule and date!")
			w.Header().Set("Content-Type", "application/json")
			var iceCreamError errors_models.IceCreamErrors
			iceCreamError.Error = fmt.Errorf("already has existing completedTask for this rule: %q and date: %q", completedTasksPost.RuleId.String(), completedTasksPost.Date).Error()
			iceCreamError.Info = "Already has existing completedTask for this rule and date!"
			json.NewEncoder(w).Encode(iceCreamError)
		} else {
			res, err := completed_tasks_controllers.CreateCompletedTask(completedTasksPost)

			if err != nil {
				fmt.Fprintf(w, "Error creating completedTasks!")
				var iceCreamError errors_models.IceCreamErrors
				iceCreamError.Error = err.Error()
				iceCreamError.Info = "Error Error creating completedTasks!"
				json.NewEncoder(w).Encode(iceCreamError)
			} else {
				_ = completed_tasks_controllers.GetCompletedTasksByCompletedTaskId(res.InsertedID)

				user, err := users_controllers.GetUserByFireBaseUserId(completedTasksPost.FireBaseUserId)

				if err != nil {
					log.Println("Error finding user to update coins after completing task:", err)
					w.Header().Set("Content-Type", "application/json")
					var iceCreamError errors_models.IceCreamErrors
					iceCreamError.Error = err.Error()
					iceCreamError.Info = "Error finding user to update coins after completing task"
					json.NewEncoder(w).Encode(iceCreamError)
				} else {
					user.Balance = user.Balance + COIN_AFTER_COMPLETED_TASK

					updatedUser, err := users_controllers.UpdateUserByUserId(user.ID, user)

					if err != nil {
						log.Println("Error updating user's coins after completing task:", err)
						w.Header().Set("Content-Type", "application/json")
						var iceCreamError errors_models.IceCreamErrors
						iceCreamError.Error = err.Error()
						iceCreamError.Info = "Error updating user's coins after completing task"
						json.NewEncoder(w).Encode(iceCreamError)
					} else {
						w.Header().Set("Content-Type", "application/json")
						json.NewEncoder(w).Encode(updatedUser)
					}
				}
			}
		}
	}
}

func DeleteCompletedTaskByCompletedTaskId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: delete garden by id")
	vars := mux.Vars(r)
	utils.EnableCors(&w)

	if r.Method == "DELETE" {
		paramsCompletedTaskId := vars["completedTaskId"]
		paramsFireBaseUserId := vars["fireBaseUserId"]

		oid, err := primitive.ObjectIDFromHex(paramsCompletedTaskId)

		if err != nil {
			log.Println("Error converting params completedTaskId to ObjectId:", err)
			w.Header().Set("Content-Type", "application/json")
			var iceCreamError errors_models.IceCreamErrors
			iceCreamError.Error = err.Error()
			iceCreamError.Info = "Invalid completedTaskId provided"
			json.NewEncoder(w).Encode(iceCreamError)
		} else {
			res, err := completed_tasks_controllers.DeleteCompletedTaskByCompletedTaskId(oid)

			if err != nil {
				log.Println("Error deleting completedTask:", err)
				w.Header().Set("Content-Type", "application/json")
				var iceCreamError errors_models.IceCreamErrors
				iceCreamError.Error = err.Error()
				iceCreamError.Info = "Error deleting completedTask"
				json.NewEncoder(w).Encode(iceCreamError)
			}

			if res.DeletedCount == 1 {
				user, err := users_controllers.GetUserByFireBaseUserId(paramsFireBaseUserId)

				if err != nil {
					log.Println("Error finding user to update coins after deleting completed task:", err)
					w.Header().Set("Content-Type", "application/json")
					var iceCreamError errors_models.IceCreamErrors
					iceCreamError.Error = err.Error()
					iceCreamError.Info = "Error finding user to update coins after deleting completed task"
					json.NewEncoder(w).Encode(iceCreamError)
				} else {
					user.Balance = user.Balance - COIN_AFTER_COMPLETED_TASK

					updatedUser, err := users_controllers.UpdateUserByUserId(user.ID, user)

					if err != nil {
						log.Println("Error updating user's coins after completing task:", err)
						w.Header().Set("Content-Type", "application/json")
						var iceCreamError errors_models.IceCreamErrors
						iceCreamError.Error = err.Error()
						iceCreamError.Info = "Error updating user's coins after completing task"
						json.NewEncoder(w).Encode(iceCreamError)
					} else {
						w.Header().Set("Content-Type", "application/json")
						json.NewEncoder(w).Encode(updatedUser)
					}
				}
			} else {
				log.Println("could not find matching completedTask ID:", oid)
				w.Header().Set("Content-Type", "application/json")
				var iceCreamError errors_models.IceCreamErrors
				iceCreamError.Error = fmt.Sprintf("could not find matching completedTask ID: %s", oid)
				iceCreamError.Info = "Error deleting rule: no matching oid"
				json.NewEncoder(w).Encode(iceCreamError)
			}
		}
	}
}
