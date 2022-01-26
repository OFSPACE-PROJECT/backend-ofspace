package data

import (
	"ofspace-be/features/booking"
	"time"
)

type Booking struct {
	ID            uint `gorm:"primaryKey"`
	CostumerId    uint
	ConsultantId  uint
	User          User `gorm:"foreignKey:CostumerId;omitempty"`
	BuildingId    uint
	Building      Building `gorm:"foreignKey:BuildingId;omitempty"`
	UnitId        uint
	Unit          Unit `gorm:"foreignKey:UnitId;omitempty"`
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
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
	Phone string
}

type Unit struct {
	ID       uint `gorm:"primaryKey"`
	UnitType string
}

type Building struct {
	ID   uint
	Name string
}

func toUserCore(u User) booking.User {
	return booking.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Phone: u.Phone,
	}
}
func toBuildingCore(u Building) booking.Building {
	return booking.Building{
		ID:   u.ID,
		Name: u.Name,
	}
}
func toUnitCore(u Unit) booking.Unit {
	return booking.Unit{
		ID:       u.ID,
		UnitType: u.UnitType,
	}
}
func fromBuildingCore(u booking.Building) Building {
	return Building{
		ID:   u.ID,
		Name: u.Name,
	}
}
func fromUnitCore(u booking.Unit) Unit {
	return Unit{
		ID:       u.ID,
		UnitType: u.UnitType,
	}
}
func fromBookingCore(c *booking.Core) Booking {
	return Booking{
		ID:            c.ID,
		CostumerId:    c.CostumerId,
		ConsultantId:  c.ConsultantId,
		BuildingId:    c.BuildingId,
		Building:      fromBuildingCore(c.Building),
		UnitId:        c.UnitId,
		Unit:          fromUnitCore(c.Unit),
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
		Unit:          toUnitCore(c.Unit),
		ConfirmedName: c.ConfirmedName,
		Building:      toBuildingCore(c.Building),
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
