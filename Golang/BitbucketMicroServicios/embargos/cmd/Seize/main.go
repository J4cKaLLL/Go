package main

import (
	"os"
	"os/signal"

	entrypoint "seizures/internal/entrypoint/v1/http"
	validator "seizures/internal/infrastructure/server/http/middleware"
	seizures "seizures/internal/seizures"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()
	e.Validator = validator.NewValidator()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	usecase := seizures.NewSeizureUseCase()
	entrypoint.NewServerHandler(e, usecase)

	if err := e.Start(":8082"); err != nil {
		e.Logger.Error("Shutdown HTTP server: " + err.Error())
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

}
