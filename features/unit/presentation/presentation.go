package presentation

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ofspace-be/features/unit"
	"ofspace-be/features/unit/presentation/request"
	"ofspace-be/features/unit/presentation/response"
	"strconv"
)

type UnitPresentation struct {
	unitBusiness unit.Business
}

func NewUnitPresentation(bb unit.Business) *UnitPresentation {
	return &UnitPresentation{unitBusiness: bb}
}

func (bp *UnitPresentation) CreateUnit(c echo.Context) error {
	CreateUnit := request.CreateUnit{}
	err := c.Bind(&CreateUnit)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err2 := bp.unitBusiness.CreateUnit(ctx, CreateUnit.ToCore())
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToUnitResponse(data),
	})

}
func (bp *UnitPresentation) GetAllUnit(c echo.Context) error {
	id := c.QueryParam("building_id")
	ids, err0 := strconv.Atoi(id)
	if err0 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err0.Error(),
		})
	}
	ctx := c.Request().Context()
	fac, err := bp.unitBusiness.GetAllUnit(ctx, uint(ids))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToListUnitCore(fac),
	})
}

//func (bp *UnitPresentation) GetUnverifiedUnitById(c echo.Context) error {
//	Id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		return c.JSON(http.StatusForbidden, map[string]interface{}{
//			"message": err.Error(),
//		})
//	}
//
//	status := "unverified"
//	ctx := c.Request().Context()
//	build, err2 := bp.unitBusiness.GetUnverifiedUnitById(ctx, uint(Id), status)
//	if err2 != nil {
//		return c.JSON(http.StatusForbidden, map[string]interface{}{
//			"message": err2.Error(),
//		})
//	}
//
//	return c.JSON(http.StatusOK, map[string]interface{}{
//		"message": "Success",
//		"data":    response.ToUnitResponse(build),
//	})
//}

func (bp *UnitPresentation) GetUnitById(c echo.Context) error {
	Id, err := strconv.Atoi(c.Param("id"))
	buildId, err2 := strconv.Atoi(c.QueryParam("building_id"))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	//status := "verified"
	ctx := c.Request().Context()
	build, err2 := bp.unitBusiness.GetUnitById(ctx, uint(Id), uint(buildId))
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToUnitResponse(build),
	})
}
func (bp *UnitPresentation) GetUnitByType(c echo.Context) error {
	Id := c.QueryParam("unit_type")
	buildId, err1 := strconv.Atoi(c.Param("id"))
	if err1 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err1.Error(),
		})
	}
	//status := "verified"
	ctx := c.Request().Context()
	build, err2 := bp.unitBusiness.GetUnitByType(ctx, uint(buildId), Id)
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToUnitResponse(build),
	})
}
func (bp *UnitPresentation) UpdateUnit(c echo.Context) error {
	cUpdate := request.UpdateUnit{}
	err := c.Bind(&cUpdate)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	data, err2 := bp.unitBusiness.UpdateUnit(ctx, cUpdate.ToUpdateCore())
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToUnitResponse(data),
	})
}
func (bp *UnitPresentation) CreateInteriorPhoto(c echo.Context) error {
	CreateExt := request.CreateInteriorPhoto{}
	//Id, err0 := strconv.Atoi(c.Param("id"))
	//if err0 != nil {
	//	return c.JSON(http.StatusForbidden, map[string]interface{}{
	//		"message": err0.Error(),
	//	})
	//}
	err := c.Bind(&CreateExt)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	//data, err2 := bp.unitBusiness.CreateInteriorPhoto(ctx, uint(Id), CreateExt.ToInteriorCore())
	data, err2 := bp.unitBusiness.CreateInteriorPhoto(ctx, CreateExt.ToInteriorCore())
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToInteriorResponse(data),
	})

}
func (bp *UnitPresentation) UpdateInteriorPhoto(c echo.Context) error {
	cUpdate := request.UpdateInteriorPhoto{}
	err := c.Bind(&cUpdate)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	data, err2 := bp.unitBusiness.UpdateInteriorPhoto(ctx, cUpdate.ToUpdateInteriorCore())
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToInteriorResponse(data),
	})
}
func (bp *UnitPresentation) GetInteriorPhoto(c echo.Context) error {
	id := c.Param("id")
	photoId := c.Param("photo_id")
	ids, err0 := strconv.Atoi(id)
	if err0 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err0.Error(),
		})
	}
	photoIds, err3 := strconv.Atoi(photoId)
	if err3 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err3.Error(),
		})
	}
	ctx := c.Request().Context()
	fac, err := bp.unitBusiness.GetInteriorPhoto(ctx, uint(ids), uint(photoIds))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToInteriorResponse(fac),
	})
}
func (bp *UnitPresentation) GetAllInteriorPhoto(c echo.Context) error {
	id := c.Param("id")
	ids, err0 := strconv.Atoi(id)
	if err0 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err0.Error(),
		})
	}
	ctx := c.Request().Context()
	fac, err := bp.unitBusiness.GetAllInteriorPhoto(ctx, uint(ids))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToListInteriorCore(fac),
	})
}
func (bp *UnitPresentation) DeleteInteriorPhoto(c echo.Context) error {
	id := c.Param("id")
	photoId := c.Param("photo_id")
	ids, err0 := strconv.Atoi(id)
	if err0 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err0.Error(),
		})
	}
	photoIds, err3 := strconv.Atoi(photoId)
	if err3 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err3.Error(),
		})
	}
	ctx := c.Request().Context()
	_, err := bp.unitBusiness.DeleteInteriorPhoto(ctx, uint(ids), uint(photoIds))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func (bp *UnitPresentation) AddFacilityToUnit(c echo.Context) error {
	ctx := c.Request().Context()
	fId, err2 := strconv.Atoi(c.QueryParam("facility_id"))
	bId, err5 := strconv.Atoi(c.QueryParam("unit_id"))
	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err2.Error(),
		})
	}
	if err5 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err5.Error(),
		})
	}
	data, err := bp.unitBusiness.AddFacilityToUnit(ctx, uint(fId), uint(bId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"detail":  data,
	})
}

func (bp *UnitPresentation) GetAllUnitFacility(c echo.Context) error {
	id := c.Param("id")
	ids, err0 := strconv.Atoi(id)
	if err0 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err0.Error(),
		})
	}
	ctx := c.Request().Context()
	fac, err := bp.unitBusiness.GetAllUnitFacility(ctx, uint(ids))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.FromUnitCore(fac),
	})
}

func (bp *UnitPresentation) GetUnitFacility(c echo.Context) error {
	id := c.Param("id")
	pId := c.Param("facility_id")
	ids, err0 := strconv.Atoi(id)
	if err0 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err0.Error(),
		})
	}
	fIds, err3 := strconv.Atoi(pId)
	if err3 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err3.Error(),
		})
	}
	ctx := c.Request().Context()
	fac, err := bp.unitBusiness.GetUnitFacility(ctx, uint(ids), uint(fIds))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToUnitFacilityResponse(fac),
	})
}

func (bp *UnitPresentation) DeleteUnitFacility(c echo.Context) error {
	id := c.Param("id")
	pId := c.Param("facility_id")
	ids, err0 := strconv.Atoi(id)
	if err0 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err0.Error(),
		})
	}
	pIds, err3 := strconv.Atoi(pId)
	if err3 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err3.Error(),
		})
	}
	ctx := c.Request().Context()
	_, err := bp.unitBusiness.DeleteUnitFacility(ctx, uint(ids), uint(pIds))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}
