package presentation

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ofspace-be/features/building"
	"ofspace-be/features/building/presentation/request"
	"ofspace-be/features/building/presentation/response"
	"strconv"
)

type BuildingPresentation struct {
	buildingBusiness building.Business
}

func NewBuildingPresentation(bb building.Business) *BuildingPresentation {
	return &BuildingPresentation{buildingBusiness: bb}
}

func (bp *BuildingPresentation) CreateBuilding(c echo.Context) error {
	CreateBuilding := request.CreateBuilding{}
	err := c.Bind(&CreateBuilding)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err2 := bp.buildingBusiness.CreateBuilding(ctx, CreateBuilding.ToCore())
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToBuildingResponse(data),
	})

}
func (bp *BuildingPresentation) GetAllBuilding(c echo.Context) error {
	id := c.QueryParam("complex_id")
	ids, err0 := strconv.Atoi(id)
	if err0 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err0.Error(),
		})
	}
	ctx := c.Request().Context()
	fac, err := bp.buildingBusiness.GetAllBuilding(ctx, uint(ids))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToListBuildingCore(fac),
	})
}
func (bp *BuildingPresentation) GetAllVerifiedBuilding(c echo.Context) error {
	id := c.QueryParam("complex_id")
	ids, err0 := strconv.Atoi(id)
	if err0 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err0.Error(),
		})
	}
	status := "verified"
	ctx := c.Request().Context()
	fac, err := bp.buildingBusiness.GetAllVerifiedBuilding(ctx, uint(ids), status)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToListBuildingCore(fac),
	})
}
func (bp *BuildingPresentation) SearchBuildingByName(c echo.Context) error {
	name := c.QueryParam("name")
	status := "verified"
	ctx := c.Request().Context()
	fac, err := bp.buildingBusiness.SearchBuildingByName(ctx, name, status)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToListBuildingCore(fac),
	})
}

//func (bp *BuildingPresentation) GetUnverifiedBuildingById(c echo.Context) error {
//	Id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		return c.JSON(http.StatusForbidden, map[string]interface{}{
//			"message": err.Error(),
//		})
//	}
//
//	status := "unverified"
//	ctx := c.Request().Context()
//	build, err2 := bp.buildingBusiness.GetUnverifiedBuildingById(ctx, uint(Id), status)
//	if err2 != nil {
//		return c.JSON(http.StatusForbidden, map[string]interface{}{
//			"message": err2.Error(),
//		})
//	}
//
//	return c.JSON(http.StatusOK, map[string]interface{}{
//		"message": "Success",
//		"data":    response.ToBuildingResponse(build),
//	})
//}

func (bp *BuildingPresentation) GetBuildingById(c echo.Context) error {
	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	//status := "verified"
	ctx := c.Request().Context()
	build, err2 := bp.buildingBusiness.GetBuildingById(ctx, uint(Id))
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToBuildingResponse(build),
	})
}
func (bp *BuildingPresentation) UpdateBuilding(c echo.Context) error {
	cUpdate := request.UpdateBuilding{}
	err := c.Bind(&cUpdate)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	data, err2 := bp.buildingBusiness.UpdateBuilding(ctx, cUpdate.ToUpdateCore())
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToBuildingResponse(data),
	})
}
func (bp *BuildingPresentation) CreateExteriorPhoto(c echo.Context) error {
	CreateExt := request.CreateExteriorPhoto{}
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
	//data, err2 := bp.buildingBusiness.CreateExteriorPhoto(ctx, uint(Id), CreateExt.ToExteriorCore())
	data, err2 := bp.buildingBusiness.CreateExteriorPhoto(ctx, CreateExt.ToExteriorCore())
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToExteriorResponse(data),
	})

}
func (bp *BuildingPresentation) UpdateExteriorPhoto(c echo.Context) error {
	cUpdate := request.UpdateExteriorPhoto{}
	err := c.Bind(&cUpdate)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	data, err2 := bp.buildingBusiness.UpdateExteriorPhoto(ctx, cUpdate.ToUpdateExteriorCore())
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToExteriorResponse(data),
	})
}
func (bp *BuildingPresentation) GetExteriorPhoto(c echo.Context) error {
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
	fac, err := bp.buildingBusiness.GetExteriorPhoto(ctx, uint(ids), uint(photoIds))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToExteriorResponse(fac),
	})
}
func (bp *BuildingPresentation) GetAllExteriorPhoto(c echo.Context) error {
	id := c.Param("id")
	ids, err0 := strconv.Atoi(id)
	if err0 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err0.Error(),
		})
	}
	ctx := c.Request().Context()
	fac, err := bp.buildingBusiness.GetAllExteriorPhoto(ctx, uint(ids))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToListExteriorCore(fac),
	})
}
func (bp *BuildingPresentation) DeleteExteriorPhoto(c echo.Context) error {
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
	_, err := bp.buildingBusiness.DeleteExteriorPhoto(ctx, uint(ids), uint(photoIds))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func (bp *BuildingPresentation) CreateFloorPhoto(c echo.Context) error {
	CreateExt := request.CreateFloorPhoto{}
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
	//data, err2 := bp.buildingBusiness.CreateExteriorPhoto(ctx, uint(Id), CreateExt.ToExteriorCore())
	data, err2 := bp.buildingBusiness.CreateFloorPhoto(ctx, CreateExt.ToFloorCore())
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToFloorResponse(data),
	})

}
func (bp *BuildingPresentation) UpdateFloorPhoto(c echo.Context) error {
	cUpdate := request.UpdateFloorPhoto{}
	err := c.Bind(&cUpdate)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	data, err2 := bp.buildingBusiness.UpdateFloorPhoto(ctx, cUpdate.ToUpdateFloorCore())
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToFloorResponse(data),
	})
}
func (bp *BuildingPresentation) GetFloorPhoto(c echo.Context) error {
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
	fac, err := bp.buildingBusiness.GetFloorPhoto(ctx, uint(ids), uint(photoIds))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToFloorResponse(fac),
	})
}
func (bp *BuildingPresentation) GetAllFloorPhoto(c echo.Context) error {
	id := c.Param("id")
	ids, err0 := strconv.Atoi(id)
	if err0 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err0.Error(),
		})
	}
	ctx := c.Request().Context()
	fac, err := bp.buildingBusiness.GetAllFloorPhoto(ctx, uint(ids))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToListFloorCore(fac),
	})
}
func (bp *BuildingPresentation) DeleteFloorPhoto(c echo.Context) error {
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
	_, err := bp.buildingBusiness.DeleteFloorPhoto(ctx, uint(ids), uint(photoIds))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func (bp *BuildingPresentation) AddFacilityToBuilding(c echo.Context) error {

	//buildingId, err := strconv.Atoi(c.QueryParam("id"))
	//if err != nil {
	//	return c.JSON(http.StatusForbidden, map[string]interface{}{
	//		"message": err.Error(),
	//	})
	//}
	//facilityId, err2 := strconv.Atoi(c.QueryParam("facility_id"))
	//if err2 != nil {
	//	return c.JSON(http.StatusForbidden, map[string]interface{}{
	//		"message": err2.Error(),
	//	})
	//}
	//
	ctx := c.Request().Context()
	//data, err3 := bp.buildingBusiness.AddFacilityToBuilding(ctx, uint(buildingId), uint(facilityId))
	//if err3 != nil {
	//	return c.JSON(http.StatusForbidden, map[string]interface{}{
	//		"message": err3.Error(),
	//	})
	//}
	//
	//return c.JSON(http.StatusOK, map[string]interface{}{
	//	"message": "Success",
	//	"data":    response.FromBuildingFacilityCore(data),
	//})
	buildF := request.AddFacility{}
	facId := buildF.FacilityId
	buildId := buildF.BuildingId
	err2 := c.Bind(&buildF)
	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err2.Error(),
		})
	}
	// fmt.Println("detail presentation ========== ", detail)
	//data, err := bp.buildingBusiness.AddFacilityToBuilding(request.T(detail))
	data, err := bp.buildingBusiness.AddFacilityToBuilding(ctx, facId, buildId)
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

func (bp *BuildingPresentation) GetAllBuildingFacility(c echo.Context) error {
	id := c.Param("id")
	ids, err0 := strconv.Atoi(id)
	if err0 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err0.Error(),
		})
	}
	ctx := c.Request().Context()
	fac, err := bp.buildingBusiness.GetAllBuildingFacility(ctx, uint(ids))
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

func (bp *BuildingPresentation) GetBuildingFacility(c echo.Context) error {
	id := c.Param("id")
	photoId := c.Param("facility_id")
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
	fac, err := bp.buildingBusiness.GetBuildingFacility(ctx, uint(ids), uint(photoIds))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToFacilityResponse(fac),
	})
}

func (bp *BuildingPresentation) DeleteFacility(c echo.Context) error {
	id := c.Param("id")
	photoId := c.Param("facility_id")
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
	_, err := bp.buildingBusiness.DeleteFacility(ctx, uint(ids), uint(photoIds))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}
