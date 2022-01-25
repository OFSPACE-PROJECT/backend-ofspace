package presentation

import (
	"net/http"
	"ofspace-be/features/accessibility"
	"ofspace-be/features/accessibility/presentation/request"
	"ofspace-be/features/accessibility/presentation/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AccessibilityPresentation struct {
	accessibilityBusiness accessibility.Business
}

func NewAccessibilityPresentation(cb accessibility.Business) *AccessibilityPresentation {
	return &AccessibilityPresentation{accessibilityBusiness: cb}
}

func (cp *AccessibilityPresentation) CreateAccessibility(c echo.Context) error {
	CreateAccessibility := request.CreateAccessibility{}
	err := c.Bind(&CreateAccessibility)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	data, err2 := cp.accessibilityBusiness.CreateAccessibility(ctx, CreateAccessibility.ToCore())
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToAccessibilityResponse(data),
	})
}

func (cp *AccessibilityPresentation) UpdateAccessibility(c echo.Context) error {

	cUpdate := request.UpdateAccessibility{}
	err := c.Bind(&cUpdate)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	data, err2 := cp.accessibilityBusiness.UpdateAccessibility(ctx, cUpdate.ToUpdateCore())
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToAccessibilityResponse(data),
	})
}

func (cp *AccessibilityPresentation) GetAccessibility(c echo.Context) error {

	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	ctx := c.Request().Context()
	comp, err2 := cp.accessibilityBusiness.GetAccessibility(ctx, uint(Id))
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToAccessibilityResponse(comp),
	})
}

func (cp *AccessibilityPresentation) SearchAccessibility(c echo.Context) error {
	name := c.QueryParam("name")
	ctx := c.Request().Context()
	acc, err := cp.accessibilityBusiness.SearchAccessibility(ctx, name)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToListAccCore(acc),
	})
}

func (cp *AccessibilityPresentation) SearchAccessibilityByAddress(c echo.Context) error {
	name := c.QueryParam("address")
	ctx := c.Request().Context()
	acc, err := cp.accessibilityBusiness.SearchAccessibility(ctx, name)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToListAccCore(acc),
	})
}

func (cp *AccessibilityPresentation) RequestAccessibility(c echo.Context) error {
	Id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}
	name := c.QueryParam("name")
	ctx := c.Request().Context()
	data, err2 := cp.accessibilityBusiness.RequestAccessibility(ctx, uint(Id), name)
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToAccessibilityResponse(data),
	})
}
