package business

import (
	"context"
	"ofspace-be/features/wishlist"
	"time"
)

type wishlistBusiness struct {
	wishlistData   wishlist.Data
	contextTimeout time.Duration
}

func NewWishlistBusiness(wishlistData wishlist.Data, timeout time.Duration) wishlist.Business {
	return &wishlistBusiness{wishlistData: wishlistData, contextTimeout: timeout}
}

func (wb *wishlistBusiness) CreateWishlist(c context.Context, data wishlist.Core) (wishlist.Core, error) {
	ctx, error1 := context.WithTimeout(c, wb.contextTimeout)
	defer error1()
	data.UpdatedAt = time.Now()
	thisWishlist, err := wb.wishlistData.CreateWishlist(ctx, data)
	if err != nil {
		return wishlist.Core{}, err
	}
	return thisWishlist, nil
}
func (wb *wishlistBusiness) GetWishlist(c context.Context, wishlistId uint) (wishlist.Core, error) {
	ctx, error1 := context.WithTimeout(c, wb.contextTimeout)
	defer error1()
	thisWishlist, err := wb.wishlistData.GetWishlist(ctx, wishlistId)
	if err != nil {
		return wishlist.Core{}, err
	}
	return thisWishlist, nil
}
func (wb *wishlistBusiness) GetAllWishlists(c context.Context, UserId uint) ([]wishlist.Core, error) {
	ctx, error1 := context.WithTimeout(c, wb.contextTimeout)
	defer error1()
	data, err := wb.wishlistData.GetAllWishlists(ctx, UserId)
	if err != nil {
		return []wishlist.Core{}, err
	}
	return data, nil
}
func (wb *wishlistBusiness) DeleteWishlist(c context.Context, wishlistId uint) (wishlist.Core, error) {
	ctx, error1 := context.WithTimeout(c, wb.contextTimeout)
	defer error1()
	del, err2 := wb.wishlistData.DeleteWishlist(ctx, wishlistId)
	if err2 != nil {
		return wishlist.Core{}, err2
	}

	return del, nil
}
