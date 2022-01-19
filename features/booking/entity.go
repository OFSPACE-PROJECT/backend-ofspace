package booking

import (
	"context"
	"time"
)

type Core struct {
	ID            uint
	CostumerId    uint
	ConsultantId  uint
	BuildingId    uint
	UnitId        uint
	ConfirmedName string
	TotalBought   uint
	Price         float32
	DealDate      time.Time
	StartDate     time.Time
	EndDate       time.Time
	PaymentStatus string
	BookingStatus string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
type Business interface {
	GetAllBooking(ctx context.Context) ([]Core, error)
	GetAllBookingByUnit(ctx context.Context, unitId uint) ([]Core, error)
	GetOneBooking(ctx context.Context, userId uint) (Core, error)
	CreateBooking(ctx context.Context, booking Core) (Core, error)
	UpdateBooking(ctx context.Context, booking Core) (Core, error)
	SearchBookingByName(ctx context.Context, buildingId uint, name string) ([]Core, error)
	SearchBookingByPayment(ctx context.Context, buildingId uint, paymentStatus string) ([]Core, error)
	GetBookingByStatus(ctx context.Context, buildingId uint, bookingStatus string) ([]Core, error)
	FindBookingByDate(ctx context.Context, buildingId uint, startDate time.Time, endDate time.Time) ([]Core, error)
	GetSumOfTotalBoughtInUnit(ctx context.Context, unitId uint) (int, error)
}

type Data interface {
	GetAllBooking(ctx context.Context) ([]Core, error)
	GetBookingByStatus(ctx context.Context, buildingId uint, bookingStatus string) ([]Core, error)
	GetAllBookingByUnit(ctx context.Context, unitId uint) ([]Core, error)
	GetOneBooking(ctx context.Context, userId uint) (Core, error)
	CreateBooking(ctx context.Context, booking Core) (Core, error)
	UpdateBooking(ctx context.Context, booking Core) (Core, error)
	SearchBookingByName(ctx context.Context, buildingId uint, name string) ([]Core, error)
	SearchBookingByPayment(ctx context.Context, buildingId uint, paymentStatus string) ([]Core, error)
	FindBookingByDate(ctx context.Context, buildingId uint, startDate time.Time, endDate time.Time) ([]Core, error)
	GetSumOfTotalBoughtInUnit(ctx context.Context, unitId uint) (int, error)
}
