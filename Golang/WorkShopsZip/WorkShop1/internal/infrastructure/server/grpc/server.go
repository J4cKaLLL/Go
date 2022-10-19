package grpc

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	_usecase "corp/fif/inte/customers/internal/core/usecase"
	entrypoint "corp/fif/inte/customers/internal/entrypoint/api/grpc"

	infra "bitbucket.org/falabellafif/infrastructure"
)

// RunGrpcServer ...
func RunGrpcServer(ctx context.Context, grpcPort string, usecase _usecase.CustomerUseCase) error {
	listen, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}
	opts = infra.AddMiddleware(infra.Log, opts)

	// Crea Grpc Server
	server := grpc.NewServer(opts...)
	entrypoint.NewHandlerServer(server, usecase)

	// Crea goruotine para handler de signal command + c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			infra.Log.Warn("shutting down gRPC server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	infra.Log.Info(fmt.Sprintf("starting gRPC server in port %s...", grpcPort))
	return server.Serve(listen)
}
