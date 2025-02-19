// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	event "github.com/goravel/framework/contracts/event"
	mock "github.com/stretchr/testify/mock"
)

// Listener is an autogenerated mock type for the Listener type
type Listener struct {
	mock.Mock
}

// Handle provides a mock function with given fields: args
func (_m *Listener) Handle(args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Queue provides a mock function with given fields: args
func (_m *Listener) Queue(args ...interface{}) event.Queue {
	var _ca []interface{}
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 event.Queue
	if rf, ok := ret.Get(0).(func(...interface{}) event.Queue); ok {
		r0 = rf(args...)
	} else {
		r0 = ret.Get(0).(event.Queue)
	}

	return r0
}

// Signature provides a mock function with given fields:
func (_m *Listener) Signature() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewListener creates a new instance of Listener. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewListener(t interface {
	mock.TestingT
	Cleanup(func())
}) *Listener {
	mock := &Listener{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
