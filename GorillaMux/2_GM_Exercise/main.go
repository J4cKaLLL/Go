package main

import (
	"fmt"
	"net/http"

	api "github.com/J4cKaLLL/Go/GorillaMux/2_GM_EXERCISE/api"
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
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Listening...")
	srv.ListenAndServe()
}
