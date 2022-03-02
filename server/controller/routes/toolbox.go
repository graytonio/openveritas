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

func SendJSONData(rw http.ResponseWriter, data interface{}, code int) {
	response, _ := json.Marshal(data)
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(code)
	fmt.Fprintf(rw, "%s", string(response))
}

func IsError(err error) bool {
	if err != nil {
		log.Println(err.Error())
		return true
	}
	return false
}


func IsMongoError(err error) bool {
	if !IsError(err) {
		return false
	}

	// Do not consider none found a mongo error
	if strings.Contains(err.Error(), "no documents in result") {
		return false
	}

	return true
}

func IsNotFoundError(data interface{}) bool {
	if data == nil || reflect.ValueOf(data).IsNil() {
		return true
	}

	if reflect.ValueOf(data).Len() == 0 {
		return true
	}

	return false
}

func SendError(rw http.ResponseWriter, code int, message string) {
	err := models.NewError(code, message)
	SendJSONData(rw, err, code)
}

func SendDBError(rw http.ResponseWriter, err error) {
	SendError(rw, http.StatusInternalServerError, fmt.Sprintf("DB Error: %s", err.Error()))
}