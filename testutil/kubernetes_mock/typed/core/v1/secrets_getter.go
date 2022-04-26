// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// SecretsGetter is an autogenerated mock type for the SecretsGetter type
type SecretsGetter struct {
	mock.Mock
}

// Secrets provides a mock function with given fields: namespace
func (_m *SecretsGetter) Secrets(namespace string) v1.SecretInterface {
	ret := _m.Called(namespace)

	var r0 v1.SecretInterface
	if rf, ok := ret.Get(0).(func(string) v1.SecretInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.SecretInterface)
		}
	}

	return r0
}

// NewSecretsGetter creates a new instance of SecretsGetter. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewSecretsGetter(t testing.TB) *SecretsGetter {
	mock := &SecretsGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
