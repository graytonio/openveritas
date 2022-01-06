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
			propertyGetHandler(rw, r)
			return
		case http.MethodPost:
			propertyPostHandler(rw, r)
			return
		case http.MethodPut:
			propertyPutHandler(rw, r)
			return
		case http.MethodDelete:
			propertyDeleteHandler(rw, r)
			return
		default:
			rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func propertyGetHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	node_name := vars["node"]
	prop_name := vars["prop"]

	node := models.GetNode(node_name)
	if node == nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	if prop_name == "" {
		properties := models.GetAllProperties(node)
		respoonse, _ := json.Marshal(properties)
		rw.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(rw, "%s", string(respoonse))
	} else {
		property := models.GetProperty(node, prop_name)
		if property == nil {
			rw.WriteHeader(http.StatusNotFound)
			return 
		}

		response, _ := json.Marshal(property)
		rw.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(rw, "%s", string(response))
	}
}

func propertyPostHandler(rw http.ResponseWriter, r *http.Request) {
	node_name := mux.Vars(r)["node"]
	var body models.NewPropertyForm
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = models.CreateProperty(node_name, body.PropertyName, body.PropertyValue)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusAccepted)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

func propertyPutHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	node_name := vars["node"]
	prop_name := vars["prop"]

	node := models.GetNode(node_name)
	if node == nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	property := models.GetProperty(node, prop_name)
	if property == nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	var body models.UpdatePropertyForm
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	property.PropertyValue = body.PropertyValue
	
	err = models.UpdateProperty(property)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

//Delete Property
func propertyDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	node_name := vars["node"]
	prop_name := vars["prop"]
	
	node := models.GetNode(node_name)
	if node == nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	property := models.GetProperty(node, prop_name)
	if property == nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	err := models.DeleteProperty(property)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	
	rw.WriteHeader(http.StatusOK)
}