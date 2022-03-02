package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

func PropertyRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/{query}", queryProperties)
	return r
}

func queryProperties(w http.ResponseWriter, r *http.Request) {

}