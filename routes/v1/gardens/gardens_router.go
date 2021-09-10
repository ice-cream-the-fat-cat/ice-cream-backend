package gardens_router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	gardens_controllers "github.com/ice-cream-backend/controllers/v1/gardens"
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

		var newGarden gardens_models.GardenForMongo
		_ = json.NewDecoder(r.Body).Decode(&newGarden)

		if gardens_models.GardenValidation(newGarden) {
			res, err := gardens_controllers.CreateGardens(newGarden)

			if err != nil {
				utils.SendErrorBack(w, err, "Error creating garden!")
			} else {
				newGarden, err := gardens_controllers.GetGardensByGardenId(res.InsertedID)

				if err != nil {
					utils.SendErrorBack(w, err, "Invalid gardenId so could not get garden")
				} else {
					utils.SendResponseBack(w, newGarden, http.StatusCreated)
					utils.StopPerformanceTest(start, "Successful create garden took")
				}
			}
		} else {
			utils.SendErrorBack(w, fmt.Errorf("missiing required fields for creating garden: %+v", newGarden), "Missing required fields to create garden")
		}
	}
}

func GetGardenByGardenId(w http.ResponseWriter, r *http.Request) {
	start := utils.StartPerformanceTest()
	vars := mux.Vars(r)
	utils.EnableCors(&w)

	paramsGardenId := vars["gardenId"]
	paramsDate := vars["date"]

	oid, err := primitive.ObjectIDFromHex(paramsGardenId)

	if err != nil {
		utils.SendErrorBack(w, err, "Invalid gardenId provided")
	} else {
		populatedGarden, err := gardens_controllers.GetPopulatedGardenByGardenId(oid, paramsDate)

		if err != nil {
			utils.SendErrorBack(w, err, "Could not get get garden with provided gardenId")
		} else {
			utils.SendResponseBack(w, populatedGarden, http.StatusOK)
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

	utils.SendResponseBack(w, userGardens, http.StatusOK)
}

func GetGardenByGardenIdWithStartAndEndDate(w http.ResponseWriter, r *http.Request) {
	start := utils.StartPerformanceTest()
	vars := mux.Vars(r)
	utils.EnableCors(&w)

	paramsGardenId := vars["gardenId"]
	paramsStartDate := vars["startDate"]
	paramsEndDate := vars["endDate"]

	oid, err := primitive.ObjectIDFromHex(paramsGardenId)

	if err != nil {

	} else {
		populatedGarden, err := gardens_controllers.GetPopulatedGardenByGardenIdWithStartAndEndDate(oid, paramsStartDate, paramsEndDate)

		if err != nil {
			utils.SendErrorBack(w, err, "Invalid gardenId provided for start / end date getGarden")
		} else {
			utils.SendResponseBack(w, populatedGarden, http.StatusOK)
			utils.StopPerformanceTest(start, fmt.Sprintf("Successfully got fully populated garden (start / end date) for gardenId %s ", paramsGardenId))
		}
	}
}

func UpdateGardenById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: update garden by id")
	vars := mux.Vars(r)
	utils.EnableCors(&w)
	if r.Method == "PUT" {
		var garden gardens_models.Gardens
		_ = json.NewDecoder(r.Body).Decode(&garden)

		paramsGardenId := vars["gardenId"]

		oid, err := primitive.ObjectIDFromHex(paramsGardenId)

		if err != nil {
			utils.SendErrorBack(w, err, "Error converting params gardenId to ObjectId")
		} else {
			res, err := gardens_controllers.UpdateGardenByGardenId(oid, garden)

			if err != nil {
				utils.SendErrorBack(w, err, "Error updating garden")
			}

			if res.MatchedCount != 0 {
				updatedGarden, _ := gardens_controllers.GetGardensByGardenId(oid)

				utils.SendResponseBack(w, updatedGarden, http.StatusOK)
			} else {
				utils.SendErrorBack(w, fmt.Errorf("could not find matching garden ID: %s", oid), "Error updating rule: no matching oid")
			}
		}
	}
}

func DeleteGardenByGardenId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: delete garden by id")
	vars := mux.Vars(r)
	utils.EnableCors(&w)

	if r.Method == "Delete" {
		paramsGardenId := vars["gardenId"]

		oid, err := primitive.ObjectIDFromHex(paramsGardenId)

		if err != nil {
			utils.SendErrorBack(w, err, "Error converting params gardenId to ObjectId")
		} else {
			res, err := gardens_controllers.DeleteGardenByGardenId(oid)

			if err != nil {
				utils.SendErrorBack(w, err, "Error deleting garden")
			}

			if res.DeletedCount == 1 {
				var deleteResult utils_models.DeleteResult
				deleteResult.Info = "Successfully deleted Garden"
				deleteResult.Success = true
				utils.SendResponseBack(w, deleteResult, http.StatusOK)
			} else {
				utils.SendErrorBack(w, fmt.Errorf("could not find matching garden ID: %s", oid), "Error deleting rule: no matching oid")
			}
		}
	}
}
