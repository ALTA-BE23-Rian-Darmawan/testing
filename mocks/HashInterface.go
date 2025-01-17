// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// HashInterface is an autogenerated mock type for the HashInterface type
type HashInterface struct {
	mock.Mock
}

// CheckPasswordHash provides a mock function with given fields: hashed, input
func (_m *HashInterface) CheckPasswordHash(hashed string, input string) bool {
	ret := _m.Called(hashed, input)

	if len(ret) == 0 {
		panic("no return value specified for CheckPasswordHash")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(hashed, input)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// HashPassword provides a mock function with given fields: input
func (_m *HashInterface) HashPassword(input string) (string, error) {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for HashPassword")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewHashInterface creates a new instance of HashInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHashInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *HashInterface {
	mock := &HashInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
