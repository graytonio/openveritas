package models

import (
	"context"
	"log"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateIndex(collectionRef mgm.Collection, field string, unique bool) bool {
	mod := mongo.IndexModel{
		Keys: bson.M{field: 1},
		Options: options.Index().SetUnique(unique),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := collectionRef.Collection
	_, err := collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}

func CreatePairIndex(collectionRef mgm.Collection, field1 string, field2 string, unique bool) bool {
	mod := mongo.IndexModel{
		Keys: bson.D{{field1, 1}, {field2, 1}},
		Options: options.Index().SetUnique(unique),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := collectionRef.Collection
	_, err := collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}