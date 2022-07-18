package mocks

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	entity "corp/fif/inte/customers/internal/core/entity"
	pb "corp/fif/inte/customers/internal/core/repository/find_customer/proto"
)

// ErrorValidRequest ...
var ErrorValidRequest = status.New(codes.InvalidArgument, "Invalid Request").Err()

// ErrorInternal ...
var ErrorInternal = status.New(codes.Internal, "Internal Error").Err()

// PbResponseOK ...
var PbResponseOK = &pb.Customer{
	Person: &pb.Person{
		FirstName: "Daenerys",
		Document: &pb.Document{
			Number:       "1",
			TypeDocument: "DNI",
		},
	},
	IsEmployee: true,
	Address:    nil,
	Telephone:  nil,
	Email:      nil,
}

// ResponseEntity ...
var ResponseEntity = &entity.Customer{
	Person: entity.Person{
		FirstName: "Daenerys",
		Document: entity.Document{
			Number:       "1",
			TypeDocument: "DNI",
		},
	},
	IsEmployee: true,
	Address:    nil,
	Email:      nil,
	Telephone:  nil,
}
