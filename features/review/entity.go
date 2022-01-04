package review

import (
	"context"
	"time"
)

type Core struct {
	Id                    uint
	CostumerId            uint
	BuildingId            uint
	UnitType              string
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

type Business interface {
	CreateReview(ctx context.Context, review Core) (Core, error)
	GetAllReview(ctx context.Context, buildingId uint, unitType string) ([]Core, error)
	GetOneReview(ctx context.Context, buildingId uint, unitType string, id uint) (Core, error)
	UpdateReview(ctx context.Context, id uint, review Core) (Core, error)
}

type Data interface {
	CreateReview(ctx context.Context, review Core) (Core, error)
	GetAllReview(ctx context.Context, buildingId uint, unitType string) ([]Core, error)
	GetOneReview(ctx context.Context, buildingId uint, unitType string, id uint) (Core, error)
	UpdateReview(ctx context.Context, id uint, review Core) (Core, error)
}
