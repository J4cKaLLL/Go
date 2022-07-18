package main

import (
	"os"
	"os/signal"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	beers "example/beers/internal/beers"
	entrypoint "example/beers/internal/entrypoint/v1/http"
	validator "example/beers/internal/infrastructure/server/http/middleware"
)

func main() {

	e := echo.New()
	e.Validator = validator.NewValidator()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	usecase := beers.NewBeerUseCase()
	entrypoint.NewServerHandler(e, usecase)

	go func() {
		if err := e.Start(":8081"); err != nil {
			e.Logger.Error("Shutdown HTTP server: " + err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
