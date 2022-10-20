package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type API struct{}

type CryptosParams struct {
	Limit int `schema: "limit"`
}

var cryptos = []string{"Bitcoin", "Ethereum", "Solana", "Polkadot", "USDT", "USDC", "DAI", "Avalanche", "Cardano"}

func (a *API) GetCryptos(w http.ResponseWriter, r *http.Request) {
	limitMin := r.URL.Query().Get("limitMin")
	limitMax := r.URL.Query().Get("limitMax")
	lMin, err := strconv.Atoi(limitMin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lMax, err := strconv.Atoi(limitMax)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if lMin < 0 || lMax > len(cryptos) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(cryptos[lMin:lMax])

}
