package controllers

import (
	"context"
	"fmt"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Node struct {
	mgm.DefaultModel `bson:",inline"`
	NodeName         string `json:"node_name" bson:"node_name"`
}

func NewNode(node_name string) *Node {
	return &Node{
		NodeName: node_name,
	}
}

func GetAllNodes() ([]Node, error) {
	nodes := []Node{}
	NodeCollection := mgm.Coll(&Node{})

	err := NodeCollection.SimpleFind(&nodes, bson.M{})
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

func GetNodeByName(node_name string) (*Node, error) {
	node := &Node{}
	NodeCollection := mgm.Coll(node)

	err := NodeCollection.First(bson.D{{Key: "node_name", Value: node_name}}, node)
	if err != nil {
		return nil, err
	}

	return node, nil
}

func CreateNode(node *Node) error {
	NodeCollection := mgm.Coll(node)
	err := NodeCollection.Create(node)
	return err
}

func UpdateNode(node *Node) error {
	NodeCollection := mgm.Coll(node)
	err := NodeCollection.Update(node)
	return err
}

func DeleteNode(node *Node) error {
	NodeCollection := mgm.Coll(node)
	err := NodeCollection.Delete(node)
	return err
}

// Updating Hook
func (node *Node) Updated(ctx context.Context, result *mongo.UpdateResult) error {
	fmt.Println("UPDATING")
	props, err := GetAllNodeProperties(node)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, p := range props {
		p.NodeName = node.NodeName
		err = UpdateProperty(&p)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

// Deleting Hook
func (node *Node) Deleting(ctx context.Context) error {
	props, err := GetAllNodeProperties(node)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, p := range props {
		err = DeleteProperty(&p)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
