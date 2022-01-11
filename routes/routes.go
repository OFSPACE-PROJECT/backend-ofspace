package routes

import (
	"ofspace-be/factory"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	factory.New()
	presenter := factory.New()
	e := echo.New()

	// users
	e.POST("/facility", presenter.FacilityPresentation.CreateFacility)
	e.GET("/facility/:id", presenter.FacilityPresentation.GetFacility)
	e.PUT("/facility/:id", presenter.FacilityPresentation.UpdateFacility)
	e.DELETE("/facility/:id", presenter.FacilityPresentation.DeleteFacility)
	e.GET("/facility/search", presenter.FacilityPresentation.SearchFacility)
	return e
}
