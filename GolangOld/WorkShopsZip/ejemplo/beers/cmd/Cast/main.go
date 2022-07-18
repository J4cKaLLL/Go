/*package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//main main
func main() {
	Resultado := NumToString(8)
	fmt.Println(Resultado + " Enviado")
	fmt.Println(http.StatusHTTPVersionNotSupported)
}

//NumToString Funcion para convertir entero a string
func NumToString(in int) string {
	return strconv.Itoa(in)
}
*/

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
