package usecase

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	entity "corp/fif/inte/customers/internal/core/entity"
	repository "corp/fif/inte/customers/internal/core/repository"
	transmutation "corp/fif/inte/customers/internal/core/transmutation"
)

// CustomerUseCase ...
type CustomerUseCase interface {
	FindCustomer(documentNumber, documentType, country string) (*entity.Customer, error)
}

type customerUseCase struct {
	CustomerRepository repository.CustomerRepository
}

// NewFindCustomerUseCase ...
func NewFindCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		CustomerRepository: repo,
	}
}

func (c *customerUseCase) FindCustomer(documentNumber, documentType, country string) (*entity.Customer, error) {

	//Valida request
	if err := validateRequest(documentNumber, documentType, country); err != nil {
		return nil, err
	}

	resp, err := c.CustomerRepository.FindCustomer(documentNumber, documentType, country)
	if err != nil {
		return nil, err
	}

	return transmutation.NewTransmutation().EncodePBEntityFind(resp)

}

//funcion validación datos request
func validateRequest(documentNumber, documentType, country string) error {
	if documentNumber == "" || documentType == "" || country == "" {
		return status.New(codes.InvalidArgument, "Invalid Request").Err()
	}
	return nil
}
