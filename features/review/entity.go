package review

import (
	"context"
	"time"
)

type Core struct {
	Id                    uint
	CustomerId            uint
	Customer              User
	UnitId                uint
	BuildingId            uint
	Unit                  Unit
	BookingId             uint
	RatingAccess          float32
	RatingFacility        float32
	RatingManagement      float32
	RatingQuality         float32
	CostumerOverallRating float32
	Comment               string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

type User struct {
	ID   uint
	Name string
}

type Unit struct {
	Id          uint
	BuildingId  uint
	Description string
	UnitType    string
}

type Business interface {
	CreateReview(ctx context.Context, review Core) (Core, error)
	GetAllReview(ctx context.Context, unitType uint) ([]Core, error)
	//GetBuildingReview(ctx context.Context, unitType uint) ([]Core, error)
	GetOneReview(ctx context.Context, id uint) (Core, error)
	UpdateReview(ctx context.Context, data Core) (Core, error)
}

type Data interface {
	CreateReview(ctx context.Context, review Core) (Core, error)
	GetAllReview(ctx context.Context, unitType uint) ([]Core, error)
	//GetBuildingReview(ctx context.Context, unitType uint) ([]Core, error)
	GetOneReview(ctx context.Context, id uint) (Core, error)
	UpdateReview(ctx context.Context, data Core) (Core, error)
}
