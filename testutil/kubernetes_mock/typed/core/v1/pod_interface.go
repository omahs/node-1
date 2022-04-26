// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	corev1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mock "github.com/stretchr/testify/mock"

	rest "k8s.io/client-go/rest"

	testing "testing"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/client-go/applyconfigurations/core/v1"

	v1beta1 "k8s.io/api/policy/v1beta1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// PodInterface is an autogenerated mock type for the PodInterface type
type PodInterface struct {
	mock.Mock
}

// Apply provides a mock function with given fields: ctx, pod, opts
func (_m *PodInterface) Apply(ctx context.Context, pod *v1.PodApplyConfiguration, opts metav1.ApplyOptions) (*corev1.Pod, error) {
	ret := _m.Called(ctx, pod, opts)

	var r0 *corev1.Pod
	if rf, ok := ret.Get(0).(func(context.Context, *v1.PodApplyConfiguration, metav1.ApplyOptions) *corev1.Pod); ok {
		r0 = rf(ctx, pod, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.PodApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, pod, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApplyStatus provides a mock function with given fields: ctx, pod, opts
func (_m *PodInterface) ApplyStatus(ctx context.Context, pod *v1.PodApplyConfiguration, opts metav1.ApplyOptions) (*corev1.Pod, error) {
	ret := _m.Called(ctx, pod, opts)

	var r0 *corev1.Pod
	if rf, ok := ret.Get(0).(func(context.Context, *v1.PodApplyConfiguration, metav1.ApplyOptions) *corev1.Pod); ok {
		r0 = rf(ctx, pod, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.PodApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, pod, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Bind provides a mock function with given fields: ctx, binding, opts
func (_m *PodInterface) Bind(ctx context.Context, binding *corev1.Binding, opts metav1.CreateOptions) error {
	ret := _m.Called(ctx, binding, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Binding, metav1.CreateOptions) error); ok {
		r0 = rf(ctx, binding, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: ctx, pod, opts
func (_m *PodInterface) Create(ctx context.Context, pod *corev1.Pod, opts metav1.CreateOptions) (*corev1.Pod, error) {
	ret := _m.Called(ctx, pod, opts)

	var r0 *corev1.Pod
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Pod, metav1.CreateOptions) *corev1.Pod); ok {
		r0 = rf(ctx, pod, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *corev1.Pod, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, pod, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *PodInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *PodInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Evict provides a mock function with given fields: ctx, eviction
func (_m *PodInterface) Evict(ctx context.Context, eviction *v1beta1.Eviction) error {
	ret := _m.Called(ctx, eviction)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.Eviction) error); ok {
		r0 = rf(ctx, eviction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *PodInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.Pod, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *corev1.Pod
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *corev1.Pod); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEphemeralContainers provides a mock function with given fields: ctx, podName, options
func (_m *PodInterface) GetEphemeralContainers(ctx context.Context, podName string, options metav1.GetOptions) (*corev1.EphemeralContainers, error) {
	ret := _m.Called(ctx, podName, options)

	var r0 *corev1.EphemeralContainers
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *corev1.EphemeralContainers); ok {
		r0 = rf(ctx, podName, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.EphemeralContainers)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, podName, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLogs provides a mock function with given fields: name, opts
func (_m *PodInterface) GetLogs(name string, opts *corev1.PodLogOptions) *rest.Request {
	ret := _m.Called(name, opts)

	var r0 *rest.Request
	if rf, ok := ret.Get(0).(func(string, *corev1.PodLogOptions) *rest.Request); ok {
		r0 = rf(name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Request)
		}
	}

	return r0
}

// List provides a mock function with given fields: ctx, opts
func (_m *PodInterface) List(ctx context.Context, opts metav1.ListOptions) (*corev1.PodList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *corev1.PodList
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *corev1.PodList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.PodList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *PodInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*corev1.Pod, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *corev1.Pod
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *corev1.Pod); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProxyGet provides a mock function with given fields: scheme, name, port, path, params
func (_m *PodInterface) ProxyGet(scheme string, name string, port string, path string, params map[string]string) rest.ResponseWrapper {
	ret := _m.Called(scheme, name, port, path, params)

	var r0 rest.ResponseWrapper
	if rf, ok := ret.Get(0).(func(string, string, string, string, map[string]string) rest.ResponseWrapper); ok {
		r0 = rf(scheme, name, port, path, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.ResponseWrapper)
		}
	}

	return r0
}

// Update provides a mock function with given fields: ctx, pod, opts
func (_m *PodInterface) Update(ctx context.Context, pod *corev1.Pod, opts metav1.UpdateOptions) (*corev1.Pod, error) {
	ret := _m.Called(ctx, pod, opts)

	var r0 *corev1.Pod
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Pod, metav1.UpdateOptions) *corev1.Pod); ok {
		r0 = rf(ctx, pod, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *corev1.Pod, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, pod, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEphemeralContainers provides a mock function with given fields: ctx, podName, ephemeralContainers, opts
func (_m *PodInterface) UpdateEphemeralContainers(ctx context.Context, podName string, ephemeralContainers *corev1.EphemeralContainers, opts metav1.UpdateOptions) (*corev1.EphemeralContainers, error) {
	ret := _m.Called(ctx, podName, ephemeralContainers, opts)

	var r0 *corev1.EphemeralContainers
	if rf, ok := ret.Get(0).(func(context.Context, string, *corev1.EphemeralContainers, metav1.UpdateOptions) *corev1.EphemeralContainers); ok {
		r0 = rf(ctx, podName, ephemeralContainers, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.EphemeralContainers)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *corev1.EphemeralContainers, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, podName, ephemeralContainers, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: ctx, pod, opts
func (_m *PodInterface) UpdateStatus(ctx context.Context, pod *corev1.Pod, opts metav1.UpdateOptions) (*corev1.Pod, error) {
	ret := _m.Called(ctx, pod, opts)

	var r0 *corev1.Pod
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Pod, metav1.UpdateOptions) *corev1.Pod); ok {
		r0 = rf(ctx, pod, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *corev1.Pod, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, pod, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *PodInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(ctx, opts)

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPodInterface creates a new instance of PodInterface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewPodInterface(t testing.TB) *PodInterface {
	mock := &PodInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
