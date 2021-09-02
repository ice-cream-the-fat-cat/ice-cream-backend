package gardens_router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	gardens_controllers "github.com/ice-cream-backend/controllers/v1/gardens"
	errors_models "github.com/ice-cream-backend/models/v1/errors"
	gardens_models "github.com/ice-cream-backend/models/v1/gardens"
	"github.com/ice-cream-backend/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateGardens(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: create gardens for method:", r.Method)
	utils.EnableCors(&w)
	if r.Method == "POST" {

		var createdGardensPost gardens_models.Gardens
		_ = json.NewDecoder(r.Body).Decode(&createdGardensPost)

		res, err := gardens_controllers.CreateGardens(createdGardensPost)

		if err != nil {
			fmt.Fprintf(w, "Error creating garden!")
		} else {
			newGarden := gardens_controllers.GetGardensByGardenId(res.InsertedID)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newGarden)
		}
	}
}

func GetGardenByGardenId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	utils.EnableCors(&w)

	paramsGardenId := vars["gardenId"]

	oid, err := primitive.ObjectIDFromHex(paramsGardenId)

	if err != nil {
		log.Println("Error converting params gardenId to ObjectId")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Invalid gardenId provided")
	} else {
		populatedGarden := gardens_controllers.GetPopulatedGardenByGardenId(oid)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(populatedGarden)
	}
}

func UpdateGardenById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: update garden by id")
	vars := mux.Vars(r)
	utils.EnableCors(&w)

	var garden gardens_models.Gardens
	_ = json.NewDecoder(r.Body).Decode(&garden)

	paramsGardenId := vars["gardenId"]

	oid, err := primitive.ObjectIDFromHex(paramsGardenId)

	if err != nil {
		log.Println("Error converting params gardenId to ObjectId:", err)
		w.Header().Set("Content-Type", "application/json")
		var iceCreamError errors_models.IceCreamErrors
		iceCreamError.Error = err.Error()
		iceCreamError.Info = "Invalid gardenId provided"
		json.NewEncoder(w).Encode(iceCreamError)
	} else {
		res, err := gardens_controllers.UpdateGardenByGardenId(oid, garden)

		if err != nil {
			log.Println("Error updating garden:", err)
			w.Header().Set("Content-Type", "application/json")
			var iceCreamError errors_models.IceCreamErrors
			iceCreamError.Error = err.Error()
			iceCreamError.Info = "Error updating garden"
			json.NewEncoder(w).Encode(iceCreamError)
		}

		if res.MatchedCount != 0 {
			updatedGarden := gardens_controllers.GetGardensByGardenId(oid)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedGarden)
		} else {
			log.Println("could not find matching garden ID:", oid)
			w.Header().Set("Content-Type", "application/json")
			var iceCreamError errors_models.IceCreamErrors
			iceCreamError.Error = fmt.Sprintf("could not find matching garden ID: %s", oid)
			iceCreamError.Info = "Error updating rule: no matching oid"
			json.NewEncoder(w).Encode(iceCreamError)
		}
	}
}
func GetGardensByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	utils.EnableCors(&w)
	paramsUserId := vars["userFireBaseId"]

	userGardens := gardens_controllers.GetGardensByUserId(paramsUserId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userGardens)

}
