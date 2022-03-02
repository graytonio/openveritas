package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/graytonio/openveritas/server/controllers"
	"github.com/graytonio/openveritas/server/routes"
	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConfigMap struct {
	MongoDB string
	MongoURI string
	Port int
}

func main() {
	// Load ENV variables
	config := loadEnv()

	// Load DB Connection
	loadDB(config)

	// Load Routes
	r := chi.NewRouter()
	loadMiddleware(r)
	loadControllers(r)

	// Start Server
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r)
}

func loadEnv() *ConfigMap {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println(err)
	}

	return &ConfigMap{
		MongoDB: os.Getenv("MONGO_DB"),
		MongoURI: os.Getenv("MONGO_URI"),
		Port: port,
	}
}

func loadDB(config *ConfigMap) {
	err := mgm.SetDefaultConfig(nil, config.MongoDB, options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		log.Println(err)
	}

	CreateIndex(*mgm.Coll(&controllers.Node{}), "node_name", true)
	CreateIndex(*mgm.Coll(&controllers.Property{}), "prop_name", false)
	CreatePairIndex(*mgm.Coll(&controllers.Property{}), "node_name", "prop_name", true)
}

func loadMiddleware(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
}

func loadControllers(r *chi.Mux) {
	r.Mount("/node", routes.NodeRouter())
	r.Mount("/prop", routes.PropertyRouter())
}