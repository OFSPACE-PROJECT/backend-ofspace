package data

import (
	"ofspace-be/features/wishlist"
	"time"
)

type Wishlist struct {
	Id         uint `gorm:"primaryKey"`
	UserId     uint
	BuildingId uint
	UnitId     uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func fromWishlistCore(c wishlist.Core) Wishlist {
	return Wishlist{
		Id:         c.Id,
		UserId:     c.UserId,
		BuildingId: c.BuildingId,
		UnitId:     c.UnitId,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
	}
}

func toWishlistCore(c Wishlist) wishlist.Core {
	return wishlist.Core{
		Id:         c.Id,
		UserId:     c.UserId,
		BuildingId: c.BuildingId,
		UnitId:     c.UnitId,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
	}
}

func toSliceWishlistCore(wishlists []Wishlist) (result []wishlist.Core) {
	for _, wishlist1 := range wishlists {
		result = append(result, toWishlistCore(wishlist1))
	}
	return
}
