package repository_test

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	repo "corp/fif/inte/customers/internal/core/repository/find_customer"
	pb "corp/fif/inte/customers/internal/core/repository/find_customer/proto"
	mocks "corp/fif/inte/customers/internal/core/repository/find_customer/proto/mocks"
)

const BufSize = 1024 * 1024

var Lis *bufconn.Listener

func BufDialer(string, time.Duration) (net.Conn, error) {
	return Lis.Dial()
}

func init() {

	// Crea Servidor gRPC
	s := grpc.NewServer()
	Lis = bufconn.Listen(BufSize)
	mockClient := new(mocks.FindCustomerServiceClient)

	requestERR := &pb.FindRequest{
		Document: &pb.Document{
			Number:       "1",
			TypeDocument: "ERR",
		},
		Country: "chile",
	}
	requestUN := &pb.FindRequest{
		Document: &pb.Document{
			Number:       "1",
			TypeDocument: "DNI",
		},
		Country: "unavailable",
	}
	requestOK := &pb.FindRequest{
		Document: &pb.Document{
			Number:       "1",
			TypeDocument: "DNI",
		},
		Country: "chile",
	}

	mockClient.On("FindCustomer", mock.AnythingOfType(""), requestERR).Return(nil, status.New(codes.InvalidArgument, "Invalid Request").Err())
	mockClient.On("FindCustomer", mock.AnythingOfType(""), requestUN).Return(nil, status.New(codes.Unavailable, "unavailable").Err())
	mockClient.On("FindCustomer", mock.AnythingOfType(""), requestOK).Return(mocks.ResponsePB, nil)
	pb.RegisterFindCustomerServiceServer(s, mockClient)
	go func() {
		if err := s.Serve(Lis); err != nil {
			log.Printf("Server exited with error: %v", err)
		}
	}()
}

func TestFindRepositoryErr(t *testing.T) {

	ctx := context.TODO()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithDialer(BufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	r := repo.NewClientFindCustomer(conn)
	mockNOK := mocks.RequestType("invalid")
	resp, err := r.FindCustomer(mockNOK.Document.Number, mockNOK.Document.TypeDocument, mockNOK.Country)
	assert.Nil(t, resp)
	assert.Equal(t, err, status.New(codes.InvalidArgument, "Invalid Request").Err())

}

func TestFindRepositoryUnavailable(t *testing.T) {

	ctx := context.TODO()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithDialer(BufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	r := repo.NewClientFindCustomer(conn)
	mockNOK := mocks.RequestType("unavailable")
	resp, err := r.FindCustomer(mockNOK.Document.Number, mockNOK.Document.TypeDocument, mockNOK.Country)
	assert.Nil(t, resp)
	assert.Equal(t, err, status.New(codes.Unavailable, "gRPC Unavailable (C)").Err())

}

func TestFindRepositoryOK(t *testing.T) {

	ctx := context.TODO()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithDialer(BufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	r := repo.NewClientFindCustomer(conn)
	mockOK := mocks.RequestType("valid")
	resp, err := r.FindCustomer(mockOK.Document.Number, mockOK.Document.TypeDocument, mockOK.Country)
	assert.IsType(t, resp, (*pb.Customer)(nil))
	assert.Nil(t, err)

}
