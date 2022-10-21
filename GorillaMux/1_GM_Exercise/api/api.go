package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
)

type API struct{}

type CryptosParams struct {
	LimitMin int `schema: "limitMin"`
	LimitMax int `schema: "limitMax"`
}

var (
	cryptos = []string{"Bitcoin", "Ethereum", "Solana", "Polkadot", "USDT", "USDC", "DAI", "Avalanche", "Cardano"}
	decoder = schema.NewDecoder()
)

func (a *API) GetCryptos(w http.ResponseWriter, r *http.Request) {

	params := &CryptosParams{}
	//var params CryptosParams

	err := decoder.Decode(params, r.URL.Query())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if params.LimitMin < 0 || params.LimitMin > len(cryptos) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if params.LimitMax < 0 || params.LimitMax > len(cryptos) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(cryptos[params.LimitMin:params.LimitMax])
}
