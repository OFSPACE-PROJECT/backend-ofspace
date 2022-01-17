package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"ofspace-be/factory"
	mid "ofspace-be/middleware"
)

func New() *echo.Echo {
	presenter := factory.New()
	e := echo.New()
	mid.Logger(e)
	iJWT := e.Group("")
	iJWT.Use(middleware.JWT([]byte("123")))
	// user
	e.POST("/register", presenter.UserPresentation.RegisterUser)
	e.POST("/login", presenter.UserPresentation.LoginUser)
	e.GET("/user/:id", presenter.UserPresentation.GetUserByID)
	e.PUT("/user", presenter.UserPresentation.UpdateUser)
	e.DELETE("/user/:id", presenter.UserPresentation.DeleteUser)

	// complex
	iJWT.POST("/complex", presenter.ComplexPresentation.CreateComplex)
	e.GET("/complex/:id", presenter.ComplexPresentation.GetComplex)
	iJWT.POST("/complex/:id", presenter.ComplexPresentation.RequestComplex)
	e.GET("/complex/search", presenter.ComplexPresentation.SearchComplex)
	iJWT.PUT("/complex", presenter.ComplexPresentation.UpdateComplex)

	// accessibility
	iJWT.POST("/accessibility", presenter.AccessibilityPresentation.CreateAccessibility)
	e.GET("/accessibility/:id", presenter.AccessibilityPresentation.GetAccessibility)
	iJWT.POST("/accessibility/:id", presenter.AccessibilityPresentation.RequestAccessibility)
	e.GET("/accessibility/search", presenter.AccessibilityPresentation.SearchAccessibility)
	iJWT.PUT("/accessibility", presenter.AccessibilityPresentation.UpdateAccessibility)

	// facility
	iJWT.POST("/facility", presenter.FacilityPresentation.CreateFacility)
	e.GET("/facility/:id", presenter.FacilityPresentation.GetFacility)
	e.GET("/facility/search", presenter.FacilityPresentation.SearchFacility)
	iJWT.PUT("/facility", presenter.FacilityPresentation.UpdateFacility)
	iJWT.DELETE("/facility/:id", presenter.FacilityPresentation.DeleteFacility)

	//building
	e.POST("/building", presenter.BuildingPresentation.CreateBuilding)
	e.GET("/building/:id", presenter.BuildingPresentation.GetBuildingById)
	iJWT.GET("/building", presenter.BuildingPresentation.GetAllBuilding)
	e.GET("/building/status", presenter.BuildingPresentation.GetAllVerifiedBuilding)
	e.GET("/building/search", presenter.BuildingPresentation.SearchBuildingByName)
	iJWT.PUT("/building", presenter.BuildingPresentation.UpdateBuilding)
	//building exterior photo
	iJWT.POST("/building/exterior", presenter.BuildingPresentation.CreateExteriorPhoto)
	e.GET("/building/:id/exterior/:photo_id", presenter.BuildingPresentation.GetExteriorPhoto)
	e.GET("/building/:id/exterior", presenter.BuildingPresentation.GetAllExteriorPhoto)
	iJWT.PUT("/building/:id/exterior/:photo_id", presenter.BuildingPresentation.UpdateExteriorPhoto)
	iJWT.DELETE("/building/:id/exterior/:photo_id", presenter.BuildingPresentation.DeleteExteriorPhoto)
	//building floor photo
	iJWT.POST("/building/floor", presenter.BuildingPresentation.CreateFloorPhoto)
	e.GET("/building/:building_id/floor/:photo_id", presenter.BuildingPresentation.GetFloorPhoto)
	e.GET("/building/:building_id/floor", presenter.BuildingPresentation.GetAllFloorPhoto)
	iJWT.PUT("/building/:building_id/floor/:photo_id", presenter.BuildingPresentation.UpdateFloorPhoto)
	iJWT.DELETE("/building/:building_id/floor/:photo_id", presenter.BuildingPresentation.DeleteFloorPhoto)
	//building facilities
	iJWT.POST("/building/facility", presenter.BuildingPresentation.AddFacilityToBuilding)
	e.GET("/building/:building_id/facility/:facility_id", presenter.BuildingPresentation.GetBuildingFacility)
	e.GET("/building/:building_id/facility", presenter.BuildingPresentation.GetAllBuildingFacility)
	iJWT.DELETE("/building/:building_id/facility/:photo_id", presenter.BuildingPresentation.DeleteFacility)

	//unit
	e.POST("/unit", presenter.UnitPresentation.CreateUnit)
	e.GET("/unit/:id", presenter.UnitPresentation.GetUnitById)
	iJWT.GET("/unit", presenter.UnitPresentation.GetAllUnit)
	e.GET("/unit/type", presenter.UnitPresentation.GetUnitByType)
	iJWT.PUT("/unit", presenter.UnitPresentation.UpdateUnit)
	//unit interior photo
	iJWT.POST("/unit/interior", presenter.UnitPresentation.CreateInteriorPhoto)
	e.GET("/unit/:id/interior/:photo_id", presenter.UnitPresentation.GetInteriorPhoto)
	e.GET("/unit/:id/interior", presenter.UnitPresentation.GetAllInteriorPhoto)
	iJWT.PUT("/unit/:id/interior/:photo_id", presenter.UnitPresentation.UpdateInteriorPhoto)
	iJWT.DELETE("/unit/:id/interior/:photo_id", presenter.UnitPresentation.DeleteInteriorPhoto)
	//unit facilities
	iJWT.POST("/unit/facility", presenter.UnitPresentation.AddFacilityToUnit)
	e.GET("/unit/:unit_id/facility/:facility_id", presenter.UnitPresentation.GetUnitFacility)
	e.GET("/unit/:unit_id/facility", presenter.UnitPresentation.GetAllUnitFacility)
	iJWT.DELETE("/unit/:unit_id/facility/:photo_id", presenter.UnitPresentation.DeleteUnitFacility)

	return e
}
