package response

import (
	"ofspace-be/features/booking"
	"time"
)

type Booking struct {
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
}

func FromBookingCore(b booking.Core) Booking {
	return Booking{
		ID:            b.ID,
		CostumerId:    b.CostumerId,
		ConsultantId:  b.ConsultantId,
		BuildingId:    b.BuildingId,
		UnitId:        b.UnitId,
		ConfirmedName: b.ConfirmedName,
		TotalBought:   b.TotalBought,
		Price:         b.Price,
		DealDate:      b.DealDate,
		StartDate:     b.StartDate,
		EndDate:       b.EndDate,
		PaymentStatus: b.PaymentStatus,
	}
}
func FromListBookingCore(core []booking.Core) (response []Booking) {
	for _, book := range core {
		response = append(response, FromBookingCore(book))
	}
	return
}
