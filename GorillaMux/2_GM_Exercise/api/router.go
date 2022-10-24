package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/cryptos", a.getCryptos).Methods(http.MethodGet)
	r.HandleFunc("/cryptos", a.postCryptoById).Methods(http.MethodPost)
	r.HandleFunc("/cryptos/{id}", a.getCryptoById).Methods(http.MethodGet)
}
