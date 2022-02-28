package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/graytonio/openveritas/server/models"
)

func sendJSONData(rw http.ResponseWriter, data interface{}) {
	response, _ := json.Marshal(data)
	rw.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(rw, "%s", string(response))
}

func isError(err error) bool {
	if err != nil {
		log.Println(err.Error())
		return true
	}

	return false
}


func isMongoError(err error) bool {
	if !isError(err) {
		return false
	}

	// Do not consider none found a mongo error
	if strings.Contains(err.Error(), "no documents in result") {
		return false
	}

	return true
}

func isNotFoundError(data interface{}) bool {
	if data == nil || reflect.ValueOf(data).IsNil() {
		return true
	}

	return false
}

func sendError(rw http.ResponseWriter, code int, message string) {
	rw.WriteHeader(code)
	err := models.NewError(code, message)
	sendJSONData(rw, err)
}

func sendDBError(rw http.ResponseWriter, err error) {
	sendError(rw, http.StatusInternalServerError, fmt.Sprintf("DB Error: %s", err.Error()))
}