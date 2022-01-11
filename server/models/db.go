package models

import (
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB(mongo_db string, mongo_uri string) {
	err := mgm.SetDefaultConfig(nil, mongo_db, options.Client().ApplyURI(mongo_uri))
	if err != nil { log.Fatalln(err.Error()) }
	log.Println("Database connection established")
	
	initIndexes()
}

func initIndexes() {
	CreateIndex(*mgm.Coll(&Node{}), "name", true)
	CreateIndex(*mgm.Coll(&Property{}), "property_name", false)
	CreatePairIndex(*mgm.Coll(&Property{}), "node_name", "property_name", true)
}