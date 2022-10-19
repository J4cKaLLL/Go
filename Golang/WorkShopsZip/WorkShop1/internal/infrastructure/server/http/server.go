package http

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_usecase "corp/fif/inte/customers/internal/core/usecase"
	entrypoint "corp/fif/inte/customers/internal/entrypoint/api/http"

	infra "bitbucket.org/falabellafif/infrastructure"
)

// RunHTTPServer ...
func RunHTTPServer(ctx context.Context, httpPort string, usecase _usecase.CustomerUseCase) error {
	// Crea instancia de servidor web Echo
	e := echo.New()

	// Deshabilita Banner de Echo
	e.HideBanner = true

	// Inicializa los Middleware y asigna a instancia de Echo
	middl := infra.InitMiddleware()
	e.Use(middl.CORS)
	e.Use(middl.ZapLogger(infra.Log))
	e.Use(middleware.Recover())
	entrypoint.NewHandlerServer(e, usecase)

	// Crea goruotine para handler de signal command + c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
		}
		e.Shutdown(ctx)
	}()

	infra.Log.Info(fmt.Sprintf("starting HTTP server in port %s...", httpPort))

	return e.Start(":" + httpPort)
}
