package models

import (
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Property struct {
	mgm.DefaultModel `bson:",inline"`
	NodeID primitive.ObjectID `json:"node_id" bson:"node_id"`
	NodeName string `json:"node_name" bson:"node_name"`
	PropertyName string `json:"property_name" bson:"property_name"`
	PropertyValue interface{} `json:"property_value" bson:"property_value"`
}

type NewPropertyForm struct {
	PropertyName string `json:"property_name"`
	PropertyValue interface{} `json:"property_value"`
}

type UpdatePropertyForm struct {
	PropertyValue interface{} `json:"property_value"`
}

func NewProperty(node *Node, property_name string, property_value interface{}) *Property {
	return &Property{
		NodeID: node.ID,
		NodeName: node.Name,
		PropertyName: property_name,
		PropertyValue: property_value,
	}
}

func GetAllPropertiesOfNode(node *Node) *[]Property {
	properties := []Property{}
	err := mgm.Coll(&Property{}).SimpleFind(&properties, bson.M{"node_id": node.ID})
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return &properties
}

func GetAllProperties(prop_name string) (*[]Property, error) {
	properties := []Property{}
	err := mgm.Coll(&Property{}).SimpleFind(&properties, bson.M{"property_name": prop_name})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &properties, nil
}

func GetProperty(node *Node, property_name string) (*Property, error) {
	property := &Property{}
	err := mgm.Coll(property).First(bson.M{"property_name": property_name, "node_id": node.ID}, property)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return property, nil
}

func CreateProperty(node_name string, property_name string, property_value interface{}) (*Property, error) {
	node, err := GetNode(node_name)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	property := NewProperty(node, property_name, property_value)
	err = mgm.Coll(property).Create(property)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return property, nil
}

func UpdateProperty(newProperty *Property) (*Property, error) {
	err := mgm.Coll(newProperty).Update(newProperty)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}	
	return newProperty, nil
}

func DeleteProperty(property *Property) error {
	err := mgm.Coll(property).Delete(property)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}