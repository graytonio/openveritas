package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/graytonio/openveritas/server/models"
)

func NodeHandler(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			nodeGetHandler(rw, r)
			return
		case http.MethodPost:
			nodePostHandler(rw, r)
			return
		case http.MethodPut:
			nodePutHandler(rw, r)
			return
		case http.MethodDelete:
			nodeDeleteHandler(rw, r)
			return
		default:
			rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func nodeGetHandler(rw http.ResponseWriter, r *http.Request) {
	node_name := mux.Vars(r)["node"]
	if node_name == "" {
		response, _ := json.Marshal(models.GetAllNodes())
		rw.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(rw, "%s", string(response))
	} else {
		node := models.GetNode(node_name)
		if node == nil {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		respoonse, _ := json.Marshal(node)
		rw.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(rw, "%s", string(respoonse))
	}
}

func nodePostHandler(rw http.ResponseWriter, r *http.Request) {
	var body models.NewNodeForm
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = models.CreateNode(body.Name)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key error collection: veritas.nodes index: name_1") {
			rw.WriteHeader(http.StatusConflict)
			return
		}
		
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	
	rw.WriteHeader(http.StatusCreated)
}

func nodePutHandler(rw http.ResponseWriter, r *http.Request) {
	node_name := mux.Vars(r)["node"]
	node := models.GetNode(node_name)
	if node == nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	
	var body models.NewNodeForm
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	if body.Name != "" { node.Name = body.Name }

	err = models.UpdateNode(node)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func nodeDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	node_name := mux.Vars(r)["node"]
	node := models.GetNode(node_name)
	if node == nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	err := models.DeleteNode(node)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}