package presentation

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ofspace-be/features/booking"
	"ofspace-be/features/booking/presentation/request"
	"ofspace-be/features/booking/presentation/response"
	"strconv"
)

type BookingPresentation struct {
	bookingBusiness booking.Business
}

func NewBookingPresentation(wb booking.Business) *BookingPresentation {
	return &BookingPresentation{bookingBusiness: wb}
}

func (bp *BookingPresentation) CreateBooking(c echo.Context) error {
	CreateBooking := request.CreateBooking{}
	err := c.Bind(&CreateBooking)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	data, err2 := bp.bookingBusiness.CreateBooking(ctx, CreateBooking.ToCore())
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.FromBookingCore(data),
	})
}
func (bp *BookingPresentation) UpdateBooking(c echo.Context) error {
	cUpdate := request.UpdateBooking{}
	err := c.Bind(&cUpdate)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	data, err2 := bp.bookingBusiness.UpdateBooking(ctx, cUpdate.ToUpdateCore())
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.FromBookingCore(data),
	})
}
func (bp *BookingPresentation) GetAllBooking(c echo.Context) error {
	ctx := c.Request().Context()
	fac, err := bp.bookingBusiness.GetAllBooking(ctx)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.FromListBookingCore(fac),
	})
}
func (bp *BookingPresentation) GetAllBookingByUnit(c echo.Context) error {
	id := c.QueryParam("unit_id")
	ids, err0 := strconv.Atoi(id)
	if err0 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err0.Error(),
		})
	}
	ctx := c.Request().Context()
	fac, err := bp.bookingBusiness.GetAllBookingByUnit(ctx, uint(ids))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.FromListBookingCore(fac),
	})
}

func (bp *BookingPresentation) GetAllBookingByBuilding(c echo.Context) error {
	id := c.QueryParam("building_id")
	ids, err0 := strconv.Atoi(id)
	if err0 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err0.Error(),
		})
	}
	ctx := c.Request().Context()
	fac, err := bp.bookingBusiness.GetAllBookingByBuilding(ctx, uint(ids))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.FromListBookingCore(fac),
	})
}

func (bp *BookingPresentation) GetAllBookingByUser(c echo.Context) error {
	id := c.QueryParam("user_id")
	ids, err0 := strconv.Atoi(id)
	if err0 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err0.Error(),
		})
	}
	ctx := c.Request().Context()
	fac, err := bp.bookingBusiness.GetAllBookingByUser(ctx, uint(ids))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.FromListBookingCore(fac),
	})
}
func (bp *BookingPresentation) GetOneBooking(c echo.Context) error {
	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	ctx := c.Request().Context()
	comp, err2 := bp.bookingBusiness.GetOneBooking(ctx, uint(Id))
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.FromBookingCore(comp),
	})
}
func (bp *BookingPresentation) SearchBookingByName(c echo.Context) error {
	name := c.QueryParam("confirmed_name")
	buildingId, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()
	fac, err := bp.bookingBusiness.SearchBookingByName(ctx, uint(buildingId), name)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.FromListBookingCore(fac),
	})
}

func (bp *BookingPresentation) GetBookingByStatus(c echo.Context) error {
	status := c.QueryParam("booking_status")
	buildingId, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()
	fac, err := bp.bookingBusiness.SearchBookingByName(ctx, uint(buildingId), status)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.FromListBookingCore(fac),
	})
}

func (bp *BookingPresentation) SearchBookingByPayment(c echo.Context) error {
	id := c.QueryParam("payment_status")
	buildingId, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()
	fac, err := bp.bookingBusiness.SearchBookingByPayment(ctx, uint(buildingId), id)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.FromListBookingCore(fac),
	})
}

//func (bp *BookingPresentation) GetSumOfTotalBoughtInUnit(c echo.Context) error {
//	unitId, _ := strconv.Atoi(c.Param("id"))
//	ctx := c.Request().Context()
//	fac, err := bp.bookingBusiness.GetSumOfTotalBoughtInUnit(ctx, uint(unitId))
//	if err != nil {
//		return c.JSON(http.StatusForbidden, map[string]interface{}{
//			"message": err.Error(),
//		})
//	}
//
//	return c.JSON(http.StatusOK, map[string]interface{}{
//		"message": "Success",
//		"data":    fac,
//	})
//}
//
//func (bp *BookingPresentation) GetEarningsInUnitWithDateFilter(c echo.Context) error {
//	unitId, _ := strconv.Atoi(c.Param("id"))
//	layoutFormat := "2006-01-02"
//	value1 := c.QueryParam("start_date")
//	value2 := c.QueryParam("end_date")
//	date1, _ := time.Parse(layoutFormat, value1)
//	date2, _ := time.Parse(layoutFormat, value2)
//	ctx := c.Request().Context()
//	fac, err := bp.bookingBusiness.GetEarningsInUnitWithDateFilter(ctx, uint(unitId), date1, date2)
//	if err != nil {
//		return c.JSON(http.StatusForbidden, map[string]interface{}{
//			"message": err.Error(),
//		})
//	}
//
//	return c.JSON(http.StatusOK, map[string]interface{}{
//		"message": "Success",
//		"data":    fac,
//	})
//}
//
//func (bp *BookingPresentation) GetSumOfPaymentConfirmed(c echo.Context) error {
//	unitId, _ := strconv.Atoi(c.Param("id"))
//	layoutFormat := "2006-01-02"
//	value1 := c.QueryParam("start_date")
//	value2 := c.QueryParam("end_date")
//	date1, _ := time.Parse(layoutFormat, value1)
//	date2, _ := time.Parse(layoutFormat, value2)
//	ctx := c.Request().Context()
//	fac, err := bp.bookingBusiness.GetSumOfPaymentConfirmed(ctx, uint(unitId), date1, date2)
//	if err != nil {
//		return c.JSON(http.StatusForbidden, map[string]interface{}{
//			"message": err.Error(),
//		})
//	}
//
//	return c.JSON(http.StatusOK, map[string]interface{}{
//		"message": "Success",
//		"data":    fac,
//	})
//}
//
//func (bp *BookingPresentation) FindBookingByDate(c echo.Context) error {
//	layoutFormat := "2006-01-02"
//	value1 := c.QueryParam("start_date")
//	value2 := c.QueryParam("end_date")
//	var date1 time.Time
//	var date2 time.Time
//	date1, _ = time.Parse(layoutFormat, value1)
//	date2, _ = time.Parse(layoutFormat, value2)
//	buildingId, _ := strconv.Atoi(c.Param("id"))
//	ctx := c.Request().Context()
//	fac, err := bp.bookingBusiness.FindBookingByDate(ctx, uint(buildingId), date1, date2)
//	if err != nil {
//		return c.JSON(http.StatusForbidden, map[string]interface{}{
//			"message": err.Error(),
//		})
//	}
//
//	return c.JSON(http.StatusOK, map[string]interface{}{
//		"message": "Success",
//		"data":    response.FromListBookingCore(fac),
//	})
//}
