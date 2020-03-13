// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	libcnb "github.com/buildpacks/libcnb"
	mock "github.com/stretchr/testify/mock"
)

// Detector is an autogenerated mock type for the Detector type
type Detector struct {
	mock.Mock
}

// Detect provides a mock function with given fields: context
func (_m *Detector) Detect(context libcnb.DetectContext) (libcnb.DetectResult, error) {
	ret := _m.Called(context)

	var r0 libcnb.DetectResult
	if rf, ok := ret.Get(0).(func(libcnb.DetectContext) libcnb.DetectResult); ok {
		r0 = rf(context)
	} else {
		r0 = ret.Get(0).(libcnb.DetectResult)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(libcnb.DetectContext) error); ok {
		r1 = rf(context)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
