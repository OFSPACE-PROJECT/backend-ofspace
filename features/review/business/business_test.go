package business

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"ofspace-be/features/review"
	mocks "ofspace-be/features/review/mock"
	"os"
	"testing"
	"time"
)

var (
	facData     mocks.Data
	facBusiness review.Business
	facValue    review.Core
	facValue2   review.Core
	facList     []review.Core
)

func TestMain(m *testing.M) {
	facBusiness = NewReviewBusiness(&facData, time.Hour*1)
	facValue = review.Core{
		Id:                    1,
		CustomerId:            2,
		Customer:              review.User{},
		UnitId:                2,
		BuildingId:            2,
		Unit:                  review.Unit{},
		BookingId:             1,
		RatingAccess:          0,
		RatingFacility:        0,
		RatingManagement:      0,
		RatingQuality:         0,
		CostumerOverallRating: 0,
		Comment:               "a",
		CreatedAt:             time.Time{},
		UpdatedAt:             time.Time{},
	}
	facValue2 = review.Core{
		Id:                    2,
		CustomerId:            2,
		Customer:              review.User{},
		UnitId:                1,
		BuildingId:            3,
		Unit:                  review.Unit{},
		BookingId:             2,
		RatingAccess:          0,
		RatingFacility:        0,
		RatingManagement:      0,
		RatingQuality:         0,
		CostumerOverallRating: 0,
		Comment:               "sda",
		CreatedAt:             time.Time{},
		UpdatedAt:             time.Time{},
	}
	facList = []review.Core{
		{
			Id:                    2,
			CustomerId:            2,
			Customer:              review.User{},
			UnitId:                1,
			BuildingId:            3,
			Unit:                  review.Unit{},
			BookingId:             2,
			RatingAccess:          0,
			RatingFacility:        0,
			RatingManagement:      0,
			RatingQuality:         0,
			CostumerOverallRating: 0,
			Comment:               "sda",
			CreatedAt:             time.Time{},
			UpdatedAt:             time.Time{},
		},
	}
	os.Exit(m.Run())
}

func TestCreateReview(t *testing.T) {
	t.Run("CreateReview sukses", func(t *testing.T) {
		facData.On("CreateReview", mock.Anything, mock.Anything).Return(facValue, nil).Once()
		data, err := facBusiness.CreateReview(context.Background(), facValue)
		assert.NotEqual(t, data, err)
	})
	t.Run("CreateReview gagal", func(t *testing.T) {
		facData.On("CreateReview", mock.Anything, mock.Anything).Return(facValue, errors.New("err")).Once()
		data, err := facBusiness.CreateReview(context.Background(), facValue)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestUpdateReview(t *testing.T) {
	t.Run("UpdateReview sukses", func(t *testing.T) {
		facData.On("GetOneReview", mock.Anything, mock.Anything).Return(facValue, nil).Once()
		facData.On("UpdateReview", mock.Anything, mock.Anything).Return(facValue, nil).Once()
		data, err := facBusiness.UpdateReview(context.Background(), facValue2)
		assert.Equal(t, data, facValue)
		assert.Nil(t, err)
	})
	t.Run("UpdateReview gagal", func(t *testing.T) {
		facData.On("GetOneReview", mock.Anything, mock.Anything).Return(facValue, nil).Once()
		facData.On("UpdateReview", mock.Anything, mock.Anything).Return(facValue, errors.New("err")).Once()
		data, err := facBusiness.UpdateReview(context.Background(), facValue2)
		assert.NotEqual(t, data, facValue)
		assert.Error(t, err)
	})
}

func TestGetOneReview(t *testing.T) {
	t.Run("GetOneReview sukses", func(t *testing.T) {
		facData.On("GetOneReview", mock.Anything, mock.AnythingOfType("uint")).Return(facValue, nil).Once()
		data, _ := facBusiness.GetOneReview(context.Background(), 1)
		assert.Equal(t, data, facValue)
	})
	t.Run("GetOneReview gagal", func(t *testing.T) {
		facData.On("GetOneReview", mock.Anything, mock.AnythingOfType("uint")).Return(facValue, errors.New("err")).Once()
		data, err := facBusiness.GetOneReview(context.Background(), 3)
		assert.NotEqual(t, data, facValue)
		assert.Error(t, err)
	})
}

func TestGetAllReview(t *testing.T) {
	t.Run("GetAllReview sukses", func(t *testing.T) {
		facData.On("GetAllReview", mock.Anything, mock.AnythingOfType("uint")).Return(facList, nil).Once()
		data, err := facBusiness.GetAllReview(context.Background(), 1)
		assert.Equal(t, data, facList)
		assert.Nil(t, err)
	})
	t.Run("GetAllReview gagal", func(t *testing.T) {
		facData.On("GetAllReview", mock.Anything, mock.AnythingOfType("uint")).Return([]review.Core{}, errors.New("err")).Once()
		data, err := facBusiness.GetAllReview(context.Background(), 4)
		assert.Equal(t, data, []review.Core{})
		assert.Error(t, err)
	})
}
