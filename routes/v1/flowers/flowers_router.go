package flowers_router

import (
	"net/http"

	flowers_controllers "github.com/ice-cream-backend/controllers/v1/flowers"

	"github.com/ice-cream-backend/utils"
)

func GetFlowers(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)

	flowerList, err := flowers_controllers.GetFlowers()

	if err != nil {
		utils.SendErrorBack(w, err, "Error getting all flowers")
	}

	utils.SendResponseBack(w, flowerList, http.StatusOK)
}
