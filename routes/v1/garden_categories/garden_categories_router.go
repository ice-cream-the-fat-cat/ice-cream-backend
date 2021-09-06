package garden_categories_router

import (
	"encoding/json"
	"fmt"
	"net/http"

	garden_categories_controllers "github.com/ice-cream-backend/controllers/v1/garden_categories"
	errors_models "github.com/ice-cream-backend/models/v1/errors"
	"github.com/ice-cream-backend/utils"
)

func GetGardenCategories(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: get gardenCategories")
	utils.EnableCors(&w)

	gardenCategories, err := garden_categories_controllers.GetGardenCategories()

	if err != nil {
		var iceCreamError errors_models.IceCreamErrors
		iceCreamError.Error = err.Error()
		iceCreamError.Info = "Could not get gardenCategories"
		json.NewEncoder(w).Encode(iceCreamError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gardenCategories)
	}
}