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
	CreateFacility(ctx context.Context, core Core) (Core, error)
	GetFacility(ctx context.Context, id uint) (Core, error)
	SearchFacility(ctx context.Context, name string) ([]Core, error)
	UpdateFacility(ctx context.Context, core Core) (Core, error)
	DeleteFacility(ctx context.Context, id uint) (Core, error)
}

type Data interface {
	CreateFacility(ctx context.Context, core Core) (Core, error)
	GetFacility(ctx context.Context, id uint) (Core, error)
	UpdateFacility(ctx context.Context, core Core) (Core, error)
	DeleteFacility(ctx context.Context, id uint) (Core, error)
	SearchFacility(ctx context.Context, name string) ([]Core, error)
}
