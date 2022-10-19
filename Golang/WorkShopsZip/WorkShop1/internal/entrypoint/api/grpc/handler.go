package grpc

import (
	"context"
	"strings"

	"google.golang.org/grpc"

	transmutation "corp/fif/inte/customers/internal/core/transmutation"
	usecase "corp/fif/inte/customers/internal/core/usecase"
	pb "corp/fif/inte/customers/internal/entrypoint/api/grpc/proto"
)

// Server ...
type server struct {
	cUseCase usecase.CustomerUseCase
}

// NewHandlerServer ...
func NewHandlerServer(s *grpc.Server, customerUseCase usecase.CustomerUseCase) {
	customerServer := &server{
		cUseCase: customerUseCase,
	}

	pb.RegisterCustomerServiceServer(s, customerServer)
}

// FindCustomer ...
func (s *server) FindCustomer(ctx context.Context, in *pb.FindRequest) (*pb.Customer, error) {
	var documentNumber, documentType, country string
	if in != nil {
		documentNumber, documentType, country = in.Document.Number, in.Document.TypeDocument, strings.ToLower(in.Country)

	}

	//Implementa caso de uso (FindCustomer)
	cu, err := s.cUseCase.FindCustomer(documentNumber, documentType, country)
	if err != nil {
		return nil, err
	}

	//Transforma Entity a Objeto Grpc
	return transmutation.NewTransmutation().EncodeEntityPB(cu)
}

// UpdateCustomer ...
func (s *server) UpdateCustomer(ctx context.Context, in *pb.UpdateRequest) (*pb.Empty, error) {
	return nil, nil
}

// FindCustomerContract ...
func (s *server) FindCustomerContract(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	return nil, nil
}

// FindCustomerProducts ...
func (s *server) FindCustomerProducts(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	return nil, nil
}
