package mocks

import (
	pb "corp/fif/inte/customers/internal/core/repository/find_customer/proto"
)

// RequestType ...
func RequestType(request string) *pb.FindRequest {
	switch request {
	case "valid":
		return &validRequest
	case "unavailable":
		return &unavailableRequest
	default:
		return &invalidRequest
	}
}

var invalidRequest = pb.FindRequest{
	Document: &pb.Document{
		Number:       "1",
		TypeDocument: "ERR",
	},
	Country: "chile",
}

var validRequest = pb.FindRequest{
	Document: &pb.Document{
		Number:       "1",
		TypeDocument: "DNI",
	},
	Country: "chile",
}

var unavailableRequest = pb.FindRequest{
	Document: &pb.Document{
		Number:       "1",
		TypeDocument: "DNI",
	},
	Country: "unavailable",
}

// ResponsePB ...
var ResponsePB = &pb.Customer{
	Person: &pb.Person{
		Document: &pb.Document{
			Number:       "1",
			TypeDocument: "DNI",
		},
	},
}
