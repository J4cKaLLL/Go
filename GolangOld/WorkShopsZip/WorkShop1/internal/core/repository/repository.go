package repository

import (
	pbFind "corp/fif/inte/customers/internal/core/repository/find_customer/proto"
)

// CustomerRepository ...
type CustomerRepository interface {
	FindCustomer(documentNumber, documentType, country string) (*pbFind.Customer, error)
}
