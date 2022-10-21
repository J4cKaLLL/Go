package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/getcryptos", a.GetCryptos).Methods(http.MethodGet)
}
