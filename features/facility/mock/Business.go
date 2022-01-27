// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"
	facility "ofspace-be/features/facility"

	mock "github.com/stretchr/testify/mock"
)

// Business is an autogenerated mock type for the Business type
type Business struct {
	mock.Mock
}

// CreateFacility provides a mock function with given fields: ctx, core
func (_m *Business) CreateFacility(ctx context.Context, core facility.Facility) (facility.Facility, error) {
	ret := _m.Called(ctx, core)

	var r0 facility.Facility
	if rf, ok := ret.Get(0).(func(context.Context, facility.Facility) facility.Facility); ok {
		r0 = rf(ctx, core)
	} else {
		r0 = ret.Get(0).(facility.Facility)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, facility.Facility) error); ok {
		r1 = rf(ctx, core)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFacility provides a mock function with given fields: ctx, id
func (_m *Business) DeleteFacility(ctx context.Context, id uint) (facility.Facility, error) {
	ret := _m.Called(ctx, id)

	var r0 facility.Facility
	if rf, ok := ret.Get(0).(func(context.Context, uint) facility.Facility); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(facility.Facility)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllFacility provides a mock function with given fields: ctx
func (_m *Business) GetAllFacility(ctx context.Context) ([]facility.Facility, error) {
	ret := _m.Called(ctx)

	var r0 []facility.Facility
	if rf, ok := ret.Get(0).(func(context.Context) []facility.Facility); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]facility.Facility)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFacility provides a mock function with given fields: ctx, id
func (_m *Business) GetFacility(ctx context.Context, id uint) (facility.Facility, error) {
	ret := _m.Called(ctx, id)

	var r0 facility.Facility
	if rf, ok := ret.Get(0).(func(context.Context, uint) facility.Facility); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(facility.Facility)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchFacility provides a mock function with given fields: ctx, name
func (_m *Business) SearchFacility(ctx context.Context, name string) ([]facility.Facility, error) {
	ret := _m.Called(ctx, name)

	var r0 []facility.Facility
	if rf, ok := ret.Get(0).(func(context.Context, string) []facility.Facility); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]facility.Facility)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFacility provides a mock function with given fields: ctx, core
func (_m *Business) UpdateFacility(ctx context.Context, core facility.Facility) (facility.Facility, error) {
	ret := _m.Called(ctx, core)

	var r0 facility.Facility
	if rf, ok := ret.Get(0).(func(context.Context, facility.Facility) facility.Facility); ok {
		r0 = rf(ctx, core)
	} else {
		r0 = ret.Get(0).(facility.Facility)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, facility.Facility) error); ok {
		r1 = rf(ctx, core)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
