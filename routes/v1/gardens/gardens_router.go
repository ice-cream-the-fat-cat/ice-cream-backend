package gardens_router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GardensIdGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "GardensIdGet  version:%v  id:%v\n", vars["version"], vars["id"])
	fmt.Println("Endpoint hit: gardens router")
}

func GardensUserIdGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "GardensUserIdGet  version:%v  userid:%v\n", vars["version"], vars["userid"])
	fmt.Println("Endpoint hit: gardens router")
}
