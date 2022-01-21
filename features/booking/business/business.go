package business

import (
	"context"
	"ofspace-be/features/booking"
	"ofspace-be/helpers/email"
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
	err := email.SendEmail(build)
	if err != nil {
		return booking.Core{}, err
	}
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
	if up.BookingStatus == "deal" || up.BookingStatus == "rented" {
		err = email.SendEmail(up)
		if err != nil {
			return booking.Core{}, err
		}
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

func (b *bookingBusiness) GetAllBookingByUnit(c context.Context, unitId uint) ([]booking.Core, error) {
	ctx, error1 := context.WithTimeout(c, b.contextTimeout)
	defer error1()
	data, err := b.bookingData.GetAllBookingByUnit(ctx, unitId)
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

func (b *bookingBusiness) SearchBookingByName(c context.Context, buildingId uint, name string) ([]booking.Core, error) {
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

func (b *bookingBusiness) GetBookingByStatus(c context.Context, buildingId uint, bookingStatus string) ([]booking.Core, error) {
	ctx, error1 := context.WithTimeout(c, b.contextTimeout)
	defer error1()
	data, err := b.bookingData.GetBookingByStatus(ctx, buildingId, bookingStatus)
	if err != nil {
		return []booking.Core{}, err
	}
	return data, nil
}

func (b *bookingBusiness) FindBookingByDate(c context.Context, buildingId uint, startDate time.Time, endDate time.Time) ([]booking.Core, error) {
	ctx, error1 := context.WithTimeout(c, b.contextTimeout)
	defer error1()
	data, err := b.bookingData.FindBookingByDate(ctx, buildingId, startDate, endDate)
	if err != nil {
		return []booking.Core{}, err
	}
	return data, nil
}

func (b *bookingBusiness) GetSumOfTotalBoughtInUnit(c context.Context, unitId uint) (int, error) {
	ctx, error1 := context.WithTimeout(c, b.contextTimeout)
	defer error1()
	fromBooking, err2 := b.bookingData.GetAllBookingByUnit(ctx, unitId)
	if err2 != nil {
		return 0, err2
	}
	totalSold := 0
	for _, j := range fromBooking {
		if j.BookingStatus == "rented" {
			totalSold += int(j.TotalBought)
		}
	}
	return totalSold, nil
}

func (b *bookingBusiness) GetEarningsInUnitWithDateFilter(c context.Context, unitId uint, startDate time.Time, endDate time.Time) (int, error) {
	ctx, error1 := context.WithTimeout(c, b.contextTimeout)
	defer error1()
	fromBooking, err2 := b.bookingData.GetAllBookingByUnit(ctx, unitId)
	if err2 != nil {
		return 0, err2
	}
	var totalEarning float32 = 0
	var canceled string = "canceled"
	for _, j := range fromBooking {
		if j.DealDate.After(startDate.Local()) && j.DealDate.Before(endDate.Local()) && j.BookingStatus != canceled {
			totalEarning += j.Price
		}
	}
	return int(totalEarning), nil
}

func (b *bookingBusiness) GetSumOfPaymentConfirmed(c context.Context, unitId uint, startDate time.Time, endDate time.Time) (int, error) {
	ctx, error1 := context.WithTimeout(c, b.contextTimeout)
	defer error1()
	fromBooking, err2 := b.bookingData.GetAllBookingByUnit(ctx, unitId)
	if err2 != nil {
		return 0, err2
	}
	var totalConfirm int = 0
	for _, j := range fromBooking {
		if j.DealDate.After(startDate) && j.DealDate.Before(endDate) && j.PaymentStatus == "confirmed" {
			totalConfirm += 1
		}
	}
	return int(totalConfirm), nil
}
