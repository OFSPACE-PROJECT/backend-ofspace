package wishlist

import (
	"context"
	"time"
)

type Core struct {
	Id         uint
	UserId     uint
	BuildingId uint
	UnitId     uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
type Business interface {
	CreateWishlist(ctx context.Context, wishlist Core) (Core, error)
	GetWishlist(ctx context.Context, wishlistId uint) (Core, error)
	GetAllWishlists(ctx context.Context, UserId uint) ([]Core, error)
	DeleteWishlist(ctx context.Context, wishlistId uint) (Core, error)
}

type Data interface {
	CreateWishlist(ctx context.Context, wishlist Core) (Core, error)
	GetWishlist(ctx context.Context, wishlistId uint) (Core, error)
	GetAllWishlists(ctx context.Context, UserId uint) ([]Core, error)
	DeleteWishlist(ctx context.Context, wishlistId uint) (Core, error)
}
