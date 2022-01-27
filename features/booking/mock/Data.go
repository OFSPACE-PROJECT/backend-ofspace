// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"
	booking "ofspace-be/features/booking"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Data is an autogenerated mock type for the Data type
type Data struct {
	mock.Mock
}

// CreateBooking provides a mock function with given fields: ctx, _a1
func (_m *Data) CreateBooking(ctx context.Context, _a1 booking.Core) (booking.Core, error) {
	ret := _m.Called(ctx, _a1)

	var r0 booking.Core
	if rf, ok := ret.Get(0).(func(context.Context, booking.Core) booking.Core); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(booking.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, booking.Core) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindBookingByDate provides a mock function with given fields: ctx, buildingId, startDate, endDate
func (_m *Data) FindBookingByDate(ctx context.Context, buildingId uint, startDate time.Time, endDate time.Time) ([]booking.Core, error) {
	ret := _m.Called(ctx, buildingId, startDate, endDate)

	var r0 []booking.Core
	if rf, ok := ret.Get(0).(func(context.Context, uint, time.Time, time.Time) []booking.Core); ok {
		r0 = rf(ctx, buildingId, startDate, endDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]booking.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, time.Time, time.Time) error); ok {
		r1 = rf(ctx, buildingId, startDate, endDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllBooking provides a mock function with given fields: ctx
func (_m *Data) GetAllBooking(ctx context.Context) ([]booking.Core, error) {
	ret := _m.Called(ctx)

	var r0 []booking.Core
	if rf, ok := ret.Get(0).(func(context.Context) []booking.Core); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]booking.Core)
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

// GetAllBookingByBuilding provides a mock function with given fields: ctx, buildingId
func (_m *Data) GetAllBookingByBuilding(ctx context.Context, buildingId uint) ([]booking.Core, error) {
	ret := _m.Called(ctx, buildingId)

	var r0 []booking.Core
	if rf, ok := ret.Get(0).(func(context.Context, uint) []booking.Core); ok {
		r0 = rf(ctx, buildingId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]booking.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, buildingId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllBookingByUnit provides a mock function with given fields: ctx, unitId
func (_m *Data) GetAllBookingByUnit(ctx context.Context, unitId uint) ([]booking.Core, error) {
	ret := _m.Called(ctx, unitId)

	var r0 []booking.Core
	if rf, ok := ret.Get(0).(func(context.Context, uint) []booking.Core); ok {
		r0 = rf(ctx, unitId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]booking.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, unitId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllBookingByUser provides a mock function with given fields: ctx, userId
func (_m *Data) GetAllBookingByUser(ctx context.Context, userId uint) ([]booking.Core, error) {
	ret := _m.Called(ctx, userId)

	var r0 []booking.Core
	if rf, ok := ret.Get(0).(func(context.Context, uint) []booking.Core); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]booking.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBookingByStatus provides a mock function with given fields: ctx, buildingId, bookingStatus
func (_m *Data) GetBookingByStatus(ctx context.Context, buildingId uint, bookingStatus string) ([]booking.Core, error) {
	ret := _m.Called(ctx, buildingId, bookingStatus)

	var r0 []booking.Core
	if rf, ok := ret.Get(0).(func(context.Context, uint, string) []booking.Core); ok {
		r0 = rf(ctx, buildingId, bookingStatus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]booking.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, string) error); ok {
		r1 = rf(ctx, buildingId, bookingStatus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneBooking provides a mock function with given fields: ctx, Id
func (_m *Data) GetOneBooking(ctx context.Context, Id uint) (booking.Core, error) {
	ret := _m.Called(ctx, Id)

	var r0 booking.Core
	if rf, ok := ret.Get(0).(func(context.Context, uint) booking.Core); ok {
		r0 = rf(ctx, Id)
	} else {
		r0 = ret.Get(0).(booking.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, Id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchBookingByName provides a mock function with given fields: ctx, buildingId, name
func (_m *Data) SearchBookingByName(ctx context.Context, buildingId uint, name string) ([]booking.Core, error) {
	ret := _m.Called(ctx, buildingId, name)

	var r0 []booking.Core
	if rf, ok := ret.Get(0).(func(context.Context, uint, string) []booking.Core); ok {
		r0 = rf(ctx, buildingId, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]booking.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, string) error); ok {
		r1 = rf(ctx, buildingId, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchBookingByPayment provides a mock function with given fields: ctx, buildingId, paymentStatus
func (_m *Data) SearchBookingByPayment(ctx context.Context, buildingId uint, paymentStatus string) ([]booking.Core, error) {
	ret := _m.Called(ctx, buildingId, paymentStatus)

	var r0 []booking.Core
	if rf, ok := ret.Get(0).(func(context.Context, uint, string) []booking.Core); ok {
		r0 = rf(ctx, buildingId, paymentStatus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]booking.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, string) error); ok {
		r1 = rf(ctx, buildingId, paymentStatus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBooking provides a mock function with given fields: ctx, _a1
func (_m *Data) UpdateBooking(ctx context.Context, _a1 booking.Core) (booking.Core, error) {
	ret := _m.Called(ctx, _a1)

	var r0 booking.Core
	if rf, ok := ret.Get(0).(func(context.Context, booking.Core) booking.Core); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(booking.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, booking.Core) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}