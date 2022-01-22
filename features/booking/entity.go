package booking

import (
	"context"
	"ofspace-be/features/unit"
	"time"
)

type Core struct {
	ID            uint
	CostumerId    uint
	ConsultantId  uint
	User          User
	BuildingId    uint
	UnitId        uint
	Unit          unit.Core
	Building      Building
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

type User struct {
	ID    uint
	Name  string
	Email string
}

type Building struct {
	ID   uint
	Name string
}

type Business interface {
	GetAllBooking(ctx context.Context) ([]Core, error)
	GetAllBookingByUnit(ctx context.Context, unitId uint) ([]Core, error)
	GetAllBookingByUser(ctx context.Context, userId uint) ([]Core, error)
	GetOneBooking(ctx context.Context, Id uint) (Core, error)
	CreateBooking(ctx context.Context, booking Core) (Core, error)
	UpdateBooking(ctx context.Context, booking Core) (Core, error)
	SearchBookingByName(ctx context.Context, buildingId uint, name string) ([]Core, error)
	SearchBookingByPayment(ctx context.Context, buildingId uint, paymentStatus string) ([]Core, error)
	GetBookingByStatus(ctx context.Context, buildingId uint, bookingStatus string) ([]Core, error)
	FindBookingByDate(ctx context.Context, buildingId uint, startDate time.Time, endDate time.Time) ([]Core, error)
	GetSumOfTotalBoughtInUnit(ctx context.Context, unitId uint) (int, error)
	GetEarningsInUnitWithDateFilter(ctx context.Context, unitId uint, startDate time.Time, endDate time.Time) (int, error)
	GetSumOfPaymentConfirmed(ctx context.Context, unitId uint, startDate time.Time, endDate time.Time) (int, error)
}

type Data interface {
	GetAllBooking(ctx context.Context) ([]Core, error)
	GetBookingByStatus(ctx context.Context, buildingId uint, bookingStatus string) ([]Core, error)
	GetAllBookingByUnit(ctx context.Context, unitId uint) ([]Core, error)
	GetAllBookingByUser(ctx context.Context, userId uint) ([]Core, error)
	GetOneBooking(ctx context.Context, Id uint) (Core, error)
	CreateBooking(ctx context.Context, booking Core) (Core, error)
	UpdateBooking(ctx context.Context, booking Core) (Core, error)
	SearchBookingByName(ctx context.Context, buildingId uint, name string) ([]Core, error)
	SearchBookingByPayment(ctx context.Context, buildingId uint, paymentStatus string) ([]Core, error)
	FindBookingByDate(ctx context.Context, buildingId uint, startDate time.Time, endDate time.Time) ([]Core, error)

	//GetSumOfTotalBoughtInUnit(ctx context.Context, unitId uint) (int, error)
	//GetEarningsInUnitWithDateFilter(ctx context.Context, unitId uint, startDate time.Time, endDate time.Time) (int, error)
}
