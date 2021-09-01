package flowers_router

import (
	"encoding/json"
	"net/http"

	flowers_controllers "github.com/ice-cream-backend/controllers/v1/flowers"

	"github.com/ice-cream-backend/utils"
)

func GetFlower(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)

	flowerList := flowers_controllers.GetFlower()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flowerList)

}
