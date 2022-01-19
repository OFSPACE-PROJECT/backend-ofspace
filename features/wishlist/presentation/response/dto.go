package response

import (
	"ofspace-be/features/wishlist"
)

type Wishlist struct {
	Id         uint `json:"id"`
	UserId     uint `json:"user_id"`
	BuildingId uint `json:"building_id"`
	UnitId     uint `json:"unit_id"`
}

func ToWishlistResponse(c wishlist.Core) Wishlist {
	return Wishlist{
		Id:         c.Id,
		UserId:     c.UserId,
		BuildingId: c.BuildingId,
		UnitId:     c.UnitId,
	}
}
func FromWishlistCore(b wishlist.Core) Wishlist {
	return Wishlist{
		Id:         b.Id,
		UserId:     b.UserId,
		BuildingId: b.BuildingId,
		UnitId:     b.UnitId,
	}
}
func ToListWishlistCore(core []wishlist.Core) (response []Wishlist) {
	for _, wish := range core {
		response = append(response, FromWishlistCore(wish))
	}
	return
}
