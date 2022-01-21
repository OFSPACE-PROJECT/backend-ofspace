package data

import (
	"ofspace-be/features/booking"
	"time"
)

type Booking struct {
	ID            uint `gorm:"primaryKey"`
	CostumerId    uint
	ConsultantId  uint
	User          User `gorm:"foreignKey:CostumerId"`
	BuildingId    uint
	UnitId        uint
	ConfirmedName string
	TotalBought   uint
	Price         float32
	DealDate      time.Time
	StartDate     time.Time
	EndDate       time.Time
	PaymentStatus string `gorm:"default:unconfirmed"`
	BookingStatus string `gorm:"default:deal"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type User struct {
	ID    uint
	Name  string
	Email string
}

func toUserCore(u User) booking.User {
	return booking.User{
		ID:   u.ID,
		Name: u.Name,
	}
}

func fromBookingCore(c booking.Core) Booking {
	return Booking{
		ID:            c.ID,
		CostumerId:    c.CostumerId,
		ConsultantId:  c.ConsultantId,
		BuildingId:    c.BuildingId,
		UnitId:        c.UnitId,
		ConfirmedName: c.ConfirmedName,
		TotalBought:   c.TotalBought,
		Price:         c.Price,
		DealDate:      c.DealDate,
		StartDate:     c.StartDate,
		EndDate:       c.EndDate,
		PaymentStatus: c.PaymentStatus,
		BookingStatus: c.BookingStatus,
		CreatedAt:     c.CreatedAt,
		UpdatedAt:     c.UpdatedAt,
	}
}

func toBookingCore(c Booking) booking.Core {
	return booking.Core{
		ID:            c.ID,
		CostumerId:    c.CostumerId,
		User:          toUserCore(c.User),
		ConsultantId:  c.ConsultantId,
		BuildingId:    c.BuildingId,
		UnitId:        c.UnitId,
		ConfirmedName: c.ConfirmedName,
		TotalBought:   c.TotalBought,
		Price:         c.Price,
		DealDate:      c.DealDate,
		StartDate:     c.StartDate,
		EndDate:       c.EndDate,
		PaymentStatus: c.PaymentStatus,
		BookingStatus: c.BookingStatus,
		CreatedAt:     c.CreatedAt,
		UpdatedAt:     c.UpdatedAt,
	}
}

func toSliceBookingCore(bookings []Booking) (result []booking.Core) {
	for _, booking1 := range bookings {
		result = append(result, toBookingCore(booking1))
	}
	return
}
