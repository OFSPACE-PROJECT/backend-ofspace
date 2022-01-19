package presentation

import (
	"github.com/labstack/echo"
	"net/http"
	"ofspace-be/features/wishlist/presentation/request"
	"ofspace-be/features/wishlist/presentation/response"
	"strconv"

	"ofspace-be/features/wishlist"
)

type WishlistPresentation struct {
	wishlistBusiness wishlist.Business
}

func NewFacilityPresentation(wb wishlist.Business) *WishlistPresentation {
	return &WishlistPresentation{wishlistBusiness: wb}
}

func (wp *WishlistPresentation) CreateWishlist(c echo.Context) error {
	CreateWishlist := request.CreateWishlist{}
	err := c.Bind(&CreateWishlist)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	data, err2 := wp.wishlistBusiness.CreateWishlist(ctx, CreateWishlist.ToCore())
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToWishlistResponse(data),
	})
}
func (wp *WishlistPresentation) GetWishlist(c echo.Context) error {
	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	ctx := c.Request().Context()
	comp, err2 := wp.wishlistBusiness.GetWishlist(ctx, uint(Id))
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToWishlistResponse(comp),
	})
}
func (wp *WishlistPresentation) GetAllWishlists(c echo.Context) error {
	id := c.QueryParam("user_id")
	ids, err0 := strconv.Atoi(id)
	if err0 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err0.Error(),
		})
	}
	ctx := c.Request().Context()
	fac, err := wp.wishlistBusiness.GetAllWishlists(ctx, uint(ids))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToListWishlistCore(fac),
	})
}
func (wp *WishlistPresentation) DeleteWishlist(c echo.Context) error {
	id := c.Param("id")
	ids, err0 := strconv.Atoi(id)
	if err0 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err0.Error(),
		})
	}
	ctx := c.Request().Context()
	_, err := wp.wishlistBusiness.DeleteWishlist(ctx, uint(ids))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}
