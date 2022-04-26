// Code generated by mockery 2.12.1. DO NOT EDIT.

package mocks

import (
	context "context"

	manifest "github.com/ovrclk/akash/provider/manifest"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// StatusClient is an autogenerated mock type for the StatusClient type
type StatusClient struct {
	mock.Mock
}

// Status provides a mock function with given fields: _a0
func (_m *StatusClient) Status(_a0 context.Context) (*manifest.Status, error) {
	ret := _m.Called(_a0)

	var r0 *manifest.Status
	if rf, ok := ret.Get(0).(func(context.Context) *manifest.Status); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*manifest.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStatusClient creates a new instance of StatusClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewStatusClient(t testing.TB) *StatusClient {
	mock := &StatusClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
