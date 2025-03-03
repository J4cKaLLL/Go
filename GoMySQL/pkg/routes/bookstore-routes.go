package routes

import (
	"github.com/J4cKaLLL/Go/GoMySQL/pkg/controllers"
	"github.com/gorilla/mux"
)

// Selecciona una ruta con el emparejador

var RegisterBookStoreRoutes = func(router *mux.Router) {
	//router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")

	// router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	// router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	// router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
