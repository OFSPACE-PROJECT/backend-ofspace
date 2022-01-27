package complex

import (
	"context"
	"ofspace-be/features/building"
	"time"
)

type Core struct {
	Id        uint
	Name      string
	Address   string
	Latitude  string
	Longitude string
	Buildings []building.Core
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	CreateComplex(ctx context.Context, complex Core) (Core, error)
	GetComplex(ctx context.Context, id uint) (Core, error)
	GetAllComplex(ctx context.Context) ([]Core, error)
	SearchComplex(ctx context.Context, name string) ([]Core, error)
	SearchComplexByAddress(ctx context.Context, address string) ([]Core, error)
	UpdateComplex(ctx context.Context, complex Core) (Core, error)
	//RequestComplex(ctx context.Context, id uint, name string) (Core, error)
}

type Data interface {
	CreateComplex(ctx context.Context, complex Core) (Core, error)
	GetComplex(ctx context.Context, id uint) (Core, error)
	GetAllComplex(ctx context.Context) ([]Core, error)
	SearchComplex(ctx context.Context, name string) ([]Core, error)
	SearchComplexByAddress(ctx context.Context, address string) ([]Core, error)
	UpdateComplex(ctx context.Context, complex Core) (Core, error)
	//RequestComplex(ctx context.Context, id uint, name string) (Core, error)
}
