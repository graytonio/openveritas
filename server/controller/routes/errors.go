package routes

import (
	"log"
	"net/http"
	"reflect"
	"strings"
)

func checkRouteError(err error, rw http.ResponseWriter) bool {
	if err != nil {
		log.Println(err.Error())
		if strings.Contains(err.Error(), "no documents in result") {
			return false
		} else if strings.Contains(err.Error(), "duplicate key") {
			rw.WriteHeader(http.StatusConflict)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
		}
		return true
	}

	return false
}

func checkNotFound(data interface{}, rw http.ResponseWriter) bool {
	log.Printf("Checking if %v exists", data)
	if data == nil || reflect.ValueOf(data).IsNil() {
		log.Println("NOT FOUND")
		rw.WriteHeader(http.StatusNotFound)
		return true
	}

	return false
}