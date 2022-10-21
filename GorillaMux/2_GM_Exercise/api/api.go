package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
)

var cryptos = []string{"Bitcoin", "Ethereum", "Solana", "Polkadot", "BUSD", "USDT", "Avalanche", "Vechain", "Cardano", "Algorand"}

type API struct{}

type limites struct {
	LimitMin int `schema: "limitMin"`
	LimitMax int `schema: "limitMax"`
}

func (a *API) GetCryptos(w http.ResponseWriter, r *http.Request) {

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
	json.NewEncoder(w).Encode(cryptos[limits.LimitMin:limits.LimitMax])
}
