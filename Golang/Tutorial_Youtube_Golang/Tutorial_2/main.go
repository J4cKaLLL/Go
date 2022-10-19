package main

import (
	"github.com/labstack/echo"
)

//SIn static
/*func main() {
	e := echo.New()
	e.File("/", "Public/index.html")
	e.File("/styles.css", "Public/styles.css")
	e.File("/script.js", "Public/script.js")
	e.Start(":1324")
}*/

//Con Static
func main() {
	e := echo.New()
	e.Static("/", "Public")
	e.Start(":1324")
}
