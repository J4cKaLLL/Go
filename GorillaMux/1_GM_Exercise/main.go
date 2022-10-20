package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	api "github.com/J4cKaLLL/Go/GorillaMux/1_GM_EXERCISE/api"
	"github.com/gorilla/mux"
)

func HW(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{Hello World}")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HW).Methods(http.MethodGet)

	a := &api.API{}
	a.RegisterRoutes(r)

	srv := &http.Server{
		Addr:    ":8082",
		Handler: r,
	}
	fmt.Println("Listening...")
	srv.ListenAndServe()
}
