// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	discoveryv1 "k8s.io/api/discovery/v1"

	mock "github.com/stretchr/testify/mock"

	v1 "k8s.io/api/core/v1"
)

// ObjectCacheInterface is an autogenerated mock type for the ObjectCacheInterface type
type ObjectCacheInterface struct {
	mock.Mock
}

// GetAllPods provides a mock function with given fields:
func (_m *ObjectCacheInterface) GetAllPods() ([]*v1.Pod, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllPods")
	}

	var r0 []*v1.Pod
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*v1.Pod, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*v1.Pod); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.Pod)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNamespace provides a mock function with given fields: name
func (_m *ObjectCacheInterface) GetNamespace(name string) (*v1.Namespace, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for GetNamespace")
	}

	var r0 *v1.Namespace
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*v1.Namespace, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *v1.Namespace); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Namespace)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNamespaces provides a mock function with given fields:
func (_m *ObjectCacheInterface) GetNamespaces() ([]*v1.Namespace, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNamespaces")
	}

	var r0 []*v1.Namespace
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*v1.Namespace, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*v1.Namespace); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.Namespace)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNode provides a mock function with given fields: name
func (_m *ObjectCacheInterface) GetNode(name string) (*v1.Node, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for GetNode")
	}

	var r0 *v1.Node
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*v1.Node, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *v1.Node); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodes provides a mock function with given fields:
func (_m *ObjectCacheInterface) GetNodes() ([]*v1.Node, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNodes")
	}

	var r0 []*v1.Node
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*v1.Node, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*v1.Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.Node)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPod provides a mock function with given fields: namespace, name
func (_m *ObjectCacheInterface) GetPod(namespace string, name string) (*v1.Pod, error) {
	ret := _m.Called(namespace, name)

	if len(ret) == 0 {
		panic("no return value specified for GetPod")
	}

	var r0 *v1.Pod
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*v1.Pod, error)); ok {
		return rf(namespace, name)
	}
	if rf, ok := ret.Get(0).(func(string, string) *v1.Pod); ok {
		r0 = rf(namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Pod)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPods provides a mock function with given fields: namespace
func (_m *ObjectCacheInterface) GetPods(namespace string) ([]*v1.Pod, error) {
	ret := _m.Called(namespace)

	if len(ret) == 0 {
		panic("no return value specified for GetPods")
	}

	var r0 []*v1.Pod
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*v1.Pod, error)); ok {
		return rf(namespace)
	}
	if rf, ok := ret.Get(0).(func(string) []*v1.Pod); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.Pod)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetService provides a mock function with given fields: namespace, name
func (_m *ObjectCacheInterface) GetService(namespace string, name string) (*v1.Service, error) {
	ret := _m.Called(namespace, name)

	if len(ret) == 0 {
		panic("no return value specified for GetService")
	}

	var r0 *v1.Service
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*v1.Service, error)); ok {
		return rf(namespace, name)
	}
	if rf, ok := ret.Get(0).(func(string, string) *v1.Service); ok {
		r0 = rf(namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Service)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServiceEndpointSlices provides a mock function with given fields: namespace, svcName, network
func (_m *ObjectCacheInterface) GetServiceEndpointSlices(namespace string, svcName string, network string) ([]*discoveryv1.EndpointSlice, error) {
	ret := _m.Called(namespace, svcName, network)

	if len(ret) == 0 {
		panic("no return value specified for GetServiceEndpointSlices")
	}

	var r0 []*discoveryv1.EndpointSlice
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) ([]*discoveryv1.EndpointSlice, error)); ok {
		return rf(namespace, svcName, network)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) []*discoveryv1.EndpointSlice); ok {
		r0 = rf(namespace, svcName, network)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*discoveryv1.EndpointSlice)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(namespace, svcName, network)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewObjectCacheInterface creates a new instance of ObjectCacheInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewObjectCacheInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ObjectCacheInterface {
	mock := &ObjectCacheInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
