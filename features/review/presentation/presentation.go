package presentation

import (
	"net/http"
	"ofspace-be/features/review"
	"ofspace-be/features/review/presentation/request"
	"ofspace-be/features/review/presentation/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReviewPresentation struct {
	ReviewBusiness review.Business
}

func NewReviewPresentation(up review.Business) *ReviewPresentation {
	return &ReviewPresentation{ReviewBusiness: up}
}

func (up *ReviewPresentation) CreateReview(c echo.Context) error {
	review := request.CreateReview{}
	c.Bind(&review)
	ctx := c.Request().Context()
	data, err := up.ReviewBusiness.CreateReview(ctx, review.ToCore())
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.FromReviewCore(data),
	})
}

func (up *ReviewPresentation) GetAllReview(c echo.Context) error {

	Id, err := strconv.Atoi(c.Param("unit"))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}
	ctx := c.Request().Context()
	data, err := up.ReviewBusiness.GetAllReview(ctx, uint(Id))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToListReview(data),
	})
}

func (up *ReviewPresentation) GetOneReview(c echo.Context) error {

	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}
	ctx := c.Request().Context()
	data, err := up.ReviewBusiness.GetOneReview(ctx, uint(Id))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.FromReviewCore(data),
	})
}

func (up *ReviewPresentation) UpdateReview(c echo.Context) error {

	reviewUpdate := request.ReviewUpdate{}
	c.Bind(&reviewUpdate)
	ctx := c.Request().Context()
	_, err := up.ReviewBusiness.UpdateReview(ctx, reviewUpdate.ToCoreUpdate())
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}
