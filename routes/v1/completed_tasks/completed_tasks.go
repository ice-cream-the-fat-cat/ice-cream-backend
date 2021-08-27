package completed_tasks_router

import (
	"encoding/json"
	"fmt"
	"net/http"

	completed_tasks_controllers "github.com/ice-cream-backend/controllers/v1/completed_tasks"
	completed_tasks_models "github.com/ice-cream-backend/models/v1/completed_tasks"
	"github.com/ice-cream-backend/utils"
)

func CreateCompletedTasks(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint hit: create completedTasks")
	utils.EnableCors(&w)

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