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
	node_name := vars["node_name"]
	var data interface{}
	var err error

	if node_name == "" {
		data, err = models.GetAllNodes()
	} else {
		data, err = models.GetNode(node_name)
	}

	if IsMongoError(err) {
		SendDBError(rw, err)
		return
	} else if IsNotFoundError(data) {
		SendError(rw, http.StatusNotFound, "Node Not Found")
		return
	}

	SendJSONData(rw, data, 200)
}

// Update/Create Node
func nodePutHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	node_name := vars["node_name"]

	if node_name == "" {
		SendError(rw, http.StatusBadRequest, "node_name is required parameter")
		return
	}

	var body models.NewNodeForm
	err := json.NewDecoder(r.Body).Decode(&body)
	if IsError(err) {
		SendError(rw, http.StatusBadRequest, fmt.Sprintf("Error parsing request body: %s", err.Error()))
		return
	}

	node, err := models.GetNode(node_name)
	if IsMongoError(err,) {
		SendDBError(rw, err)
		return
	}

	if IsNotFoundError(node) {
		node = models.NewNode(node_name)
	} else {
		node.Name = body.Name
	}

	_, err = models.UpdateOrCreateNode(node)
	if IsMongoError(err) {
		SendDBError(rw, err)
		return
	}
}

// Delete Node
func nodeDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	node_name := vars["node_name"]
	node, err := models.GetNode(node_name)
	if IsMongoError(err) {
		SendDBError(rw, err)
		return
	} else if IsNotFoundError(node) {
		SendError(rw, http.StatusNotFound, "Node not Found")
		return
	}

	err = models.DeleteNode(node)
	if IsMongoError(err) {
		SendDBError(rw, err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}
