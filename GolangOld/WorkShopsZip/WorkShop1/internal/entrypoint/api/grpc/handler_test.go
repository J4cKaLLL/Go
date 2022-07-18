package grpc_test

import (
	"context"
	"errors"
	"log"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	usecase "corp/fif/inte/customers/internal/core/usecase/mocks"
	entrypoint "corp/fif/inte/customers/internal/entrypoint/api/grpc"
	mocks "corp/fif/inte/customers/internal/entrypoint/api/grpc/mocks"
	pb "corp/fif/inte/customers/internal/entrypoint/api/grpc/proto"
)

const bufSize = 1024 * 1024

var (
	lis *bufconn.Listener
)

func init() {
	// New Server Grpc
	s := grpc.NewServer()
	lis = bufconn.Listen(bufSize)

	mockUseCase := new(usecase.CustomerUseCase)
	mockUseCase.On("FindCustomer", "1", "DNI", "Argentina").Return(nil, nil)
	mockUseCase.On("FindCustomer", "2", "DNI", "Argentina").Return(nil, errors.New("Internal Server Error"))
	mockUseCase.On("FindCustomer", "", "", "").Return(nil, errors.New("Internal Server Error"))

	entrypoint.NewHandlerServer(s, mockUseCase)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

}

func bufDialer(string, time.Duration) (net.Conn, error) {
	return lis.Dial()
}

func TestFindCustomer(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	assert.Nil(t, err)
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)
	resp, err := client.FindCustomer(ctx, mocks.RequestType("invalid"))
	assert.Nil(t, resp)

	client = pb.NewCustomerServiceClient(conn)
}
