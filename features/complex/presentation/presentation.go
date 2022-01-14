package presentation

import (
	"net/http"
	"ofspace-be/features/complex"
	"ofspace-be/features/complex/presentation/request"
	"ofspace-be/features/complex/presentation/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ComplexPresentation struct {
	complexBusiness complex.Business
}

// type json map[string]interface{}

func NewComplexPresentation(cb complex.Business) *ComplexPresentation {
	return &ComplexPresentation{complexBusiness: cb}
}

func (cp *ComplexPresentation) CreateComplex(c echo.Context) error {
	CreateComplex := request.CreateComplex{}
	err := c.Bind(&CreateComplex)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	data, err2 := cp.complexBusiness.CreateComplex(ctx, CreateComplex.ToCore())
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToComplexResponse(data),
	})
}

func (cp *ComplexPresentation) UpdateComplex(c echo.Context) error {

	cUpdate := request.UpdateComplex{}
	err := c.Bind(&cUpdate)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	data, err2 := cp.complexBusiness.UpdateComplex(ctx, cUpdate.ToUpdateCore())
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToComplexResponse(data),
	})
}

func (cp *ComplexPresentation) GetComplex(c echo.Context) error {

	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	ctx := c.Request().Context()
	comp, err2 := cp.complexBusiness.GetComplex(ctx, uint(Id))
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToComplexResponse(comp),
	})
}

func (cp *ComplexPresentation) SearchComplex(c echo.Context) error {
	name := c.QueryParam("name")
	ctx := c.Request().Context()
	comp, err := cp.complexBusiness.SearchComplex(ctx, name)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToListCompCore(comp),
	})
}

func (cp *ComplexPresentation) RequestComplex(c echo.Context) error {
	Id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}
	name := c.QueryParam("name")
	ctx := c.Request().Context()
	data, err2 := cp.complexBusiness.RequestComplex(ctx, uint(Id), name)
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToComplexResponse(data),
	})
}
