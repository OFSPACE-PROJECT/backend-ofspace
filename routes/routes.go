package routes

import (
	"github.com/labstack/echo/v4"
	"ofspace-be/factory"
)

func New() *echo.Echo {
	presenter := factory.New()
	e := echo.New()
	// complex
	e.POST("/complex", presenter.ComplexPresentation.CreateComplex)
	e.GET("/complex/:complex_id", presenter.ComplexPresentation.GetComplex)
	e.POST("/complex/:complex_id", presenter.ComplexPresentation.RequestComplex)
	e.GET("/complex/search", presenter.ComplexPresentation.SearchComplex)
	e.PUT("/complex/:complex_id", presenter.ComplexPresentation.UpdateComplex)
	return e
}
