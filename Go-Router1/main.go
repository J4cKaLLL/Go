package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/J4cKaLLL/Go-MuxRouter/api"
	"github.com/gorilla/mux"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{\"message\": \"HelloWorld\"}")
}

func main() {

	r := mux.NewRouter()
	a := &api.API{}
	a.RegisterRoutes(r)
	r.HandleFunc("/", handleIndex).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}
	log.Println("Listening...")
	srv.ListenAndServe()
}
