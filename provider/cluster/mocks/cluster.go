// Code generated by mockery 2.12.1. DO NOT EDIT.

package mocks

import (
	clustertypesv1beta2 "github.com/ovrclk/akash/provider/cluster/types/v1beta2"
	mock "github.com/stretchr/testify/mock"

	testing "testing"

	typesv1beta2 "github.com/ovrclk/akash/types/v1beta2"

	v1beta2 "github.com/ovrclk/akash/x/market/types/v1beta2"
)

// Cluster is an autogenerated mock type for the Cluster type
type Cluster struct {
	mock.Mock
}

// Reserve provides a mock function with given fields: _a0, _a1
func (_m *Cluster) Reserve(_a0 v1beta2.OrderID, _a1 typesv1beta2.ResourceGroup) (clustertypesv1beta2.Reservation, error) {
	ret := _m.Called(_a0, _a1)

	var r0 clustertypesv1beta2.Reservation
	if rf, ok := ret.Get(0).(func(v1beta2.OrderID, typesv1beta2.ResourceGroup) clustertypesv1beta2.Reservation); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clustertypesv1beta2.Reservation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(v1beta2.OrderID, typesv1beta2.ResourceGroup) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unreserve provides a mock function with given fields: _a0
func (_m *Cluster) Unreserve(_a0 v1beta2.OrderID) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(v1beta2.OrderID) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCluster creates a new instance of Cluster. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewCluster(t testing.TB) *Cluster {
	mock := &Cluster{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
