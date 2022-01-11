package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graytonio/openveritas/server/models"
)

func PropertyQueryHandler(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		propertyGetHander(rw, r)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func propertyGetHander(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prop_name := vars["prop"]

	props := models.GetAllProperties(prop_name)
	if props == nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	response, _ := json.Marshal(props)
	rw.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(rw, "%s", string(response))
}