package business

import (
	"ofspace-be/features/booking"
	mocks "ofspace-be/features/booking/mock"
	"os"
	"testing"
	"time"
)

var (
	compData     mocks.Data
	compBusiness booking.Business
	compValue    booking.Core
	compValue2   booking.Core
	compList     []booking.Core
)

func TestMain(m *testing.M) {
	compBusiness = NewBookingBusiness(&compData, time.Hour*1)
	compValue = booking.Core{
		ID:            1,
		CostumerId:    1,
		ConsultantId:  1,
		User:          booking.User{},
		BuildingId:    0,
		UnitId:        0,
		Unit:          booking.Unit{},
		Building:      booking.Building{},
		ConfirmedName: "",
		TotalBought:   0,
		Price:         0,
		DealDate:      time.Time{},
		StartDate:     time.Time{},
		EndDate:       time.Time{},
		PaymentStatus: "",
		BookingStatus: "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	}

	compValue2 = booking.Core{
		ID:            1,
		CostumerId:    1,
		ConsultantId:  1,
		User:          booking.User{},
		BuildingId:    0,
		UnitId:        0,
		Unit:          booking.Unit{},
		Building:      booking.Building{},
		ConfirmedName: "",
		TotalBought:   0,
		Price:         0,
		DealDate:      time.Time{},
		StartDate:     time.Time{},
		EndDate:       time.Time{},
		PaymentStatus: "",
		BookingStatus: "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	}
	compList = []booking.Core{
		{
			ID:            1,
			CostumerId:    1,
			ConsultantId:  1,
			User:          booking.User{},
			BuildingId:    0,
			UnitId:        0,
			Unit:          booking.Unit{},
			Building:      booking.Building{},
			ConfirmedName: "",
			TotalBought:   0,
			Price:         0,
			DealDate:      time.Time{},
			StartDate:     time.Time{},
			EndDate:       time.Time{},
			PaymentStatus: "",
			BookingStatus: "",
			CreatedAt:     time.Time{},
			UpdatedAt:     time.Time{},
		},
	}

	os.Exit(m.Run())
}

//func TestCreateBooking(t *testing.T) {
//	t.Run("create complex sukses", func(t *testing.T) {
//		compData.On("CreateBooking", mock.Anything, mock.Anything).Return(compValue, nil).Once()
//		data, err := compBusiness.CreateBooking(context.Background(), compValue)
//		assert.NotEqual(t, data, err)
//	})
//	t.Run("create complex gagal", func(t *testing.T) {
//		compData.On("CreateBooking", mock.Anything, mock.Anything).Return(compValue, errors.New("err")).Once()
//		data, err := compBusiness.CreateBooking(context.Background(), compValue)
//		assert.NotEqual(t, data, err)
//		assert.Error(t, err)
//	})
//}

//func TestGetBooking(t *testing.T) {
//	t.Run("get complex sukses", func(t *testing.T) {
//		compData.On("GetOneBooking", mock.Anything, mock.AnythingOfType("uint")).Return(compValue, nil).Once()
//		data, err := compBusiness.GetOneBooking(context.Background(), 1)
//		assert.Equal(t, data, compValue)
//		assert.Nil(t, err)
//	})
//	t.Run("get complex gagal", func(t *testing.T) {
//		compData.On("GetOneBooking", mock.Anything, mock.AnythingOfType("uint")).Return(compValue, errors.New("err")).Once()
//		data, err := compBusiness.GetOneBooking(context.Background(), 2)
//		assert.Equal(t, data, booking.Core{})
//		assert.Error(t, err)
//	})
//}
//
//func TestSearchBookingByName(t *testing.T) {
//	t.Run("SearchBookingByName sukses", func(t *testing.T) {
//		compData.On("SearchBookingByName", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(compList, nil).Once()
//		data, err := compBusiness.SearchBookingByName(context.Background(), 1, compList[0].ConfirmedName)
//		assert.Equal(t, data, compList)
//		assert.Nil(t, err)
//	})
//
//	t.Run("SearchBookingByPaymentStatus gagal", func(t *testing.T) {
//		compData.On("SearchBookingByPaymentStatus", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(compList, errors.New("err")).Once()
//		data, err := compBusiness.SearchBookingByPaymentStatus(context.Background(), 1, "address")
//		assert.Equal(t, data, []booking.Core{})
//		assert.Error(t, err)
//	})
//}
//func TestSearchBookingByPaymentStatus(t *testing.T) {
//	t.Run("SearchBookingByPaymentStatus sukses", func(t *testing.T) {
//		compData.On("SearchBookingByPaymentStatus", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(compList, nil).Once()
//		data, err := compBusiness.SearchBookingByPaymentStatus(context.Background(), 1, compList[0].PaymentStatus)
//		assert.Equal(t, data, compList)
//		assert.Nil(t, err)
//	})
//
//	t.Run("SearchBookingByPaymentStatus gagal", func(t *testing.T) {
//		compData.On("SearchBookingByPaymentStatus", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(compList, errors.New("err")).Once()
//		data, err := compBusiness.SearchBookingByPaymentStatus(context.Background(), 1, "address")
//		assert.Equal(t, data, []booking.Core{})
//		assert.Error(t, err)
//	})
//}
//
//func TestUpdateBooking(t *testing.T) {
//	t.Run("UpdateBooking sukses", func(t *testing.T) {
//		compData.On("GetBooking", mock.Anything, mock.AnythingOfType("uint")).Return(compValue, nil).Once()
//		compData.On("UpdateBooking", mock.Anything, mock.Anything).Return(compValue2, nil).Once()
//		data, err := compBusiness.UpdateBooking(context.Background(), compValue2)
//		assert.Equal(t, data, compValue2)
//		assert.Nil(t, err)
//	})
//	t.Run("UpdateBooking gagal", func(t *testing.T) {
//		compData.On("GetBooking", mock.Anything, mock.AnythingOfType("uint")).Return(compValue, errors.New("err")).Once()
//		compData.On("UpdateBooking", mock.Anything, mock.Anything).Return(compValue2, errors.New("err")).Once()
//		data, err := compBusiness.UpdateBooking(context.Background(), compValue2)
//		assert.Equal(t, data, booking.Core{})
//		assert.Error(t, err)
//	})
//}
