package router

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func testRoute() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Test")
	}
}
