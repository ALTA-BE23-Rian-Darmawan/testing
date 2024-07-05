// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	users "BE23TODO/features/Users"

	mock "github.com/stretchr/testify/mock"
)

// DataUserInterface is an autogenerated mock type for the DataUserInterface type
type DataUserInterface struct {
	mock.Mock
}

// AccountByEmail provides a mock function with given fields: email
func (_m *DataUserInterface) AccountByEmail(email string) (*users.User, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for AccountByEmail")
	}

	var r0 *users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*users.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *users.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAccount provides a mock function with given fields: account
func (_m *DataUserInterface) CreateAccount(account users.User) error {
	ret := _m.Called(account)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(users.User) error); ok {
		r0 = rf(account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDataUserInterface creates a new instance of DataUserInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDataUserInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *DataUserInterface {
	mock := &DataUserInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
