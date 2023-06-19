// Code generated by mockery v2.30.1. DO NOT EDIT.

package repomocks

import (
	model "rent-service/pkg/model"

	mock "github.com/stretchr/testify/mock"
)

// RentRepository is an autogenerated mock type for the RentRepository type
type RentRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: rentInput
func (_m *RentRepository) Create(rentInput model.RentInput) (model.Rent, error) {
	ret := _m.Called(rentInput)

	var r0 model.Rent
	var r1 error
	if rf, ok := ret.Get(0).(func(model.RentInput) (model.Rent, error)); ok {
		return rf(rentInput)
	}
	if rf, ok := ret.Get(0).(func(model.RentInput) model.Rent); ok {
		r0 = rf(rentInput)
	} else {
		r0 = ret.Get(0).(model.Rent)
	}

	if rf, ok := ret.Get(1).(func(model.RentInput) error); ok {
		r1 = rf(rentInput)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: userId
func (_m *RentRepository) GetAll(userId string) ([]model.Rent, error) {
	ret := _m.Called(userId)

	var r0 []model.Rent
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]model.Rent, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(string) []model.Rent); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Rent)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRentRepository creates a new instance of RentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRentRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *RentRepository {
	mock := &RentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
