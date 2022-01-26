package routes

import (
	"ofspace-be/factory"
	mid "ofspace-be/middleware"
	"os"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	presenter := factory.New()
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	mid.Logger(e)
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
	JWTSecret := os.Getenv("JWT_SECRET")
	iJWT := e.Group("")
	iJWT.Use(middleware.JWT([]byte(JWTSecret)))
	// user
	e.POST("/register", presenter.UserPresentation.RegisterUser)
	e.POST("/login", presenter.UserPresentation.LoginUser)
	iJWT.GET("/user/:id", presenter.UserPresentation.GetUserByID)
	iJWT.PUT("/user", presenter.UserPresentation.UpdateUser)
	iJWT.DELETE("/user/:id", presenter.UserPresentation.DeleteUser)
	iJWT.GET("/user/search", presenter.UserPresentation.SearchUserByName)
	iJWT.GET("/user/status", presenter.UserPresentation.SearchUserByAdminStatus)
	iJWT.GET("/user", presenter.UserPresentation.GetAllUser)

	// complex
	iJWT.POST("/complex", presenter.ComplexPresentation.CreateComplex)
	e.GET("/complex/:id", presenter.ComplexPresentation.GetComplex)
	e.GET("/complex", presenter.ComplexPresentation.GetAllComplex)
	iJWT.POST("/complex/:id", presenter.ComplexPresentation.RequestComplex)
	e.GET("/complex/search", presenter.ComplexPresentation.SearchComplex)
	e.GET("/complex/address", presenter.ComplexPresentation.SearchComplexByAddress)
	iJWT.PUT("/complex", presenter.ComplexPresentation.UpdateComplex)

	// accessibility
	iJWT.POST("/accessibility", presenter.AccessibilityPresentation.CreateAccessibility)
	e.GET("/accessibility/:id", presenter.AccessibilityPresentation.GetAccessibility)
	iJWT.POST("/accessibility/:id", presenter.AccessibilityPresentation.RequestAccessibility)
	e.GET("/accessibility/search", presenter.AccessibilityPresentation.SearchAccessibility)
	e.GET("/accessibility/address", presenter.AccessibilityPresentation.SearchAccessibilityByAddress)
	iJWT.PUT("/accessibility", presenter.AccessibilityPresentation.UpdateAccessibility)

	// facility
	iJWT.POST("/facility", presenter.FacilityPresentation.CreateFacility)
	e.GET("/facility/:id", presenter.FacilityPresentation.GetFacility)
	e.GET("/facility", presenter.FacilityPresentation.GetAllFacility)
	e.GET("/facility/search", presenter.FacilityPresentation.SearchFacility)
	iJWT.PUT("/facility", presenter.FacilityPresentation.UpdateFacility)
	iJWT.DELETE("/facility/:id", presenter.FacilityPresentation.DeleteFacility)

	//building
	e.POST("/building", presenter.BuildingPresentation.CreateBuilding)
	e.GET("/building/:id", presenter.BuildingPresentation.GetBuildingById)
	iJWT.GET("/building/admin", presenter.BuildingPresentation.GetAllBuilding)
	e.GET("/building", presenter.BuildingPresentation.GetAllVerifiedBuilding)
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
	e.GET("/building/:id/floor/:photo_id", presenter.BuildingPresentation.GetFloorPhoto)
	e.GET("/building/:id/floor", presenter.BuildingPresentation.GetAllFloorPhoto)
	iJWT.PUT("/building/:building_id/floor/:photo_id", presenter.BuildingPresentation.UpdateFloorPhoto)
	iJWT.DELETE("/building/:id/floor/:photo_id", presenter.BuildingPresentation.DeleteFloorPhoto)
	//building facilities
	//iJWT.POST("/building/facility", presenter.BuildingPresentation.AddFacilityToBuilding)
	iJWT.POST("/building/facility", presenter.BuildingPresentation.AddFacilityToBuilding)
	e.GET("/building/:id/facility/:facility_id", presenter.BuildingPresentation.GetBuildingFacility)
	e.GET("/building/:id/facility", presenter.BuildingPresentation.GetAllBuildingFacility)
	//iJWT.DELETE("/building/:id/facility/:photo_id", presenter.BuildingPresentation.DeleteFacility)
	iJWT.DELETE("/building/:id/facility/:facility_id", presenter.BuildingPresentation.DeleteFacility)

	//unit
	iJWT.POST("/unit", presenter.UnitPresentation.CreateUnit)
	e.GET("/unit/:id", presenter.UnitPresentation.GetUnitById)
	e.GET("/unit", presenter.UnitPresentation.GetAllUnit)
	e.GET("/unit/:id/type", presenter.UnitPresentation.GetUnitByType)
	iJWT.PUT("/unit", presenter.UnitPresentation.UpdateUnit)
	//unit interior photo
	iJWT.POST("/unit/interior", presenter.UnitPresentation.CreateInteriorPhoto)
	e.GET("/unit/:id/interior/:photo_id", presenter.UnitPresentation.GetInteriorPhoto)
	e.GET("/unit/:id/interior", presenter.UnitPresentation.GetAllInteriorPhoto)
	iJWT.PUT("/unit/:id/interior/:photo_id", presenter.UnitPresentation.UpdateInteriorPhoto)
	iJWT.DELETE("/unit/:id/interior/:photo_id", presenter.UnitPresentation.DeleteInteriorPhoto)
	//unit facilities
	iJWT.POST("/unit/facility", presenter.UnitPresentation.AddFacilityToUnit)
	e.GET("/unit/:id/facility/:facility_id", presenter.UnitPresentation.GetUnitFacility)
	e.GET("/unit/:id/facility", presenter.UnitPresentation.GetAllUnitFacility)
	iJWT.DELETE("/unit/:id/facility/:photo_id", presenter.UnitPresentation.DeleteUnitFacility)

	//wishlist
	iJWT.POST("/wishlist", presenter.WishlistPresentation.CreateWishlist)
	iJWT.GET("/wishlist/:id", presenter.WishlistPresentation.GetWishlist)
	iJWT.GET("/wishlist", presenter.WishlistPresentation.GetAllWishlists)
	iJWT.DELETE("/wishlist/:id", presenter.WishlistPresentation.DeleteWishlist)

	// review
	iJWT.POST("/review", presenter.ReviewPresentation.CreateReview)
	e.GET("/review/:id", presenter.ReviewPresentation.GetOneReview)
	e.GET("/review/all/:unit", presenter.ReviewPresentation.GetAllReview)
	iJWT.PUT("/review", presenter.ReviewPresentation.UpdateReview)

	//booking
	iJWT.POST("/booking", presenter.BookingPresentation.CreateBooking)
	iJWT.PUT("/booking", presenter.BookingPresentation.UpdateBooking)
	iJWT.GET("/booking", presenter.BookingPresentation.GetAllBooking)
	iJWT.GET("/booking/:id", presenter.BookingPresentation.GetOneBooking)
	iJWT.GET("/booking/:id/name", presenter.BookingPresentation.SearchBookingByName)
	iJWT.GET("/booking/:id/payment", presenter.BookingPresentation.SearchBookingByPayment)
	iJWT.GET("/booking/:id/status", presenter.BookingPresentation.GetBookingByStatus)
	iJWT.GET("/booking/:id/date", presenter.BookingPresentation.FindBookingByDate)
	iJWT.GET("/booking/:id/sum", presenter.BookingPresentation.GetSumOfTotalBoughtInUnit)
	iJWT.GET("/booking/:id/earning", presenter.BookingPresentation.GetEarningsInUnitWithDateFilter)
	iJWT.GET("/booking/:id/sumpayment", presenter.BookingPresentation.GetSumOfPaymentConfirmed)
	iJWT.GET("/booking/user", presenter.BookingPresentation.GetAllBookingByUser)
	iJWT.GET("/booking/building", presenter.BookingPresentation.GetAllBookingByBuilding)
	iJWT.GET("/booking/unit", presenter.BookingPresentation.GetAllBookingByUnit)

	return e
}
