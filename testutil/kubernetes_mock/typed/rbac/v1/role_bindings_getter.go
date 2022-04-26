// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	v1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

// RoleBindingsGetter is an autogenerated mock type for the RoleBindingsGetter type
type RoleBindingsGetter struct {
	mock.Mock
}

// RoleBindings provides a mock function with given fields: namespace
func (_m *RoleBindingsGetter) RoleBindings(namespace string) v1.RoleBindingInterface {
	ret := _m.Called(namespace)

	var r0 v1.RoleBindingInterface
	if rf, ok := ret.Get(0).(func(string) v1.RoleBindingInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.RoleBindingInterface)
		}
	}

	return r0
}

// NewRoleBindingsGetter creates a new instance of RoleBindingsGetter. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRoleBindingsGetter(t testing.TB) *RoleBindingsGetter {
	mock := &RoleBindingsGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
