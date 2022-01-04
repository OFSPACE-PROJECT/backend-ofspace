package complex

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
	CreateComplex(ctx context.Context, complex Core) (Core, error)
	GetComplex(ctx context.Context, id uint) (Core, error)
	SearchComplex(ctx context.Context, name string) ([]Core, error)
	UpdateComplex(ctx context.Context, id uint) (Core, error)
	RequestComplex(ctx context.Context, id uint, name string) (Core, error)
}

type Data interface {
	CreateComplex(ctx context.Context, complex Core) (Core, error)
	GetComplex(ctx context.Context, id uint) (Core, error)
	SearchComplex(ctx context.Context, name string) ([]Core, error)
	UpdateComplex(ctx context.Context, id uint) (Core, error)
	RequestComplex(ctx context.Context, id uint, name string) (Core, error)
}
