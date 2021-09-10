package users_router

import (
	"log"
	"net/http"

	users_controllers "github.com/ice-cream-backend/controllers/v1/users"
	users_models "github.com/ice-cream-backend/models/v1/users"

	"github.com/gorilla/mux"
	"github.com/ice-cream-backend/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	utils.EnableCors(&w)
	paramsUserId := vars["fireBaseUserId"]

	log.Println(paramsUserId)

	_, err := users_controllers.GetUserByFireBaseUserId(paramsUserId)
	if err != nil {
		log.Println(w, "Error no user data!")
		var createdUserPost users_models.Users
		createdUserPost.ID = primitive.NewObjectID()
		createdUserPost.FireBaseUserId = paramsUserId
		createdUserPost.Balance = 0
		createdUserPost.FlowerCollections = []primitive.ObjectID{}

		_, err := users_controllers.CreateUser(createdUserPost)
		if err != nil {
			utils.SendErrorBack(w, err, "Error creating user!")
		}

	}

	userData, err := users_controllers.GetUserByFireBaseUserId(paramsUserId)
	if err != nil {
		log.Println(w, "Error no user data!")
		utils.SendErrorBack(w, err, "Error no user data after user creation was done!")
	}

	utils.SendResponseBack(w, userData, http.StatusOK)
}
