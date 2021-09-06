package flowersStore_router

import (
	"encoding/json"
	"log"
	"net/http"

	flowersStore_controllers "github.com/ice-cream-backend/controllers/v1/flowersStore"
	users_controllers "github.com/ice-cream-backend/controllers/v1/users"
	errors_models "github.com/ice-cream-backend/models/v1/errors"
	flowersStore_models "github.com/ice-cream-backend/models/v1/flowersStore"
	"github.com/ice-cream-backend/utils"
)

func BuyNewFlower(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)

	var flowersStore flowersStore_models.FlowersStore
	_ = json.NewDecoder(r.Body).Decode(&flowersStore)

	_, err := flowersStore_controllers.BuyNewFlower(flowersStore)
	if err != nil {
		var iceCreamError errors_models.IceCreamErrors
		iceCreamError.Error = err.Error()
		iceCreamError.Info = "Invalid flowersStore request provided"
		json.NewEncoder(w).Encode(iceCreamError)
	} else {
		userData, err := users_controllers.GetUserByFireBaseUserId(flowersStore.FireBaseUserId)
		if err != nil {
			log.Println(w, "Error no user data!")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userData)
	}
}
