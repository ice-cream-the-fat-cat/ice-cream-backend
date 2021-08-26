package completedTasks_router

import (
	"encoding/json"
	"fmt"
	"net/http"

	completedTasks_models "github.com/ice-cream-backend/models/v1/completedTasks"
	"github.com/ice-cream-backend/utils"
)

func CreateCompletedTasks(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint hit: create completedTasks")
	utils.EnableCors(&w)

	var completedTasksPost completedTasks_models.CompletedTasks
	_ = json.NewDecoder(r.Body).Decode(&completedTasksPost)

}