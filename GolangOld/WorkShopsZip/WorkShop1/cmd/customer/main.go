package main

import (
	"context"
	"fmt"
	"strings"

	file "corp/fif/inte/customers/internal/infrastructure/config/file"
	flag "corp/fif/inte/customers/internal/infrastructure/config/flag"
	grpc "corp/fif/inte/customers/internal/infrastructure/server/grpc"
	http "corp/fif/inte/customers/internal/infrastructure/server/http"

	infra "bitbucket.org/falabellafif/infrastructure"

	repository "corp/fif/inte/customers/internal/core/repository/find_customer"
	use_case "corp/fif/inte/customers/internal/core/usecase"
)

func main() {
	ctx := context.Background()

	// Inicializa los flag del binario
	configFlag := flag.NewFlagConfig()
	if configFlag.Help {
		flag.ShowHelp()
		return
	}

	// Carga configuraciones del archivo config.yaml
	configFile, err := file.NewFileConfig()
	if err != nil {
		fmt.Printf("failed to initialize config: %v", err)
		panic(err)
	}

	// Inicializa los logs
	if err := infra.Init(configFlag.LogLevel, configFlag.LogTimeFormat); err != nil {
		fmt.Printf("failed to initialize logger: %v", err)
		panic(err)
	}

	// Crea un mapa que trae la lista de conexiones de clientes grpc
	client := make(map[string]string)
	for _, element := range configFile.Repository.Client {
		client[strings.ToLower(element.Service)] = element.IP + ":" + element.Port
	}

	// Carga Dial Grpc
	if err := infra.RunDialGrpc(client); err != nil {
		infra.Log.Error("Error: " + err.Error())
	}
	defer infra.CloseDialGrpc()

	//Inicializa los repositorios que implementan las conexiones creadas anteriormente

	find := repository.NewClientFindCustomer(infra.ClientConn["find"])
	uc := use_case.NewFindCustomerUseCase(find)

	// Inicia el servicios HTTP y gRPC
	go func() {
		http.RunHTTPServer(ctx, configFlag.HTTPPort, uc)
	}()

	grpc.RunGrpcServer(ctx, configFlag.GRPCPort, uc)

}
