package routes

import (
	"log"
	"net/http"
	"reflect"
	"strings"
)

func isError(err error) bool {
	if err != nil {
		log.Println(err.Error())
		return true
	}

	return false
}

func handleBodyParseError(err error, rw http.ResponseWriter) bool {
	if isError(err) {
		rw.WriteHeader(http.StatusBadRequest)
		return true
	}

	return false
}

func handleMongoError(err error, rw http.ResponseWriter) bool {
	if !isError(err) {
		return false
	}

	// Do not consider none found a mongo error
	if strings.Contains(err.Error(), "no documents in result") {
		return false
	}

	rw.WriteHeader(http.StatusInternalServerError)
	return true
}

func handleNotFoundError(data interface{}, rw http.ResponseWriter) bool {
	if data == nil || reflect.ValueOf(data).IsNil() {
		rw.WriteHeader(http.StatusNotFound)
		return true
	}

	return false
}
