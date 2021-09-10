package completed_tasks_router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	completed_tasks_controllers "github.com/ice-cream-backend/controllers/v1/completed_tasks"
	users_controllers "github.com/ice-cream-backend/controllers/v1/users"
	completed_tasks_models "github.com/ice-cream-backend/models/v1/completed_tasks"
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
			err := fmt.Errorf("already has existing completedTask for this rule: %q and date: %q", completedTasksPost.RuleId.String(), completedTasksPost.Date)
			utils.SendErrorBack(w, err, "Already has existing completedTask for this rule and date!")
		} else {
			res, err := completed_tasks_controllers.CreateCompletedTask(completedTasksPost)

			if err != nil {
				utils.SendErrorBack(w, err, "Error Error creating completedTasks!")
			} else {
				_ = completed_tasks_controllers.GetCompletedTasksByCompletedTaskId(res.InsertedID)

				user, err := users_controllers.GetUserByFireBaseUserId(completedTasksPost.FireBaseUserId)

				if err != nil {
					utils.SendErrorBack(w, err, "Error finding user to update coins after completing task")
				} else {
					user.Balance = user.Balance + COIN_AFTER_COMPLETED_TASK

					updatedUser, err := users_controllers.UpdateUserByUserId(user.ID, user)

					if err != nil {
						utils.SendErrorBack(w, err, "Error updating user's coins after completing task")
					} else {
						if updatedUser.MatchedCount != 0 {
							userData, err := users_controllers.GetUserByFireBaseUserId(user.FireBaseUserId)

							if err != nil {
								utils.SendErrorBack(w, err, "Error getting updated user data")
							} else {
								utils.SendResponseBack(w, userData, http.StatusOK)
							}
						} else {
							err := fmt.Errorf("could not find matching user ID: %v", user.ID)
							utils.SendErrorBack(w, err, "Error updating user's balance")
						}
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
			utils.SendErrorBack(w, err, "Invalid completedTaskId provided")
		} else {
			res, err := completed_tasks_controllers.DeleteCompletedTaskByCompletedTaskId(oid)

			if err != nil {
				utils.SendErrorBack(w, err, "Error deleting completedTask")
			}

			if res.DeletedCount == 1 {
				user, err := users_controllers.GetUserByFireBaseUserId(paramsFireBaseUserId)

				if err != nil {
					utils.SendErrorBack(w, err, "Error finding user to update coins after deleting completed task")
				} else {
					user.Balance = user.Balance - COIN_AFTER_COMPLETED_TASK

					updatedUser, err := users_controllers.UpdateUserByUserId(user.ID, user)

					if err != nil {
						utils.SendErrorBack(w, err, "Error updating user's coins after deleting completed task")
					} else {
						if updatedUser.MatchedCount != 0 {
							userData, err := users_controllers.GetUserByFireBaseUserId(user.FireBaseUserId)

							if err != nil {
								utils.SendErrorBack(w, err, "Error getting updated user data after deleting completedTask")
							} else {
								utils.SendResponseBack(w, userData, http.StatusOK)
							}
						} else {
							err := fmt.Errorf("could not find matching user ID: %v", user.ID)
							utils.SendErrorBack(w, err, "Error updating user's balance")
						}
					}
				}
			} else {
				err := fmt.Errorf("could not find matching completedTask ID: %s", oid)
				utils.SendErrorBack(w, err, "Error deleting rule: no matching oid")
			}
		}
	}
}
