// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	testing "github.com/goravel/framework/contracts/testing"
	mock "github.com/stretchr/testify/mock"
)

// Testing is an autogenerated mock type for the Testing type
type Testing struct {
	mock.Mock
}

// Docker provides a mock function with given fields:
func (_m *Testing) Docker() testing.Docker {
	ret := _m.Called()

	var r0 testing.Docker
	if rf, ok := ret.Get(0).(func() testing.Docker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(testing.Docker)
		}
	}

	return r0
}

// NewTesting creates a new instance of Testing. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTesting(t interface {
	mock.TestingT
	Cleanup(func())
}) *Testing {
	mock := &Testing{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
