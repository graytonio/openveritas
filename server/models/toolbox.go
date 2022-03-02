package models

import (
	"log"
	"strings"
)

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