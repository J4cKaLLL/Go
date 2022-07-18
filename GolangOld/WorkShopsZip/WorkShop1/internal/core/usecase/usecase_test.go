package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	mocks "corp/fif/inte/customers/internal/core/repository/mocks"
	usecase "corp/fif/inte/customers/internal/core/usecase"
	test_data "corp/fif/inte/customers/internal/core/usecase/mocks"
)

var uc usecase.CustomerUseCase


func init() {
	mockRepository := new(mocks.CustomerRepository)
	mockRepository.On("FindCustomer", "1", "1", "1").Return(nil, test_data.ErrorInternal)
	mockRepository.On("FindCustomer", "DNI", "1", "chile").Return(test_data.PbResponseOK, nil)

	uc = usecase.NewFindCustomerUseCase(mockRepository)
}

func TestFindCustomerUseCaseNilRequest(t *testing.T) {
	resp, err := uc.FindCustomer("", "", "")
	assert.Nil(t, resp)
	assert.EqualError(t, err, "rpc error: code = InvalidArgument desc = Invalid Request")
}

func TestFindCustomerUseCaseError(t *testing.T) {

	resp, err := uc.FindCustomer("1", "1", "1")
	assert.Nil(t, resp)
	assert.EqualError(t, err, "rpc error: code = Internal desc = Internal Error")
}

func TestFindCustomerUseCaseOK(t *testing.T) {

	resp, err := uc.FindCustomer("DNI", "1", "chile")
	assert.Nil(t, err)
	assert.Equal(t, resp, test_data.ResponseEntity)
}
