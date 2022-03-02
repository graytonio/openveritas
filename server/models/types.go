package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Node struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
}

type NewNodeForm struct {
	Name string `json:"name"`
}

type Property struct {
	mgm.DefaultModel `bson:",inline"`
	NodeID           primitive.ObjectID `json:"node_id" bson:"node_id"`
	NodeName         string             `json:"node_name" bson:"node_name"`
	PropertyName     string             `json:"property_name" bson:"property_name"`
	PropertyValue    interface{}        `json:"property_value" bson:"property_value"`
}

type PropertyForm struct {
	PropertyName  string      `json:"property_name"`
	PropertyValue interface{} `json:"property_value"`
}

type Error struct {
	Message string `json:"message"`
	Code int `json:"code"`
}

func NewError(code int, message string) *Error {
	return &Error {
		Message: message,
		Code: code,
	}
}