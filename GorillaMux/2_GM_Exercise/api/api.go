package api

import (
	"encoding/json"
	"net/http"
)

var cryptos = []string{"Bitcoin", "Ethereum", "Solana", "Polkadot", "BUSD", "USDT", "Avalanche", "Vechain", "Cardano", "Algorand"}

type API struct{}

func (a *API) GetCryptos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(cryptos)
}
