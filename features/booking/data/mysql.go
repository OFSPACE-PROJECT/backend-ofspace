package data

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"ofspace-be/features/booking"
	"time"
)

type BookingData struct {
	Connect *gorm.DB
}

func NewBookingData(connect *gorm.DB) booking.Data {
	return &BookingData{
		Connect: connect,
	}
}

func (b *BookingData) CreateBooking(ctx context.Context, core booking.Core) (booking.Core, error) {
	booking1 := fromBookingCore(core)
	result := b.Connect.Create(&booking1)
	if result.Error != nil {
		return booking.Core{}, result.Error
	}
	return toBookingCore(booking1), nil
}

func (b *BookingData) UpdateBooking(ctx context.Context, core booking.Core) (booking.Core, error) {
	booking1 := fromBookingCore(core)
	result := b.Connect.Where("id= ?", booking1.ID).Updates(&Booking{
		ConfirmedName: booking1.ConfirmedName,
		TotalBought:   booking1.TotalBought,
		Price:         booking1.Price,
		DealDate:      booking1.DealDate,
		StartDate:     booking1.StartDate,
		EndDate:       booking1.EndDate,
		PaymentStatus: booking1.PaymentStatus,
		UpdatedAt:     booking1.UpdatedAt,
	})
	if result.Error != nil {
		return booking.Core{}, result.Error
	}
	return toBookingCore(booking1), nil
}

func (b *BookingData) GetAllBooking(ctx context.Context) ([]booking.Core, error) {
	var bookings []Booking
	result := b.Connect.Debug().Find(&bookings)

	if result.Error != nil {
		fmt.Println(result.Error)
		return []booking.Core{}, result.Error
	}
	return toSliceBookingCore(bookings), nil
}

func (b *BookingData) GetAllBookingByUnit(ctx context.Context, unitId uint) ([]booking.Core, error) {
	var bookings []Booking
	result := b.Connect.Debug().Find(&bookings, "unit_id", unitId)

	if result.Error != nil {
		fmt.Println(result.Error)
		return []booking.Core{}, result.Error
	}
	return toSliceBookingCore(bookings), nil
}

func (b *BookingData) GetOneBooking(ctx context.Context, userId uint) (booking.Core, error) {
	var booking1 Booking
	result := b.Connect.First(&booking1, "id= ?", userId)
	if result.Error != nil {
		return booking.Core{}, result.Error
	}
	return toBookingCore(booking1), nil
}

func (b *BookingData) SearchBookingByName(ctx context.Context, buildingId uint, name string) ([]booking.Core, error) {
	var booking1 []Booking
	result := b.Connect.Find(&booking1, "building_id= ? && confirmed_name= ?", buildingId, name)
	if result.Error != nil {
		return []booking.Core{}, result.Error
	}
	return toSliceBookingCore(booking1), nil
}

func (b *BookingData) SearchBookingByPayment(ctx context.Context, buildingId uint, paymentStatus string) ([]booking.Core, error) {
	var booking1 []Booking
	result := b.Connect.Find(&booking1, "building_id= ? && payment_status= ?", buildingId, paymentStatus)
	if result.Error != nil {
		return []booking.Core{}, result.Error
	}
	return toSliceBookingCore(booking1), nil
}

func (b *BookingData) GetBookingByStatus(ctx context.Context, buildingId uint, bookingStatus string) ([]booking.Core, error) {
	var booking1 []Booking
	result := b.Connect.Find(&booking1, "building_id= ? && booking_status= ?", buildingId, bookingStatus)
	if result.Error != nil {
		return []booking.Core{}, result.Error
	}
	return toSliceBookingCore(booking1), nil
}

func (b *BookingData) FindBookingByDate(ctx context.Context, buildingId uint, startDate time.Time, endDate time.Time) ([]booking.Core, error) {
	var booking1 []Booking
	result := b.Connect.Find(&booking1, "building_id= ? && start_date= ? && end_date= ?", buildingId, startDate, endDate)
	if result.Error != nil {
		return []booking.Core{}, result.Error
	}
	return toSliceBookingCore(booking1), nil
}

func (b *BookingData) GetSumOfTotalBoughtInUnit(ctx context.Context, unitId uint) (int, error) {
	var booking1 []Booking
	result := b.Connect.Find("booking_status='rented' && unit_id= ?", unitId)
	if result.Error != nil {
		return 0, result.Error
	}
	sum := 0
	for _, j := range booking1 {
		sum += int(j.TotalBought)
	}
	return sum, nil
}
