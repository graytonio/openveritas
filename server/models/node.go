package models

import (
	"context"
	"errors"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type Node struct {
	mgm.DefaultModel `bson:",inline"`
	Name string `json:"name" bson:"name"`
}

type NewNodeForm struct {
	Name string `json:"name"`
}

func NewNode(name string) *Node {
	return &Node{
		Name: name,
	}
}

func GetAllNodes() *[]Node {
	result := []Node{}
	mgm.Coll(&Node{}).SimpleFind(&result, bson.M{})
	return &result
}

func GetNode(name string) *Node {
	node := &Node{}
	err := mgm.Coll(node).First(bson.M{"name": name}, node)
	if err != nil {	
		log.Println(err.Error())
		return nil
	}
	return node
}

func CreateNode(name string) error {
	node := NewNode(name)
	err := mgm.Coll(node).Create(node)
	if err != nil {
		log.Println(err.Error()) 
		return err
	}
	return nil
}

func UpdateNode(newNode *Node) error {
	err := mgm.Coll(newNode).Update(newNode)
	if err != nil { 
		log.Println(err.Error())
		return err
	}
	return nil
}

func DeleteNode(node *Node) error {	
	err := mgm.Coll(node).Delete(node)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

//Update hook for updating related properties
func (model *Node) Updating(ctx context.Context) error {
	properties := GetAllPropertiesOfNode(model)

	success := true

	for _, p := range *properties {
		p.NodeID = model.ID
		p.NodeName = model.Name
		err := UpdateProperty(&p)
		if err != nil {
			log.Println(err.Error())
			success = true
		}
	}

	if !success {
		return errors.New("failed to update properties")
	}

	return nil
}