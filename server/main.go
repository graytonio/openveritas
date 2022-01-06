package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/graytonio/openveritas/server/controller"
	"github.com/graytonio/openveritas/server/models"
	"github.com/joho/godotenv"
)

var srv *http.Server
var wait time.Duration

var mongo_db, mongo_uri string
var port int

func init() {
	var wait time.Duration
    flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
    flag.Parse()

	initEnv()
	models.InitDB(mongo_db, mongo_uri)
	srv = controller.InitServer(port)
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		if !strings.Contains(err.Error(), "no such file or directory") {
			log.Fatalln("error loading .env file")
		}
	}

	mongo_db = os.Getenv("MONGO_DB")
	mongo_uri = os.Getenv("MONGO_URI")
	port, _ = strconv.Atoi(os.Getenv("PORT"))
}

func waitForExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
}

func main(){
	go controller.StartServer()
	log.Printf("Web interface listening on port %d", port)

	waitForExit()
	log.Println("Shutting down")
	os.Exit(0)
}