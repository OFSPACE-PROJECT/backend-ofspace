package request

import (
	"ofspace-be/features/booking"
	"time"
)

type CreateBooking struct {
	CostumerId    uint      `json:"costumer_id"`
	ConsultantId  uint      `json:"consultant_id"`
	BuildingId    uint      `json:"building_id"`
	UnitId        uint      `json:"unit_id"`
	ConfirmedName string    `json:"confirmed_name"`
	TotalBought   uint      `json:"total_bought"`
	Price         float32   `json:"price"`
	DealDate      time.Time `json:"deal_date"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	PaymentStatus string    `json:"payment_status"`
}
type UpdateBooking struct {
	ID            uint      `json:"id"`
	CostumerId    uint      `json:"costumer_id"`
	ConsultantId  uint      `json:"consultant_id"`
	BuildingId    uint      `json:"building_id"`
	UnitId        uint      `json:"unit_id"`
	ConfirmedName string    `json:"confirmed_name"`
	TotalBought   uint      `json:"total_bought"`
	Price         float32   `json:"price"`
	DealDate      time.Time `json:"deal_date"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	PaymentStatus string    `json:"payment_status"`
	BookingStatus string    `json:"booking_status"`
}

func (c *CreateBooking) ToCore() booking.Core {
	return booking.Core{
		CostumerId:    c.BuildingId,
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
	}
}

func (c *UpdateBooking) ToUpdateCore() booking.Core {
	return booking.Core{
		ID:            c.BuildingId,
		CostumerId:    c.BuildingId,
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
	}
}
