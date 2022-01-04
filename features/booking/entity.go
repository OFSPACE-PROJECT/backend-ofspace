package booking

import (
	"context"
	"time"
)

type Core struct {
	ID            uint
	CostumerId    string
	ConsultantId  string
	UnitType      string
	UnitId        []uint
	Price         int
	DealDate      time.Time
	StartDate     time.Time
	EndDate       time.Time
	PaymentStatus bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
type Business interface {
	GetAllBooking(ctx context.Context) ([]Core, error)
	CreateBooking(ctx context.Context, booking Core) (Core, error)
	UpdateBooking(ctx context.Context, costumerId uint, paymentStatus bool) (Core, error)
	SearchBookingByName(ctx context.Context, name string) ([]Core, error)
	SearchBookingByPayment(ctx context.Context, paymentStatus bool) ([]Core, error)
	FindBookingByDate(ctx context.Context, startDate time.Time, endDate time.Time) ([]Core, error)
}

type Data interface {
	GetAllBooking(ctx context.Context) ([]Core, error)
	CreateBooking(ctx context.Context, booking Core) (Core, error)
	UpdateBooking(ctx context.Context, costumerId uint, paymentStatus bool) (Core, error)
	SearchBookingByName(ctx context.Context, name string) ([]Core, error)
	SearchBookingByPayment(ctx context.Context, paymentStatus bool) ([]Core, error)
	FindBookingByDate(ctx context.Context, startDate time.Time, endDate time.Time) ([]Core, error)
}
