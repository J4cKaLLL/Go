package main

import (
	"log"
	"net/http"

	"github.com/J4cKaLLL/Go/GoMySQL/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Obtiene un puntero de mux.Router
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
