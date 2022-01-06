package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/graytonio/openveritas/server/controller/routes"
)

var srv *http.Server

func StartServer() {
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatalln(err)
	}
}

func InitServer(port int) *http.Server {
	r := mux.NewRouter()

	r.HandleFunc("/node", routes.NodeHandler).Methods("GET", "POST")
	r.HandleFunc("/node/{node}", routes.NodeHandler).Methods("GET", "PUT", "DELETE")

	r.HandleFunc("/node/{node}/prop", routes.PropertyHandler).Methods("GET", "POST")
	r.HandleFunc("/node/{node}/prop/{prop}", routes.PropertyHandler).Methods("GET", "PUT", "DELETE")

	srv = &http.Server{
		Addr: "0.0.0.0:" + fmt.Sprint(port),
		WriteTimeout: time.Second * 15,
		ReadTimeout: time.Second * 15,
		IdleTimeout: time.Second * 60,
		Handler: r,
	}
	return srv
}