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
	node_name := mux.Vars(r)["node"]
	if node_name == "" {
		nodes, err := models.GetAllNodes()
		if handleMongoError(err, rw) || handleNotFoundError(nodes, rw) {
			return
		}
		response, _ := json.Marshal(nodes)
		rw.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(rw, "%s", string(response))
	} else {
		node, err := models.GetNode(node_name)
		if handleMongoError(err, rw) || handleNotFoundError(node, rw) {
			return
		}

		respoonse, _ := json.Marshal(node)
		rw.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(rw, "%s", string(respoonse))
	}
}

// Update/Create Node
func nodePutHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	node_name := vars["node"]

	var body models.NewNodeForm
	err := json.NewDecoder(r.Body).Decode(&body)
	if handleBodyParseError(err, rw) {
		return
	}

	node, err := models.GetNode(node_name)
	if handleMongoError(err, rw) {
		return
	}

	if node == nil {
		node = models.NewNode(body.Name)
	} else {
		node.Name = body.Name
	}

	_, err = models.UpdateOrCreateNode(node)
	if handleMongoError(err, rw) {
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func nodeDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	node_name := mux.Vars(r)["node"]
	node, err := models.GetNode(node_name)
	if handleMongoError(err, rw) || handleNotFoundError(node, rw) {
		return
	}

	err = models.DeleteNode(node)
	if handleMongoError(err, rw) {
		return
	}

	rw.WriteHeader(http.StatusOK)
}
