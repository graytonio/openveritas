package routes

import (
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

// Prase Property Query
// TODO Implement Wildcards
func propertyGetHander(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prop_name := vars["prop_name"]

	props, err := models.GetAllProperties(prop_name)
	if isMongoError(err) {
		sendDBError(rw, err)
		return
	} else if isNotFoundError(props) {
		sendError(rw, http.StatusNotFound, "Property Not Found")
		return
	}

	sendJSONData(rw, props)
}
