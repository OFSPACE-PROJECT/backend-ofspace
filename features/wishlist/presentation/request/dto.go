package request

import (
	"ofspace-be/features/wishlist"
)

type CreateWishlist struct {
	UserId     uint `json:"user_id"`
	BuildingId uint `json:"building_id"`
	UnitId     uint `json:"unit_id"`
}

func (c *CreateWishlist) ToCore() wishlist.Core {
	return wishlist.Core{
		UserId:     c.UserId,
		BuildingId: c.BuildingId,
		UnitId:     c.UnitId,
	}
}
