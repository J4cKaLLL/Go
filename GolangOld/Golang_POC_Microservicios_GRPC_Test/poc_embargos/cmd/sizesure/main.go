package main

import (
    "log"
    "os"
    "os/signal"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
    "google.golang.org/grpc"

    
    reposito "sizesure/internal/impl/impl_1"
    usec "sizesure/internal/usecase"
    entrypoint "sizesure/cmd/sizesure/internal/http_handler"
    validator "sizesure/cmd/sizesure/internal/http_handler"
    
)

func main() {
    e := echo.New()
    e.Validator = validator.NewValidator()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    // se realiza el mapeo de conexión al server grpc
    conn, err := grpc.Dial("10.128.2.166:9090", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		panic(err)
    }
    //inicializa el repository que requiere la conexión grpc anterior
    repo := reposito.NewSeizureRepository(conn)
    //inicializa el usecase que requiere del repository anterior
    usecase := usec.NewSizesuresUseCase(repo)
    //inicializa el handler que requiere del usecase anterior
    entrypoint.NewServerHandler(e, usecase)

    go func(){
        if err := e.Start(":8080"); err != nil {
            e.Logger.Error("Shutdown HTTP Server:" + err.Error())
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt)
    <-quit
}
