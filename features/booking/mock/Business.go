// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"
	booking "ofspace-be/features/booking"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Business is an autogenerated mock type for the Business type
type Business struct {
	mock.Mock
}

// CreateBooking provides a mock function with given fields: ctx, _a1
func (_m *Business) CreateBooking(ctx context.Context, _a1 booking.Core) (booking.Core, error) {
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
func (_m *Business) FindBookingByDate(ctx context.Context, buildingId uint, startDate time.Time, endDate time.Time) ([]booking.Core, error) {
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
func (_m *Business) GetAllBooking(ctx context.Context) ([]booking.Core, error) {
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
func (_m *Business) GetAllBookingByBuilding(ctx context.Context, buildingId uint) ([]booking.Core, error) {
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
func (_m *Business) GetAllBookingByUnit(ctx context.Context, unitId uint) ([]booking.Core, error) {
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
func (_m *Business) GetAllBookingByUser(ctx context.Context, userId uint) ([]booking.Core, error) {
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
func (_m *Business) GetBookingByStatus(ctx context.Context, buildingId uint, bookingStatus string) ([]booking.Core, error) {
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

// GetEarningsInUnitWithDateFilter provides a mock function with given fields: ctx, unitId, startDate, endDate
func (_m *Business) GetEarningsInUnitWithDateFilter(ctx context.Context, unitId uint, startDate time.Time, endDate time.Time) (int, error) {
	ret := _m.Called(ctx, unitId, startDate, endDate)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, uint, time.Time, time.Time) int); ok {
		r0 = rf(ctx, unitId, startDate, endDate)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, time.Time, time.Time) error); ok {
		r1 = rf(ctx, unitId, startDate, endDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneBooking provides a mock function with given fields: ctx, Id
func (_m *Business) GetOneBooking(ctx context.Context, Id uint) (booking.Core, error) {
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

// GetSumOfPaymentConfirmed provides a mock function with given fields: ctx, unitId, startDate, endDate
func (_m *Business) GetSumOfPaymentConfirmed(ctx context.Context, unitId uint, startDate time.Time, endDate time.Time) (int, error) {
	ret := _m.Called(ctx, unitId, startDate, endDate)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, uint, time.Time, time.Time) int); ok {
		r0 = rf(ctx, unitId, startDate, endDate)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, time.Time, time.Time) error); ok {
		r1 = rf(ctx, unitId, startDate, endDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSumOfTotalBoughtInUnit provides a mock function with given fields: ctx, unitId
func (_m *Business) GetSumOfTotalBoughtInUnit(ctx context.Context, unitId uint) (int, error) {
	ret := _m.Called(ctx, unitId)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, uint) int); ok {
		r0 = rf(ctx, unitId)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, unitId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchBookingByName provides a mock function with given fields: ctx, buildingId, name
func (_m *Business) SearchBookingByName(ctx context.Context, buildingId uint, name string) ([]booking.Core, error) {
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
func (_m *Business) SearchBookingByPayment(ctx context.Context, buildingId uint, paymentStatus string) ([]booking.Core, error) {
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
func (_m *Business) UpdateBooking(ctx context.Context, _a1 booking.Core) (booking.Core, error) {
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