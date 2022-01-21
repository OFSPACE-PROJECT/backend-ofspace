package data

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"ofspace-be/features/wishlist"
)

type wishlistData struct {
	Connect *gorm.DB
}

func NewWishlistData(connect *gorm.DB) wishlist.Data {
	return &wishlistData{Connect: connect}
}

func (wd *wishlistData) CreateWishlist(ctx context.Context, core wishlist.Core) (wishlist.Core, error) {
	wishlist1 := fromWishlistCore(core)
	result := wd.Connect.Create(&wishlist1)
	if result.Error != nil {
		return wishlist.Core{}, result.Error
	}
	return toWishlistCore(wishlist1), nil
}
func (wd *wishlistData) GetWishlist(ctx context.Context, wishlistId uint) (wishlist.Core, error) {
	var wishlist1 Wishlist
	result := wd.Connect.Preload("User").First(&wishlist1, "id= ?", wishlistId)
	if result.Error != nil {
		return wishlist.Core{}, result.Error
	}
	return toWishlistCore(wishlist1), nil
}
func (wd *wishlistData) GetAllWishlists(ctx context.Context, UserId uint) ([]wishlist.Core, error) {
	var wishlists []Wishlist
	result := wd.Connect.Debug().Find(&wishlists, "id", UserId)

	if result.Error != nil {
		fmt.Println(result.Error)
		return []wishlist.Core{}, result.Error
	}
	return toSliceWishlistCore(wishlists), nil
}
func (wd *wishlistData) DeleteWishlist(ctx context.Context, wishlistId uint) (wishlist.Core, error) {
	var fac Wishlist
	result := wd.Connect.Delete(&fac, "id= ?", wishlistId)

	if result.Error != nil {
		return wishlist.Core{}, result.Error
	}

	return toWishlistCore(fac), nil
}
