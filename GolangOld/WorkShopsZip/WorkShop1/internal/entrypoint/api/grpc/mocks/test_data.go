package mocks

import (
	pb "corp/fif/inte/customers/internal/entrypoint/api/grpc/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RequestType ...
func RequestType(request string) *pb.FindRequest {
	switch request {
	case "valid":
		return &validRequest
	default:
		return &invalidRequest
	}
}

var validRequest = pb.FindRequest{
	Document: &pb.Document{
		Number:       "1",
		TypeDocument: "DNI",
	},
	Country: "Argentina",
}

var invalidRequest = pb.FindRequest{
	Document: &pb.Document{
		Number:       "",
		TypeDocument: "",
	},
	Country: "",
}

// ErrorValidRequest ...
var ErrorValidRequest = status.New(codes.InvalidArgument, "Invalid Request").Err()
