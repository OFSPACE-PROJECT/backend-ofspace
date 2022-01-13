package accessibility

import (
	"context"
	"time"
)

type Core struct {
	Id        uint
	Name      string
	Latitude  string
	Longitude string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	CreateAccessibility(ctx context.Context, accessibility Core) (Core, error)
	GetAccessibility(ctx context.Context, id uint) (Core, error)
	SearchAccessibility(ctx context.Context, name string) ([]Core, error)
	UpdateAccessibility(ctx context.Context, accessibility Core) (Core, error)
	RequestAccessibility(ctx context.Context, id uint, name string) (Core, error)
}

type Data interface {
	CreateAccessibility(ctx context.Context, accessibility Core) (Core, error)
	GetAccessibility(ctx context.Context, id uint) (Core, error)
	SearchAccessibility(ctx context.Context, name string) ([]Core, error)
	UpdateAccessibility(ctx context.Context, accessibility Core) (Core, error)
	RequestAccessibility(ctx context.Context, id uint, name string) (Core, error)
}
