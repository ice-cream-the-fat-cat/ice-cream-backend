package flowersStore_router

import (
	"encoding/json"
	"log"
	"net/http"

	flowersStore_controllers "github.com/ice-cream-backend/controllers/v1/flowersStore"
	users_controllers "github.com/ice-cream-backend/controllers/v1/users"
	flowersStore_models "github.com/ice-cream-backend/models/v1/flowersStore"
	"github.com/ice-cream-backend/utils"
)

func BuyNewFlower(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	log.Println("Request came into buy new flower endpoint", r.Method)

	if r.Method == "PUT" {
		var flowersStore flowersStore_models.FlowersStore
		_ = json.NewDecoder(r.Body).Decode(&flowersStore)

		_, err := flowersStore_controllers.BuyNewFlower(flowersStore)

		if err != nil {
			utils.SendErrorBack(w, err, "Invalid flowersStore request provided")
		} else {
			userData, err := users_controllers.GetUserByFireBaseUserId(flowersStore.FireBaseUserId)

			if err != nil {
				utils.SendErrorBack(w, err, "Error no user data!")
				return
			}
			
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(userData)
		}
	}
}
