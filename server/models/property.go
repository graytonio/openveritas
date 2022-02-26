package models

import (
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func NewProperty(node *Node, property_name string, property_value interface{}) *Property {
	return &Property{
		NodeID:        node.ID,
		NodeName:      node.Name,
		PropertyName:  property_name,
		PropertyValue: property_value,
	}
}

func GetAllPropertiesOfNode(node *Node) (*[]Property, error) {
	log.Printf("Fetching all properties of node %s", node)
	properties := []Property{}
	err := mgm.Coll(&Property{}).SimpleFind(&properties, bson.M{"node_id": node.ID})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &properties, nil
}

func GetAllProperties(prop_name string) (*[]Property, error) {
	log.Printf("Fetching all nodes with property %s", prop_name)
	properties := []Property{}
	err := mgm.Coll(&Property{}).SimpleFind(&properties, bson.M{"property_name": prop_name})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &properties, nil
}

func GetProperty(node *Node, property_name string) (*Property, error) {
	log.Printf("Fetching property %s of node %s", property_name, node.Name)
	property := &Property{}
	err := mgm.Coll(property).First(bson.M{"property_name": property_name, "node_id": node.ID}, property)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return property, nil
}

func UpdateOrCreateProperty(newProperty *Property) (*Property, error) {
	log.Printf("Updating property %s to %v", newProperty.PropertyName, newProperty.PropertyValue)
	err := mgm.Coll(newProperty).Update(newProperty, options.Update().SetUpsert(true))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return newProperty, nil
}

func DeleteProperty(property *Property) error {
	log.Printf("Deleting property %s", property.PropertyName)
	err := mgm.Coll(property).Delete(property)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
