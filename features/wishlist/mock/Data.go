// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"
	wishlist "ofspace-be/features/wishlist"

	mock "github.com/stretchr/testify/mock"
)

// Data is an autogenerated mock type for the Data type
type Data struct {
	mock.Mock
}

// CreateWishlist provides a mock function with given fields: ctx, _a1
func (_m *Data) CreateWishlist(ctx context.Context, _a1 wishlist.Core) (wishlist.Core, error) {
	ret := _m.Called(ctx, _a1)

	var r0 wishlist.Core
	if rf, ok := ret.Get(0).(func(context.Context, wishlist.Core) wishlist.Core); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(wishlist.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, wishlist.Core) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteWishlist provides a mock function with given fields: ctx, wishlistId
func (_m *Data) DeleteWishlist(ctx context.Context, wishlistId uint) (wishlist.Core, error) {
	ret := _m.Called(ctx, wishlistId)

	var r0 wishlist.Core
	if rf, ok := ret.Get(0).(func(context.Context, uint) wishlist.Core); ok {
		r0 = rf(ctx, wishlistId)
	} else {
		r0 = ret.Get(0).(wishlist.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, wishlistId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllWishlists provides a mock function with given fields: ctx, UserId
func (_m *Data) GetAllWishlists(ctx context.Context, UserId uint) ([]wishlist.Core, error) {
	ret := _m.Called(ctx, UserId)

	var r0 []wishlist.Core
	if rf, ok := ret.Get(0).(func(context.Context, uint) []wishlist.Core); ok {
		r0 = rf(ctx, UserId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]wishlist.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, UserId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWishlist provides a mock function with given fields: ctx, wishlistId
func (_m *Data) GetWishlist(ctx context.Context, wishlistId uint) (wishlist.Core, error) {
	ret := _m.Called(ctx, wishlistId)

	var r0 wishlist.Core
	if rf, ok := ret.Get(0).(func(context.Context, uint) wishlist.Core); ok {
		r0 = rf(ctx, wishlistId)
	} else {
		r0 = ret.Get(0).(wishlist.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, wishlistId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}