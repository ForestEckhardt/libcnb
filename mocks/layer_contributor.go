// Code generated by mockery 2.9.4. DO NOT EDIT.

package mocks

import (
	libcnb "github.com/buildpacks/libcnb"
	mock "github.com/stretchr/testify/mock"
)

// LayerContributor is an autogenerated mock type for the LayerContributor type
type LayerContributor struct {
	mock.Mock
}

// Contribute provides a mock function with given fields: layer
func (_m *LayerContributor) Contribute(layer libcnb.Layer) (libcnb.Layer, error) {
	ret := _m.Called(layer)

	var r0 libcnb.Layer
	if rf, ok := ret.Get(0).(func(libcnb.Layer) libcnb.Layer); ok {
		r0 = rf(layer)
	} else {
		r0 = ret.Get(0).(libcnb.Layer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(libcnb.Layer) error); ok {
		r1 = rf(layer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Name provides a mock function with given fields:
func (_m *LayerContributor) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
