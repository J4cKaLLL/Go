package main

import (
	"encoding/json"
	"log"
	"net/http"

	"J4cKaLLL/Go-MuxRouter2.git/api"

	"github.com/gorilla/mux"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{\"message\": \"HelloWorldd\"}")
}

func main() {

	r := mux.NewRouter()
	//create the api object
	a := &api.API{}
	//register routes
	a.RegisterRoutes(r)
	//r.HandleFunc("/", handleIndex).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}
	log.Printf("Listening... port: %v", srv.Addr)
	srv.ListenAndServe()
}
