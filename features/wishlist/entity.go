package wishlist

import (
	"context"
	"time"
)

type Core struct {
	Id         uint
	UserId     string
	BuildingId uint
	UnitId     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
type Business interface {
	CreateWishlist(ctx context.Context, wishlist Core) (Core, error)
	GetWishlist(ctx context.Context, userId uint, wishlistId uint) (Core, error)
	GetAllWishlists(ctx context.Context, UserId string) ([]Core, error)
	DeleteWishlist(ctx context.Context, wishlistId uint) error
}

type Data interface {
	CreateWishlist(ctx context.Context, wishlist Core) (Core, error)
	GetWishlist(ctx context.Context, userId uint, wishlistId uint) (Core, error)
	GetAllWishlists(ctx context.Context, UserId string) ([]Core, error)
	DeleteWishlist(ctx context.Context, wishlistId uint) error
}
