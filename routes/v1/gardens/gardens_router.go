package gardens_router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Gardens(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Gardens API  version:%v  id:%v\n", vars["version"], vars["id"])
	fmt.Println("Endpoint hit: gardens router")
}
