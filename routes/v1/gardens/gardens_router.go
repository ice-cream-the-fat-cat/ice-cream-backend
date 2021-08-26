package gardens_router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	gardens_controllers "github.com/ice-cream-backend/controllers/v1/gardens"
	gardens_models "github.com/ice-cream-backend/models/v1/gardens"
	"github.com/ice-cream-backend/utils"
)

func CreateGardens(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: create gardens")
	utils.EnableCors(&w)

	var createdGardensPost gardens_models.Gardens
	_ = json.NewDecoder(r.Body).Decode(&createdGardensPost)

	res, err := gardens_controllers.CreateGardens(createdGardensPost)

	if err != nil {
		fmt.Fprintf(w, "Error creating garden!")
	} else {
		newCompletedTask := gardens_controllers.GetGardensById(res.InsertedID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newCompletedTask)
	}
}

func GetGardenByGardenId(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	utils.EnableCors(&w)

	gardenId := vars["gardenId"]
	log.Println("gardenId:", gardenId)
}