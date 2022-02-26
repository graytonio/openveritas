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
	case http.MethodPut:
		nodePropertyPutHandler(rw, r)
	case http.MethodDelete:
		nodePropertyDeleteHandler(rw, r)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Get Property of Node
func nodePropertyGetHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	node_name := vars["node"]
	prop_name := vars["prop"]

	node, err := models.GetNode(node_name)
	if handleMongoError(err, rw) || handleNotFoundError(node, rw) {
		return
	}

	if prop_name == "" {
		properties, err := models.GetAllPropertiesOfNode(node)
		if handleMongoError(err, rw) || handleNotFoundError(node, rw) {
			return
		}
		response, _ := json.Marshal(properties)
		rw.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(rw, "%s", string(response))
	} else {
		property, err := models.GetProperty(node, prop_name)
		if handleMongoError(err, rw) || handleNotFoundError(node, rw) {
			return
		}

		response, _ := json.Marshal(property)
		rw.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(rw, "%s", string(response))
	}
}

// Update/Create Property of Node
func nodePropertyPutHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	node_name := vars["node"]
	prop_name := vars["prop"]

	var body models.PropertyForm
	err := json.NewDecoder(r.Body).Decode(&body)
	if handleBodyParseError(err, rw) {
		return
	}

	node, err := models.GetNode(node_name)
	if handleMongoError(err, rw) || handleNotFoundError(node, rw) {
		return
	}

	property, err := models.GetProperty(node, prop_name)
	if handleMongoError(err, rw) {
		return
	}

	if property == nil {
		property = models.NewProperty(node, prop_name, body.PropertyValue)
	} else {
		property.PropertyName = body.PropertyName
		property.PropertyValue = body.PropertyValue
	}

	_, err = models.UpdateOrCreateProperty(property)
	if handleMongoError(err, rw) {
		return
	}

	rw.WriteHeader(http.StatusOK)
}

// Delete Property of Node
func nodePropertyDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	node_name := vars["node"]
	prop_name := vars["prop"]

	node, err := models.GetNode(node_name)
	if handleMongoError(err, rw) || handleNotFoundError(node, rw) {
		return
	}

	property, err := models.GetProperty(node, prop_name)
	if handleMongoError(err, rw) || handleNotFoundError(property, rw) {
		return
	}

	err = models.DeleteProperty(property)
	if handleMongoError(err, rw) {
		return
	}

	rw.WriteHeader(http.StatusOK)
}
