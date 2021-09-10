package garden_categories_router

import (
	"fmt"
	"net/http"

	garden_categories_controllers "github.com/ice-cream-backend/controllers/v1/garden_categories"
	"github.com/ice-cream-backend/utils"
)

func GetGardenCategories(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: get gardenCategories")
	utils.EnableCors(&w)

	gardenCategories, err := garden_categories_controllers.GetGardenCategories()

	if err != nil {
		utils.SendErrorBack(w, err, "Could not get gardenCategories")
	} else {
		utils.SendResponseBack(w, gardenCategories, http.StatusOK)
	}
}