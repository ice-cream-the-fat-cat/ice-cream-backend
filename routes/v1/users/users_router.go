package users_router

import (
	"encoding/json"
	"net/http"

	users_controllers "github.com/ice-cream-backend/controllers/v1/users"

	"github.com/gorilla/mux"
	"github.com/ice-cream-backend/utils"
)

func GetUserByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	utils.EnableCors(&w)
	paramsUserId := vars["userFireBaseId"]

	userGardens := users_controllers.GetUserByUserId(paramsUserId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userGardens)
}
