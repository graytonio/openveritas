package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/graytonio/openveritas/server/controllers"
)

func NodePropertyRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(fetchNodeMiddleware)
	r.Get("/", getAllNodeProperties)
	r.Get("/{prop_name}", getNodePropertyByName)
	r.Put("/{prop_name}", putNodeProperty)
	r.Delete("/{prop_name}", deleteNodeProperty)
	return r
}

type propertyContext string

func fetchNodeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		
		ctx := context.WithValue(r.Context(), propertyContext("node"), node)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllNodeProperties(w http.ResponseWriter, r *http.Request) {
	node := r.Context().Value(propertyContext("node")).(*controllers.Node)

	props, err := controllers.GetAllNodeProperties(node)
	if isDBError(err) {
		sendDBError(w, err)
		return
	}

	if !isData(props) {
		sendError(w, http.StatusNotFound, "No Properties Found")
		return
	}

	sendJSONData(w, props, 200)
}

func getNodePropertyByName(w http.ResponseWriter, r *http.Request) {
	node := r.Context().Value(propertyContext("node")).(*controllers.Node)
	prop_name := chi.URLParam(r, "prop_name")

	prop, err := controllers.GetNodePropertyByName(node, prop_name)
	if isDBError(err) {
		sendDBError(w, err)
		return
	}

	if !isData(prop) {
		sendError(w, http.StatusNotFound, fmt.Sprintf("Property %s of Node %s Not Found", prop_name, node.NodeName))
		return
	}

	sendJSONData(w, prop, 200)
}

func putNodeProperty(w http.ResponseWriter, r *http.Request) {
	node := r.Context().Value(propertyContext("node")).(*controllers.Node)
	prop_name := chi.URLParam(r, "prop_name")

	var body controllers.Property
	err := json.NewDecoder(r.Body).Decode(&body)
	if isError(err) {
		sendError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request body: %s", err.Error()))
		return
	}

	prop, err := controllers.GetNodePropertyByName(node, prop_name)
	if isDBError(err) {
		sendDBError(w, err)
		return
	}

	if isData(prop) {
		prop.PropertyValue = body.PropertyValue
		updateProperty(prop, w, r)
	} else {
		createProperty(node, body.PropertyName, body.PropertyValue, w, r)
	}
}

func createProperty(node *controllers.Node, prop_name string, prop_value interface{}, w http.ResponseWriter, r *http.Request) {
	prop := &controllers.Property{
		PropertyName: prop_name,
		PropertyValue: prop_value,
		NodeName: node.NodeName,
		NodeId: node.ID,
	}
	err := controllers.CreateProperty(prop)

	if isDBError(err) {
		sendDBError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func updateProperty(prop *controllers.Property, w http.ResponseWriter, r *http.Request) {
	err := controllers.UpdateProperty(prop)
	if isDBError(err) {
		sendDBError(w, err)
		return
	}
}

func deleteNodeProperty(w http.ResponseWriter, r *http.Request) {
	node := r.Context().Value(propertyContext("node")).(*controllers.Node)
	prop_name := chi.URLParam(r, "prop_name")
	prop, err := controllers.GetNodePropertyByName(node, prop_name)
	if isDBError(err) {
		sendDBError(w, err)
		return
	}

	if !isData(prop) {
		sendError(w, 404, fmt.Sprintf("Property %s of Node %s Not Found", prop_name, node.NodeName))
		return
	}

	err = controllers.DeleteProperty(prop)
	if isDBError(err) {
		sendDBError(w, err)
		return
	}
}