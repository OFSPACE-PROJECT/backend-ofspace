package presentation

import (
	"net/http"
	"ofspace-be/features/facility"
	"ofspace-be/features/facility/presentation/request"
	"ofspace-be/features/facility/presentation/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FacilityPresentation struct {
	facilityBusiness facility.Business
}

func NewFacilityPresentation(cb facility.Business) *FacilityPresentation {
	return &FacilityPresentation{facilityBusiness: cb}
}

func (fp *FacilityPresentation) CreateFacility(c echo.Context) error {
	CreateFacility := request.CreateFacility{}
	err := c.Bind(&CreateFacility)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	data, err2 := fp.facilityBusiness.CreateFacility(ctx, CreateFacility.ToCore())
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToFacilityResponse(data),
	})
}

func (fp *FacilityPresentation) UpdateFacility(c echo.Context) error {

	cUpdate := request.CreateFacility{}
	err := c.Bind(&cUpdate)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	data, err2 := fp.facilityBusiness.UpdateFacility(ctx, cUpdate.ToCore())
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToFacilityResponse(data),
	})
}

func (fp *FacilityPresentation) GetFacility(c echo.Context) error {

	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	ctx := c.Request().Context()
	comp, err2 := fp.facilityBusiness.GetFacility(ctx, uint(Id))
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToFacilityResponse(comp),
	})
}

func (fp *FacilityPresentation) SearchFacility(c echo.Context) error {
	name := c.QueryParam("name")
	ctx := c.Request().Context()
	fac, err := fp.facilityBusiness.SearchFacility(ctx, name)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToListFacilityCore(fac),
	})
}

func (fp *FacilityPresentation) DeleteFacility(c echo.Context) error {

	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}
	ctx := c.Request().Context()
	_, err = fp.facilityBusiness.DeleteFacility(ctx, uint(Id))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}
