// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	corev1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mock "github.com/stretchr/testify/mock"

	testing "testing"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/client-go/applyconfigurations/core/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// PersistentVolumeClaimInterface is an autogenerated mock type for the PersistentVolumeClaimInterface type
type PersistentVolumeClaimInterface struct {
	mock.Mock
}

// Apply provides a mock function with given fields: ctx, persistentVolumeClaim, opts
func (_m *PersistentVolumeClaimInterface) Apply(ctx context.Context, persistentVolumeClaim *v1.PersistentVolumeClaimApplyConfiguration, opts metav1.ApplyOptions) (*corev1.PersistentVolumeClaim, error) {
	ret := _m.Called(ctx, persistentVolumeClaim, opts)

	var r0 *corev1.PersistentVolumeClaim
	if rf, ok := ret.Get(0).(func(context.Context, *v1.PersistentVolumeClaimApplyConfiguration, metav1.ApplyOptions) *corev1.PersistentVolumeClaim); ok {
		r0 = rf(ctx, persistentVolumeClaim, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.PersistentVolumeClaim)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.PersistentVolumeClaimApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, persistentVolumeClaim, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApplyStatus provides a mock function with given fields: ctx, persistentVolumeClaim, opts
func (_m *PersistentVolumeClaimInterface) ApplyStatus(ctx context.Context, persistentVolumeClaim *v1.PersistentVolumeClaimApplyConfiguration, opts metav1.ApplyOptions) (*corev1.PersistentVolumeClaim, error) {
	ret := _m.Called(ctx, persistentVolumeClaim, opts)

	var r0 *corev1.PersistentVolumeClaim
	if rf, ok := ret.Get(0).(func(context.Context, *v1.PersistentVolumeClaimApplyConfiguration, metav1.ApplyOptions) *corev1.PersistentVolumeClaim); ok {
		r0 = rf(ctx, persistentVolumeClaim, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.PersistentVolumeClaim)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.PersistentVolumeClaimApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, persistentVolumeClaim, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, persistentVolumeClaim, opts
func (_m *PersistentVolumeClaimInterface) Create(ctx context.Context, persistentVolumeClaim *corev1.PersistentVolumeClaim, opts metav1.CreateOptions) (*corev1.PersistentVolumeClaim, error) {
	ret := _m.Called(ctx, persistentVolumeClaim, opts)

	var r0 *corev1.PersistentVolumeClaim
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.PersistentVolumeClaim, metav1.CreateOptions) *corev1.PersistentVolumeClaim); ok {
		r0 = rf(ctx, persistentVolumeClaim, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.PersistentVolumeClaim)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *corev1.PersistentVolumeClaim, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, persistentVolumeClaim, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *PersistentVolumeClaimInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
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
func (_m *PersistentVolumeClaimInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *PersistentVolumeClaimInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.PersistentVolumeClaim, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *corev1.PersistentVolumeClaim
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *corev1.PersistentVolumeClaim); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.PersistentVolumeClaim)
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

// List provides a mock function with given fields: ctx, opts
func (_m *PersistentVolumeClaimInterface) List(ctx context.Context, opts metav1.ListOptions) (*corev1.PersistentVolumeClaimList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *corev1.PersistentVolumeClaimList
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *corev1.PersistentVolumeClaimList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.PersistentVolumeClaimList)
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
func (_m *PersistentVolumeClaimInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*corev1.PersistentVolumeClaim, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *corev1.PersistentVolumeClaim
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *corev1.PersistentVolumeClaim); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.PersistentVolumeClaim)
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

// Update provides a mock function with given fields: ctx, persistentVolumeClaim, opts
func (_m *PersistentVolumeClaimInterface) Update(ctx context.Context, persistentVolumeClaim *corev1.PersistentVolumeClaim, opts metav1.UpdateOptions) (*corev1.PersistentVolumeClaim, error) {
	ret := _m.Called(ctx, persistentVolumeClaim, opts)

	var r0 *corev1.PersistentVolumeClaim
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.PersistentVolumeClaim, metav1.UpdateOptions) *corev1.PersistentVolumeClaim); ok {
		r0 = rf(ctx, persistentVolumeClaim, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.PersistentVolumeClaim)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *corev1.PersistentVolumeClaim, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, persistentVolumeClaim, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: ctx, persistentVolumeClaim, opts
func (_m *PersistentVolumeClaimInterface) UpdateStatus(ctx context.Context, persistentVolumeClaim *corev1.PersistentVolumeClaim, opts metav1.UpdateOptions) (*corev1.PersistentVolumeClaim, error) {
	ret := _m.Called(ctx, persistentVolumeClaim, opts)

	var r0 *corev1.PersistentVolumeClaim
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.PersistentVolumeClaim, metav1.UpdateOptions) *corev1.PersistentVolumeClaim); ok {
		r0 = rf(ctx, persistentVolumeClaim, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.PersistentVolumeClaim)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *corev1.PersistentVolumeClaim, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, persistentVolumeClaim, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *PersistentVolumeClaimInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
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

// NewPersistentVolumeClaimInterface creates a new instance of PersistentVolumeClaimInterface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewPersistentVolumeClaimInterface(t testing.TB) *PersistentVolumeClaimInterface {
	mock := &PersistentVolumeClaimInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
