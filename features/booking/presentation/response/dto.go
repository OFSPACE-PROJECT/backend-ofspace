package response

import (
	"ofspace-be/features/booking"
	"time"
)

type Booking struct {
	ID            uint      `json:"id"`
	CostumerId    uint      `json:"costumer_id"`
	User          User      `json:"user"`
	Building      Building  `json:"building"`
	ConsultantId  uint      `json:"consultant_id"`
	BuildingId    uint      `json:"building_id"`
	UnitId        uint      `json:"unit_id"`
	Unit          Unit      `json:"unit"`
	ConfirmedName string    `json:"confirmed_name"`
	TotalBought   uint      `json:"total_bought"`
	Price         float32   `json:"price"`
	DealDate      time.Time `json:"deal_date"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	PaymentStatus string    `json:"payment_status"`
	BookingStatus string    `json:"booking_status"`
}

type Building struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type Unit struct {
	ID       uint   `json:"id"`
	UnitType string `json:"unit_type"`
}

func FromBuildingCore(b booking.Building) Building {
	return Building{
		Id:   b.ID,
		Name: b.Name,
	}
}

func FromUnitCore(b booking.Unit) Unit {
	return Unit{
		ID:       b.ID,
		UnitType: b.UnitType,
	}
}

func FromUserCore(b booking.User) User {
	return User{
		Id:    b.ID,
		Name:  b.Name,
		Phone: b.Phone,
		Email: b.Email,
	}
}
func FromBookingCore(b booking.Core) Booking {
	return Booking{
		ID:            b.ID,
		CostumerId:    b.CostumerId,
		ConsultantId:  b.ConsultantId,
		BuildingId:    b.BuildingId,
		Unit:          FromUnitCore(b.Unit),
		User:          FromUserCore(b.User),
		Building:      FromBuildingCore(b.Building),
		UnitId:        b.UnitId,
		ConfirmedName: b.ConfirmedName,
		TotalBought:   b.TotalBought,
		Price:         b.Price,
		DealDate:      b.DealDate,
		StartDate:     b.StartDate,
		EndDate:       b.EndDate,
		PaymentStatus: b.PaymentStatus,
		BookingStatus: b.BookingStatus,
	}
}
func FromListBookingCore(core []booking.Core) (response []Booking) {
	for _, book := range core {
		response = append(response, FromBookingCore(book))
	}
	return
}
