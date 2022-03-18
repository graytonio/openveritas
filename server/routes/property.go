package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/graytonio/openveritas/server/controllers"
)

func PropertyRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/{query}", queryProperties)
	return r
}

func queryProperties(w http.ResponseWriter, r *http.Request) {
	query_string := chi.URLParam(r, "query")

	properties, err := controllers.QueryPropertyNames(query_string)
	if isDBError(err) {
		sendDBError(w, err)
		return
	}

	if !isData(properties) {
		sendError(w, http.StatusNotFound, fmt.Sprintf("Query %s Returned No Results", query_string))
		return
	}

	sendJSONData(w, properties, 200)
}
