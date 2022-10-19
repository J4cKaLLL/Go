// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import beers "example/beers/internal/beers"
import mock "github.com/stretchr/testify/mock"

// BeerUseCase is an autogenerated mock type for the BeerUseCase type
type BeerUseCase struct {
	mock.Mock
}

// AddBeers provides a mock function with given fields: _a0
func (_m *BeerUseCase) AddBeers(_a0 beers.Beer) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(beers.Beer) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchBeers provides a mock function with given fields:
func (_m *BeerUseCase) SearchBeers() (*[]beers.Beer, error) {
	ret := _m.Called()

	var r0 *[]beers.Beer
	if rf, ok := ret.Get(0).(func() *[]beers.Beer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]beers.Beer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
