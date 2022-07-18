package api

import (
	"encoding/json"
	"net/http"
)

type API struct{}

var books = []string{"Book 1", "Book 2", "Book 3"}

func (a *API) getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func (a *API) getBook(w http.ResponseWriter, r *http.Request) {}
