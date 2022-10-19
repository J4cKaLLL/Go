package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	// Se crea una instancia de echo
	e := echo.New()
	// Se crea el handler donde "/" es la raiz de la pagina web
	//crear una funcion que es el handler encargado de responder cuando se haga la peticion
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hola mundo")
	})
	//Escuche en puerto 1323
	e.Logger.Fatal(e.Start(":1323"))
}
