package completed_tasks_router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	completed_tasks_controllers "github.com/ice-cream-backend/controllers/v1/completed_tasks"
	completed_tasks_models "github.com/ice-cream-backend/models/v1/completed_tasks"
	errors_models "github.com/ice-cream-backend/models/v1/errors"
	utils_models "github.com/ice-cream-backend/models/v1/utils"
	"github.com/ice-cream-backend/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCompletedTasks(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint hit: create completedTasks")
	utils.EnableCors(&w)

	if r.Method == "POST" {
		var completedTasksPost completed_tasks_models.CompletedTasks
		_ = json.NewDecoder(r.Body).Decode(&completedTasksPost)
	
		res, err := completed_tasks_controllers.CreateCompletedTask(completedTasksPost)
	
		if err != nil {
			fmt.Fprintf(w, "Error creating completedTasks!")
		} else {
			newCompletedTask := completed_tasks_controllers.GetCompletedTasksByCompletedTaskId(res.InsertedID)
	
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newCompletedTask)
		}
	}
}

func DeleteCompletedTaskByCompletedTaskId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: delete garden by id")
	vars := mux.Vars(r)
	utils.EnableCors(&w)

	paramsCompletedTaskId := vars["completedTaskId"]

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
			var deleteResult utils_models.DeleteResult
			deleteResult.Info = "Successfully deleted Completed Task"
			deleteResult.Success = true
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(deleteResult)
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