// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ExecPlatform is an autogenerated mock type for the ExecPlatform type
type ExecPlatform struct {
	mock.Mock
}

// ExecFunc provides a mock function with given fields:
func (_m *ExecPlatform) ExecFunc() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}