package business

import (
	"context"
	"ofspace-be/features/booking"
	"time"
)

type bookingBusiness struct {
	bookingData    booking.Data
	contextTimeout time.Duration
}

func NewBookingBusiness(bookingData booking.Data, timeout time.Duration) booking.Business {
	return &bookingBusiness{bookingData: bookingData, contextTimeout: timeout}
}

func (b *bookingBusiness) CreateBooking(c context.Context, data booking.Core) (booking.Core, error) {
	ctx, error1 := context.WithTimeout(c, b.contextTimeout)
	defer error1()
	data.UpdatedAt = time.Now()
	build, err2 := b.bookingData.CreateBooking(ctx, data)
	if err2 != nil {
		return booking.Core{}, err2
	}
	return build, nil
}

func (b *bookingBusiness) UpdateBooking(c context.Context, data booking.Core) (booking.Core, error) {
	ctx, error1 := context.WithTimeout(c, b.contextTimeout)
	defer error1()
	_, err := b.bookingData.GetOneBooking(ctx, data.CostumerId)
	if err != nil {
		return booking.Core{}, err
	}
	data.UpdatedAt = time.Now()
	up, err2 := b.bookingData.UpdateBooking(ctx, data)
	if err2 != nil {
		return booking.Core{}, err
	}
	return up, nil
}

func (b *bookingBusiness) GetAllBooking(c context.Context) ([]booking.Core, error) {
	ctx, error1 := context.WithTimeout(c, b.contextTimeout)
	defer error1()
	data, err := b.bookingData.GetAllBooking(ctx)
	if err != nil {
		return []booking.Core{}, err
	}
	return data, nil
}

func (b *bookingBusiness) GetAllBookingByBuilding(c context.Context, buildingId uint) ([]booking.Core, error) {
	ctx, error1 := context.WithTimeout(c, b.contextTimeout)
	defer error1()
	data, err := b.bookingData.GetAllBookingByBuilding(ctx, buildingId)
	if err != nil {
		return []booking.Core{}, err
	}
	return data, nil
}

func (b *bookingBusiness) GetOneBooking(c context.Context, userId uint) (booking.Core, error) {
	ctx, error1 := context.WithTimeout(c, b.contextTimeout)
	defer error1()
	data, err := b.bookingData.GetOneBooking(ctx, userId)
	if err != nil {
		return booking.Core{}, err
	}
	return data, nil
}

func (b bookingBusiness) SearchBookingByName(c context.Context, buildingId uint, name string) ([]booking.Core, error) {
	ctx, error1 := context.WithTimeout(c, b.contextTimeout)
	defer error1()
	data, err := b.bookingData.SearchBookingByName(ctx, buildingId, name)
	if err != nil {
		return []booking.Core{}, err
	}
	return data, nil
}

func (b bookingBusiness) SearchBookingByPayment(c context.Context, buildingId uint, paymentStatus string) ([]booking.Core, error) {
	ctx, error1 := context.WithTimeout(c, b.contextTimeout)
	defer error1()
	data, err := b.bookingData.SearchBookingByPayment(ctx, buildingId, paymentStatus)
	if err != nil {
		return []booking.Core{}, err
	}
	return data, nil
}

func (b bookingBusiness) FindBookingByDate(c context.Context, buildingId uint, startDate time.Time, endDate time.Time) ([]booking.Core, error) {
	ctx, error1 := context.WithTimeout(c, b.contextTimeout)
	defer error1()
	data, err := b.bookingData.FindBookingByDate(ctx, buildingId, startDate, endDate)
	if err != nil {
		return []booking.Core{}, err
	}
	return data, nil
}
