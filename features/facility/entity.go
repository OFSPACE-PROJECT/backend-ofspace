package facility

import (
	"context"
	"time"
)

type Facility struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	CreateFacility(ctx context.Context, core Facility) (Facility, error)
	GetFacility(ctx context.Context, id uint) (Facility, error)
	SearchFacility(ctx context.Context, name string) ([]Facility, error)
	UpdateFacility(ctx context.Context, core Facility) (Facility, error)
	DeleteFacility(ctx context.Context, id uint) (Facility, error)
}

type Data interface {
	CreateFacility(ctx context.Context, core Facility) (Facility, error)
	GetFacility(ctx context.Context, id uint) (Facility, error)
	UpdateFacility(ctx context.Context, core Facility) (Facility, error)
	DeleteFacility(ctx context.Context, id uint) (Facility, error)
	SearchFacility(ctx context.Context, name string) ([]Facility, error)
}
