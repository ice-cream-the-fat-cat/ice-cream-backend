package gardens_router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	gardens_controller "github.com/ice-cream-backend/controller/v1/gardens"
)

func GardensIdGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "GardensIdGet  version:%v  id:%v\n", vars["version"], vars["id"])
	fmt.Println("Endpoint hit: gardens router")

	integerId, _ := strconv.Atoi(vars["id"])

	gardenData := gardens_controller.GetGardenData(integerId)
	fmt.Println(gardenData)
}

func GardensUserIdGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "GardensUserIdGet  version:%v  userid:%v\n", vars["version"], vars["userid"])
	fmt.Println("Endpoint hit: gardens router")
}

func GardensPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "GardensPost  version:%v\n", vars["version"])
	fmt.Println("Endpoint hit: gardens GardensPost")
}
