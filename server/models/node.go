package models

import (
	"context"
	"errors"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewNode(name string) *Node {
	return &Node{
		Name: name,
	}
}

func GetAllNodes() (*[]Node, error) {
	log.Println("Fetching All Nodes")
	result := []Node{}
	err := mgm.Coll(&Node{}).SimpleFind(&result, bson.M{})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &result, nil
}

func GetNode(name string) (*Node, error) {
	log.Printf("Fetching Node %s", name)
	node := &Node{}
	err := mgm.Coll(node).First(bson.M{"name": name}, node)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return node, nil
}

func UpdateOrCreateNode(newNode *Node) (*Node, error) {
	log.Printf("Updating Node %s", newNode.Name)
	err := mgm.Coll(newNode).Update(newNode, options.Update().SetUpsert(true))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return newNode, nil
}

func DeleteNode(node *Node) error {
	log.Printf("Deleting Node %s", node.Name)
	err := mgm.Coll(node).Delete(node)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

//Delete Hook to remove floating properties
func (model *Node) Deleting(ctx context.Context) error {
	properties, err := GetAllPropertiesOfNode(model)
	if err != nil {
		return err
	}

	for _, p := range *properties {
		err = DeleteProperty(&p)
		if err != nil {
			return errors.New("filed to delete floating properties")
		}
	}

	return nil
}

//Update hook for updating related properties
func (model *Node) Updating(ctx context.Context) error {
	properties, err := GetAllPropertiesOfNode(model)
	if err != nil {
		return err
	}

	for _, p := range *properties {

		p.NodeID = model.ID
		p.NodeName = model.Name

		_, err := UpdateOrCreateProperty(&p)
		if err != nil {
			log.Println(err.Error())
			return errors.New("failed to update properties")
		}
	}

	return nil
}
