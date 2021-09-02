package users_router

import (
	"encoding/json"
	"fmt"
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
	paramsUserId := vars["userFireBaseId"]

	_, err := users_controllers.GetUserByUserId(paramsUserId)

	if err != nil {
		fmt.Fprintf(w, "Error no user data!")
		var createdUserPost users_models.Users
		createdUserPost.ID = primitive.NewObjectID()
		createdUserPost.UserFireBaseId = paramsUserId
		createdUserPost.NumCoins = 0

		newUser, err := users_controllers.CreateUser(createdUserPost)
		if err != nil {
			fmt.Fprintf(w, "Error creating user!")
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newUser)
		}
	}
}
