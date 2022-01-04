package facility

import (
	"context"
	"time"
)

type Core struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	CreateFacility(ctx context.Context, name string) (Core, error)
	GetFacility(ctx context.Context, name string) (Core, error)
	UpdateFacility(ctx context.Context, name string) (Core, error)
	DeleteFacility(ctx context.Context, id uint) error
}

type Data interface {
	CreateFacility(ctx context.Context, name string) (Core, error)
	GetFacility(ctx context.Context, name string) (Core, error)
	UpdateFacility(ctx context.Context, name string) (Core, error)
	DeleteFacility(ctx context.Context, id uint) error
}
