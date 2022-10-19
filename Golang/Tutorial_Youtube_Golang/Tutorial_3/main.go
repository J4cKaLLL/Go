package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// Retorna Texto plano
/*func main() {
	e := echo.New()
	e.GET("/string", func(c echo.Context) error {
		return c.String(http.StatusOK, "<h1>Hola mundo con string</h1>")
	})
	e.Start(":8080")
}*/

// Retorna  un html con un script
/*func main() {
	e := echo.New()
	e.GET("/html", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
			<h1>Hola mundo con html</h1>
			<script>alert("HolaMUndo")</script>
		`)
	})
	e.Start(":8080")
}*/

// Retorna  un no-content no devuelve nada pero no devuelve error
func main() {
	e := echo.New()
	e.GET("/no-content", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	e.Start(":8080")
}
