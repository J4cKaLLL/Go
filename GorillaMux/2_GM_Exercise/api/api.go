package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var cryptos = []string{"Bitcoin", "Ethereum", "Solana", "Polkadot", "BUSD", "USDT", "Avalanche", "Vechain", "Cardano", "Algorand"}

type API struct{}

type limites struct {
	LimitMin int `schema: "limitMin"`
	LimitMax int `schema: "limitMax"`
}

type PostCrypto struct {
	Name string `json: "name"`
}

// Retorna todos las cryptos
func (a *API) getCryptos(w http.ResponseWriter, r *http.Request) {
	limits := &limites{}

	decoder := schema.NewDecoder()

	err := decoder.Decode(limits, r.URL.Query())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if limits.LimitMin < 0 || limits.LimitMin > len(cryptos) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if limits.LimitMax < 0 || limits.LimitMax > len(cryptos) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if limits.LimitMax == 0 {
		json.NewEncoder(w).Encode(cryptos[limits.LimitMin : len(cryptos)-1])
	} else {
		json.NewEncoder(w).Encode(cryptos[limits.LimitMin:limits.LimitMax])
	}
}

// Retorna una crypto con un Id En especifico
func (a *API) getCryptoById(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	idParam := pathParams["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if (id-1) < 0 || (id-1) > len(cryptos)-1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(cryptos[id-1])
}

func (a *API) postCryptoById(w http.ResponseWriter, r *http.Request) {
	crypt := &PostCrypto{}
	err := json.NewDecoder(r.Body).Decode(crypt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cryptos = append(cryptos, crypt.Name)
	w.WriteHeader(http.StatusCreated)
}
