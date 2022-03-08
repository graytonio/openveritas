package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func newError(code int, message string) *Error {
	return &Error{
		Message: message,
		Code:    code,
	}
}

func sendJSONData(w http.ResponseWriter, data interface{}, code int) {
	response, _ := json.Marshal(data)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintf(w, "%s", string(response))
}

func isError(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}

	return false
}

func isDBError(err error) bool {
	if !isError(err) {
		return false
	}

	// Handle No Documents Error Differently
	if err == mongo.ErrNoDocuments {
		return false
	}

	return true
}

func isData(data interface{}) bool {
	if data == nil || reflect.ValueOf(data).IsNil() {
		return false
	}

	valueof := reflect.ValueOf(data)
	if (valueof.Kind() == reflect.Array || valueof.Kind() == reflect.Slice) && valueof.Len() == 0 {
		return false
	}

	return true
}

func sendError(rw http.ResponseWriter, code int, message string) {
	err := newError(code, message)
	sendJSONData(rw, err, code)
}

func sendDBError(rw http.ResponseWriter, err error) {
	sendError(rw, http.StatusInternalServerError, fmt.Sprintf("DB Error: %s", err.Error()))
}
