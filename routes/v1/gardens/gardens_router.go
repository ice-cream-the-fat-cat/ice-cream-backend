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
	utils_models "github.com/ice-cream-backend/models/v1/utils"
	"github.com/ice-cream-backend/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateGardens(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: create gardens for method:", r.Method)
	utils.EnableCors(&w)
	if r.Method == "POST" {
		start := utils.StartPerformanceTest()

		var createdGardensPost gardens_models.Gardens
		_ = json.NewDecoder(r.Body).Decode(&createdGardensPost)

		res, err := gardens_controllers.CreateGardens(createdGardensPost)

		if err != nil {
			fmt.Fprintf(w, "Error creating garden!")
		} else {
			newGarden, err := gardens_controllers.GetGardensByGardenId(res.InsertedID)

			if err != nil {
				var iceCreamError errors_models.IceCreamErrors
				iceCreamError.Error = err.Error()
				iceCreamError.Info = "Invalid gardenId provided"
				json.NewEncoder(w).Encode(iceCreamError)
			} else {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(newGarden)
				utils.StopPerformanceTest(start, "Successful create garden took")
			}
		}
	}
}

func GetGardenByGardenId(w http.ResponseWriter, r *http.Request) {
	start := utils.StartPerformanceTest()
	vars := mux.Vars(r)
	utils.EnableCors(&w)

	paramsGardenId := vars["gardenId"]

	oid, err := primitive.ObjectIDFromHex(paramsGardenId)

	if err != nil {
		log.Println("Error converting params gardenId to ObjectId")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Invalid gardenId provided")
	} else {
		populatedGarden, err := gardens_controllers.GetPopulatedGardenByGardenId(oid)

		if err != nil {
			var iceCreamError errors_models.IceCreamErrors
			iceCreamError.Error = err.Error()
			iceCreamError.Info = "Invalid gardenId provided"
			json.NewEncoder(w).Encode(iceCreamError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(populatedGarden)
			utils.StopPerformanceTest(start, fmt.Sprintf("Successfully got fully populated garden for gardenId %s ", paramsGardenId))
		}
	}
}

func GetGardensByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	utils.EnableCors(&w)
	paramsUserId := vars["fireBaseUserId"]

	userGardens := gardens_controllers.GetGardensByUserId(paramsUserId)

	if len(userGardens) == 0 {
		userGardens = []gardens_models.Gardens{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userGardens)
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
			updatedGarden, _ := gardens_controllers.GetGardensByGardenId(oid)

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

func DeleteGardenByGardenId(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint hit: delete garden by id")
	vars := mux.Vars(r)
	utils.EnableCors(&w)

	if r.Method == "Delete" {
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
			res, err := gardens_controllers.DeleteGardenByGardenId(oid)

			if err != nil {
				log.Println("Error deleting garden:", err)
				w.Header().Set("Content-Type", "application/json")
				var iceCreamError errors_models.IceCreamErrors
				iceCreamError.Error = err.Error()
				iceCreamError.Info = "Error deleting garden"
				json.NewEncoder(w).Encode(iceCreamError)
			}

			if res.DeletedCount == 1 {
				var deleteResult utils_models.DeleteResult
				deleteResult.Info = "Successfully deleted Garden"
				deleteResult.Success = true
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(deleteResult)
			} else {
				log.Println("could not find matching garden ID:", oid)
				w.Header().Set("Content-Type", "application/json")
				var iceCreamError errors_models.IceCreamErrors
				iceCreamError.Error = fmt.Sprintf("could not find matching garden ID: %s", oid)
				iceCreamError.Info = "Error deleting rule: no matching oid"
				json.NewEncoder(w).Encode(iceCreamError)
			}
		}
	}
}