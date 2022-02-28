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
	node_name := vars["node_name"]
	prop_name := vars["prop_name"]
	var data interface{}
	var err error

	node, err := models.GetNode(node_name)
	if isMongoError(err) {
		sendDBError(rw, err)
		return
	} else if isNotFoundError(node) {
		sendError(rw, http.StatusNotFound, "Node not Found")
		return
	}

	if prop_name == "" {
		data, err = models.GetAllPropertiesOfNode(node)
	} else {
		data, err = models.GetProperty(node, prop_name)
	}

	if isMongoError(err) {
		sendDBError(rw, err)
		return
	} else if isNotFoundError(data) {
		sendError(rw, http.StatusNotFound, "Property Not Found")
	}

	sendJSONData(rw, data)
}

// Update/Create Property of Node
func nodePropertyPutHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	node_name := vars["node"]
	prop_name := vars["prop"]

	var body models.PropertyForm
	err := json.NewDecoder(r.Body).Decode(&body)
	if isError(err) {
		sendError(rw, http.StatusBadRequest, fmt.Sprintf("Error parsing request body: %s", err.Error()))
		return
	}

	node, err := models.GetNode(node_name)
	if isMongoError(err) {
		sendDBError(rw, err)
		return
	} else if isNotFoundError(err) {
		sendError(rw, http.StatusNotFound, "Node Not Found")
		return
	}

	property, err := models.GetProperty(node, prop_name)
	if isMongoError(err) {
		sendDBError(rw, err)
		return
	}

	if isNotFoundError(property) {
		property = models.NewProperty(node, prop_name, body.PropertyValue)
	} else {
		property.PropertyName = body.PropertyName
		property.PropertyValue = body.PropertyValue
	}

	_, err = models.UpdateOrCreateProperty(property)
	if isMongoError(err) {
		sendDBError(rw, err)
		return
	}
}

// Delete Property of Node
func nodePropertyDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	node_name := vars["node"]
	prop_name := vars["prop"]

	node, err := models.GetNode(node_name)
	if isMongoError(err) {
		sendDBError(rw, err)
		return
	} else if isNotFoundError(node) {
		sendError(rw, http.StatusNotFound, "Node Not Fount")
		return
	}

	property, err := models.GetProperty(node, prop_name)
	if isMongoError(err) {
		sendDBError(rw, err)
		return
	} else if isNotFoundError(property) {
		sendError(rw, http.StatusNotFound, "Property Not Found")
		return
	}

	err = models.DeleteProperty(property)
	if isMongoError(err) {
		sendDBError(rw, err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}
