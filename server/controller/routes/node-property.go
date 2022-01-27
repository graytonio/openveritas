package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graytonio/openveritas/server/models"
)

func PropertyHandler(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			nodePropertyGetHandler(rw, r)
		case http.MethodPost:
			nodePropertyPostHandler(rw, r)
		case http.MethodPut:
			nodePropertyPutHandler(rw, r)
		case http.MethodDelete:
			nodePropertyDeleteHandler(rw, r)
		default:
			rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func nodePropertyGetHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	node_name := vars["node"]
	prop_name := vars["prop"]

	node, err := models.GetNode(node_name)
	if checkRouteError(err, rw) { return }
	if checkNotFound(node, rw) { return }

	if prop_name == "" {
		properties := models.GetAllPropertiesOfNode(node)
		respoonse, _ := json.Marshal(properties)
		rw.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(rw, "%s", string(respoonse))
	} else {
		property, err := models.GetProperty(node, prop_name)
		if checkRouteError(err, rw) { return }
		if checkNotFound(property, rw) { return }

		response, _ := json.Marshal(property)
		rw.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(rw, "%s", string(response))
	}
}

func nodePropertyPostHandler(rw http.ResponseWriter, r *http.Request) {
	node_name := mux.Vars(r)["node"]
	var body models.NewPropertyForm
	err := json.NewDecoder(r.Body).Decode(&body)
	if checkRouteError(err, rw) { return }

	property, err := models.CreateProperty(node_name, body.PropertyName, body.PropertyValue)
	if checkRouteError(err, rw) { return }
	if checkNotFound(property, rw) { return }

	rw.WriteHeader(http.StatusCreated)
}

func nodePropertyPutHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	node_name := vars["node"]
	prop_name := vars["prop"]

	node, err := models.GetNode(node_name)
	if checkRouteError(err, rw) { return }
	if checkNotFound(node, rw) { return }

	var body models.UpdatePropertyForm
	err = json.NewDecoder(r.Body).Decode(&body)
	if checkRouteError(err, rw) { return }

	property, err := models.GetProperty(node, prop_name)
	if checkRouteError(err, rw) { return }
	if checkNotFound(property, rw) {
		property, err = models.CreateProperty(node.Name, prop_name, body.PropertyValue)
		if checkRouteError(err, rw) { return }
	} else {
		property.PropertyValue = body.PropertyValue
	}
	
	_, err = models.UpdateProperty(property)
	if checkRouteError(err, rw) { return }

	rw.WriteHeader(http.StatusOK)
}

//Delete Property
func nodePropertyDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	node_name := vars["node"]
	prop_name := vars["prop"]
	
	node, err := models.GetNode(node_name)
	if checkRouteError(err, rw) { return }
	if checkNotFound(node, rw) { return }

	property, err := models.GetProperty(node, prop_name)
	if checkRouteError(err, rw) { return }
	if checkNotFound(property, rw) { return }

	err = models.DeleteProperty(property)
	if checkRouteError(err, rw) { return }
	
	rw.WriteHeader(http.StatusOK)
}