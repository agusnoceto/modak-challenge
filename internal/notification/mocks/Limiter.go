// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	model "github.com/agusnoceto/modak-challenge/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// Limiter is an autogenerated mock type for the Limiter type
type Limiter struct {
	mock.Mock
}

// Allow provides a mock function with given fields: key, email
func (_m *Limiter) Allow(key model.MessageKey, email string) (bool, error) {
	ret := _m.Called(key, email)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(model.MessageKey, string) (bool, error)); ok {
		return rf(key, email)
	}
	if rf, ok := ret.Get(0).(func(model.MessageKey, string) bool); ok {
		r0 = rf(key, email)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(model.MessageKey, string) error); ok {
		r1 = rf(key, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewLimiter creates a new instance of Limiter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLimiter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Limiter {
	mock := &Limiter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
