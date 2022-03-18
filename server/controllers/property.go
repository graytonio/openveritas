package controllers

import (
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Property struct {
	mgm.DefaultModel `bson:",inline"`
	PropertyName     string             `json:"prop_name" bson:"prop_name"`
	PropertyValue    interface{}        `json:"prop_value" bson:"prop_value"`
	NodeName         string             `json:"node_name" bson:"node_name"`
	NodeId           primitive.ObjectID `json:"node_id" bson:"node_id"`
}

func NewProperty(node *Node, prop_name string, prop_value interface{}) *Property {
	return &Property{
		PropertyName:  prop_name,
		PropertyValue: prop_value,
		NodeName:      node.NodeName,
		NodeId:        node.ID,
	}
}

func GetAllNodeProperties(node *Node) ([]Property, error) {
	properties := []Property{}
	PropertyCollection := mgm.Coll(&Property{})

	err := PropertyCollection.SimpleFind(&properties, bson.D{{Key: "node_id", Value: node.ID}})
	if err != nil {
		return nil, err
	}

	return properties, nil
}

func GetNodePropertyByName(node *Node, prop_name string) (*Property, error) {
	property := &Property{}
	PropertyCollection := mgm.Coll(property)

	err := PropertyCollection.First(bson.D{{Key: "node_name", Value: node.NodeName}, {Key: "prop_name", Value: prop_name}}, property)
	if err != nil {
		return nil, err
	}

	return property, nil
}

func QueryPropertyNames(query_string string) ([]Property, error) {
	properties := []Property{}
	PropertyCollection := mgm.Coll(&Property{})

	// searchStage := bson.D{{"$search", bson.D{{"wildcard", bson.D{{"path", "title"}, {"query", query_string}}}}}}
	// searchStage := bson.D{{"prop_name", primitive.Regex{Pattern: query_string, Options: "i"}}}

	searchStage := bson.M{"prop_name": bson.M{"$regex": primitive.Regex{Pattern: query_string, Options: "i"}}}

	log.Println(searchStage)

	err := PropertyCollection.SimpleFind(&properties, searchStage)

	if err != nil {
		return nil, err
	}

	return properties, nil
}

func CreateProperty(prop *Property) error {
	PropertyCollection := mgm.Coll(prop)
	err := PropertyCollection.Create(prop)
	return err
}

func UpdateProperty(prop *Property) error {
	PropertyCollection := mgm.Coll(prop)
	err := PropertyCollection.Update(prop)
	return err
}

func DeleteProperty(prop *Property) error {
	PropertyCollection := mgm.Coll(prop)
	err := PropertyCollection.Delete(prop)
	return err
}
