package routes

import (
	"log"
	"net/http"
	"strings"
)

func checkRouteError(err error, rw http.ResponseWriter) bool {
	if err != nil {
		log.Println(err.Error())

		if strings.Contains(err.Error(), "duplicate key error") {
			rw.WriteHeader(http.StatusConflict)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
		}
		return true
	}

	return false
}

func checkNotFound(data interface{}, rw http.ResponseWriter) bool {
	if data == nil {
		rw.WriteHeader(http.StatusNotFound)
		return true
	}

	return false
}