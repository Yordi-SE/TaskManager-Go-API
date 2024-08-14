// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// InfrastructureInterface is an autogenerated mock type for the InfrastructureInterface type
type InfrastructureInterface struct {
	mock.Mock
}

// ComparePasswords provides a mock function with given fields: hashedPassword, password
func (_m *InfrastructureInterface) ComparePasswords(hashedPassword string, password string) error {
	ret := _m.Called(hashedPassword, password)

	if len(ret) == 0 {
		panic("no return value specified for ComparePasswords")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(hashedPassword, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateToken provides a mock function with given fields: user_name, id, role
func (_m *InfrastructureInterface) GenerateToken(user_name string, id primitive.ObjectID, role string) (string, error) {
	ret := _m.Called(user_name, id, role)

	if len(ret) == 0 {
		panic("no return value specified for GenerateToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, primitive.ObjectID, string) (string, error)); ok {
		return rf(user_name, id, role)
	}
	if rf, ok := ret.Get(0).(func(string, primitive.ObjectID, string) string); ok {
		r0 = rf(user_name, id, role)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, primitive.ObjectID, string) error); ok {
		r1 = rf(user_name, id, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HashPassword provides a mock function with given fields: password
func (_m *InfrastructureInterface) HashPassword(password string) (string, error) {
	ret := _m.Called(password)

	if len(ret) == 0 {
		panic("no return value specified for HashPassword")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(password)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateToken provides a mock function with given fields: tokenString
func (_m *InfrastructureInterface) ValidateToken(tokenString string) (bool, error) {
	ret := _m.Called(tokenString)

	if len(ret) == 0 {
		panic("no return value specified for ValidateToken")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(tokenString)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(tokenString)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewInfrastructureInterface creates a new instance of InfrastructureInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInfrastructureInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *InfrastructureInterface {
	mock := &InfrastructureInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
