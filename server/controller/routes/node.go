package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graytonio/openveritas/server/models"
)

func NodeHandler(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		nodeGetHandler(rw, r)
	case http.MethodPut:
		nodePutHandler(rw, r)
	case http.MethodDelete:
		nodeDeleteHandler(rw, r)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Get All/A Node
func nodeGetHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	node_name := vars["node"]
	var data interface{}
	var err error

	if node_name == "" {
		data, err = models.GetAllNodes()
	} else {
		data, err = models.GetNode(node_name)
	}

	if isMongoError(err) {
		sendDBError(rw, err)
		return
	} else if isNotFoundError(err) {
		sendError(rw, http.StatusNotFound, "Node Not Found")
		return
	}

	sendJSONData(rw, data)
}

// Update/Create Node
func nodePutHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	node_name := vars["node_name"]

	if node_name == "" {
		sendError(rw, http.StatusBadRequest, "node_name is required parameter")
		return
	}

	var body models.NewNodeForm
	err := json.NewDecoder(r.Body).Decode(&body)
	if isError(err) {
		sendError(rw, http.StatusBadRequest, fmt.Sprintf("Error parsing request body: %s", err.Error()))
		return
	}

	node, err := models.GetNode(node_name)
	if isMongoError(err,) {
		sendDBError(rw, err)
		return
	}

	if isNotFoundError(node) {
		node = models.NewNode(node_name)
	} else {
		node.Name = body.Name
	}

	_, err = models.UpdateOrCreateNode(node)
	if isMongoError(err) {
		sendDBError(rw, err)
		return
	}
}

// Delete Node
func nodeDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	node_name := vars["node_name"]
	node, err := models.GetNode(node_name)
	if isMongoError(err) {
		sendDBError(rw, err)
		return
	} else if isNotFoundError(node) {
		sendError(rw, http.StatusNotFound, "Node not Found")
		return
	}

	err = models.DeleteNode(node)
	if isMongoError(err) {
		sendDBError(rw, err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}
