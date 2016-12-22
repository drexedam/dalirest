package main

import (
	"github.com/gorilla/mux"
	"github.com/drexedam/dalirest/rest"
	"net/http"
	"log"
)

const (
	staticDirectory = "/static/"
)

func main() {

	router := mux.NewRouter()

	log.Println("Registering endpoint /lp/{id}")
	router.HandleFunc("/lp/{id}", rest.GetLightPointInfo).Methods("GET")
	log.Println("Registering endpoint /lp")
	router.HandleFunc("/lp", rest.GetLightPoints).Methods("GET")
	log.Println("Registering endpoints /lz/{id} for get and post")
	router.HandleFunc("/lz/{id}", rest.ActivateScene).Methods("POST")
	log.Println("Registering endpoints /lz for get and post")
	router.HandleFunc("/lz", rest.GetScenes).Methods("GET")

	router.
	PathPrefix(staticDirectory).
		Handler(http.StripPrefix(staticDirectory, http.FileServer(http.Dir("."+staticDirectory))))

	log.Println(http.ListenAndServe(":8080", router))
}
