package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/random_name_svc/config"
)

// Init is used to initialize the routes
func Init() {
	r := mux.NewRouter()
	r.HandleFunc("/", getRandomJokeByName).Methods(http.MethodGet)

	log.Printf("webserver serving on :%v", config.RunOnPort)
	http.ListenAndServe(fmt.Sprintf(":%v", config.RunOnPort), r)
}
