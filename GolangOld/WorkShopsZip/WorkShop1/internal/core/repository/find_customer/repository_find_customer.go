package repository

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	repository "corp/fif/inte/customers/internal/core/repository"
	pb "corp/fif/inte/customers/internal/core/repository/find_customer/proto"
)

type clientFindCustomer struct {
	cc *grpc.ClientConn
}

// NewClientFindCustomer ...
func NewClientFindCustomer(cc *grpc.ClientConn) repository.CustomerRepository {
	return &clientFindCustomer{
		cc: cc,
	}
}

func (c *clientFindCustomer) FindCustomer(documentNumber, documentType, country string) (*pb.Customer, error) {
	ctx := context.Background()
	cs := pb.NewFindCustomerServiceClient(c.cc)

	request := &pb.FindRequest{
		Document: &pb.Document{
			Number:       documentNumber,
			TypeDocument: documentType,
		},
		Country: country,
	}

	reply, err := cs.FindCustomer(ctx, request)
	if err != nil {
		e, _ := status.FromError(err)
		if e.Code() == codes.Unavailable && !strings.HasPrefix(e.Message(), "gRPC") {
			err = status.New(codes.Unavailable, "gRPC Unavailable (C)").Err()
		}

		return nil, err
	}

	return reply, nil
}
