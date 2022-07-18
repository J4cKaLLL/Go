package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type API struct {
}

type BooksParams struct {
	Limit  int `schema:"limit"`
	Offset int `schema:"offset"`
}

var (
	books   = []string{"Book 1", "Book 2", "Book 3"}
	decoder = schema.NewDecoder()
)

//retorna la lista de libros
func (a *API) getBooks(w http.ResponseWriter, r *http.Request) {
	params := &BooksParams{}

	err := decoder.Decode(params, r.URL.Query())

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if params.Offset < 0 || params.Offset > len(books) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if params.Limit < 0 || params.Limit > len(books) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(books[params.Offset:params.Limit])
}

//retorna un libro con un ID
func (a *API) getBook(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	idParam := pathParams["id"]

	id, err := strconv.Atoi(idParam)
	fmt.Println("len(books)", len(books))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if id-1 < 0 || id-1 >= len(books) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(books[id-1])
}
