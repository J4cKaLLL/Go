package main

import (
	"fmt"
	"net/http"

	api "exercise/api"

	"github.com/gorilla/mux"
)

func GetGreeting(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", GetGreeting)

	a := &api.API{}

	a.RegisterRoutes(router)

	srv := &http.Server{
		Addr:    ":8082",
		Handler: router,
	}

	fmt.Println("Listening...")
	srv.ListenAndServe()
}
