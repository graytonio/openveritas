package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/graytonio/openveritas/server/controllers"
)

func NodeRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", getAllNodes)
	r.Get("/{node_name}", getNodeByName)
	r.Put("/{node_name}", putNode)
	r.Delete("/{node_name}", deleteNode)
	r.Mount("/{node_name}/prop", NodePropertyRouter())
	return r
}

func getAllNodes(w http.ResponseWriter, r *http.Request) {
	nodes, err := controllers.GetAllNodes()
	if isDBError(err) {
		sendDBError(w, err)
		return
	}

	if !isData(nodes) {
		sendError(w, http.StatusNotFound, "No Nodes Found")
		return
	}

	sendJSONData(w, nodes, http.StatusOK)
}

func getNodeByName(w http.ResponseWriter, r *http.Request) {
	node_name := chi.URLParam(r, "node_name")
	node, err := controllers.GetNodeByName(node_name)
	if isDBError(err) {
		sendDBError(w, err)
		return
	}

	if !isData(node) {
		sendError(w, http.StatusNotFound, fmt.Sprintf("Node %s Not Found", node_name))
		return
	}

	sendJSONData(w, node, http.StatusOK)
}

func putNode(w http.ResponseWriter, r *http.Request) {
	node_name := chi.URLParam(r, "node_name")

	var body controllers.Node
	err := json.NewDecoder(r.Body).Decode(&body)
	if isError(err) {
		sendError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request body: %s", err.Error()))
		return
	}

	node, err := controllers.GetNodeByName(node_name)
	if isDBError(err) {
		sendDBError(w, err)
		return
	}

	if isData(node) {
		if isDBError(err) {
			sendDBError(w, err)
			return
		}

		node.NodeName = body.NodeName

		updateNode(node, w, r)
	} else {
		if node_name != body.NodeName {
			sendError(w, http.StatusBadRequest, "Body Parameter node_name and Url Parameter node_name Must Match")
		}

		createNode(body.NodeName, w, r)
	}
}

func createNode(node_name string, w http.ResponseWriter, r *http.Request) {
	node := &controllers.Node{
		NodeName: node_name,
	}
	err := controllers.CreateNode(node)

	if isDBError(err) {
		sendDBError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func updateNode(node *controllers.Node, w http.ResponseWriter, r *http.Request) {
	err := controllers.UpdateNode(node)
	if isDBError(err) {
		sendDBError(w, err)
		return
	}

}

func deleteNode(w http.ResponseWriter, r *http.Request) {
	node_name := chi.URLParam(r, "node_name")
	node, err := controllers.GetNodeByName(node_name)
	if isDBError(err) {
		sendDBError(w, err)
		return
	}

	if !isData(node) {
		sendError(w, 404, fmt.Sprintf("Node %s Not Found", node_name))
		return
	}

	err = controllers.DeleteNode(node)
	if isDBError(err) {
		sendDBError(w, err)
		return
	}
}
