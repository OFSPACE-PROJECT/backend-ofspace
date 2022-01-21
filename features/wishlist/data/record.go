package data

import (
	"ofspace-be/features/wishlist"
	"time"
)

type Wishlist struct {
	Id         uint `gorm:"primaryKey"`
	UserId     uint
	User       User `gorm:"foreignKey:UserId"`
	BuildingId uint
	UnitId     uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func fromUserCore(c wishlist.User) User {
	return User{
		ID:   c.ID,
		Name: c.Name,
	}
}

func toUserCore(c User) wishlist.User {
	return wishlist.User{
		ID:   c.ID,
		Name: c.Name,
	}
}

func fromWishlistCore(c wishlist.Core) Wishlist {
	return Wishlist{
		Id:         c.Id,
		UserId:     c.UserId,
		BuildingId: c.BuildingId,
		User:       fromUserCore(c.User),
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
		User:       toUserCore(c.User),
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
