package business

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"ofspace-be/features/facility"
	mocks "ofspace-be/features/facility/mock"
	"os"
	"testing"
	"time"
)

var (
	facData     mocks.Data
	facBusiness facility.Business
	facValue    facility.Facility
	facValue2   facility.Facility
	facList     []facility.Facility
)

func TestMain(m *testing.M) {
	facBusiness = NewFacilityBusiness(&facData, time.Hour*1)
	facValue = facility.Facility{
		Id:        1,
		Name:      "test",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	facValue2 = facility.Facility{
		Id:        1,
		Name:      "testa",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	facList = []facility.Facility{
		{
			Id:        1,
			Name:      "test",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}
	os.Exit(m.Run())
}

func TestCreateFacility(t *testing.T) {
	t.Run("TestCreateFacility sukses", func(t *testing.T) {
		facData.On("CreateFacility", mock.Anything, mock.Anything).Return(facValue, nil).Once()
		data, err := facBusiness.CreateFacility(context.Background(), facValue)
		assert.NotEqual(t, data, err)
	})
	t.Run("CreateFacility gagal", func(t *testing.T) {
		facData.On("CreateFacility", mock.Anything, mock.Anything).Return(facValue, errors.New("err")).Once()
		data, err := facBusiness.CreateFacility(context.Background(), facValue)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestUpdateFacility(t *testing.T) {
	t.Run("UpdateFacility sukses", func(t *testing.T) {
		facData.On("GetFacility", mock.Anything, mock.Anything).Return(facValue, nil).Once()
		facData.On("UpdateFacility", mock.Anything, mock.Anything).Return(facValue, nil).Once()
		data, err := facBusiness.UpdateFacility(context.Background(), facValue2)
		assert.Equal(t, data, facValue)
		assert.Nil(t, err)
	})
	t.Run("UpdateFacility gagal", func(t *testing.T) {
		facData.On("GetFacility", mock.Anything, mock.Anything).Return(facValue, nil).Once()
		facData.On("UpdateFacility", mock.Anything, mock.Anything).Return(facValue, errors.New("err")).Once()
		data, err := facBusiness.UpdateFacility(context.Background(), facValue2)
		assert.NotEqual(t, data, facValue)
		assert.Error(t, err)
	})
}

func TestGetFacility(t *testing.T) {
	t.Run("GetFacility sukses", func(t *testing.T) {
		facData.On("GetFacility", mock.Anything, mock.AnythingOfType("uint")).Return(facValue, nil).Once()
		data, _ := facBusiness.GetFacility(context.Background(), 1)
		assert.Equal(t, data, facValue)
	})
	t.Run("GetFacility gagal", func(t *testing.T) {
		facData.On("GetFacility", mock.Anything, mock.AnythingOfType("uint")).Return(facValue, errors.New("err")).Once()
		data, err := facBusiness.GetFacility(context.Background(), 3)
		assert.NotEqual(t, data, facValue)
		assert.Error(t, err)
	})
}

func TestDeleteFacility(t *testing.T) {
	t.Run("DeleteFacility sukses", func(t *testing.T) {
		facData.On("GetFacility", mock.Anything, mock.AnythingOfType("uint")).Return(facValue, nil).Once()
		facData.On("DeleteFacility", mock.Anything, mock.AnythingOfType("uint")).Return(facValue, nil).Once()
		data, _ := facBusiness.DeleteFacility(context.Background(), 1)
		assert.Equal(t, data, facValue)
	})
	t.Run("DeleteFacility gagal", func(t *testing.T) {
		facData.On("GetFacility", mock.Anything, mock.AnythingOfType("uint")).Return(facValue, errors.New("err")).Once()
		facData.On("DeleteFacility", mock.Anything, mock.AnythingOfType("uint")).Return(facValue, errors.New("err")).Once()
		data, err := facBusiness.DeleteFacility(context.Background(), 3)
		assert.NotEqual(t, data, facValue)
		assert.Error(t, err)
	})
}

func TestSearchFacility(t *testing.T) {
	t.Run("SearchFacility sukses", func(t *testing.T) {
		facData.On("SearchFacility", mock.Anything, mock.AnythingOfType("string")).Return(facList, nil).Once()
		data, err := facBusiness.SearchFacility(context.Background(), "test")
		assert.Equal(t, data, facList)
		assert.Nil(t, err)
	})
	t.Run("SearchFacility gagal", func(t *testing.T) {
		facData.On("SearchFacility", mock.Anything, mock.AnythingOfType("string")).Return([]facility.Facility{}, errors.New("err")).Once()
		data, err := facBusiness.SearchFacility(context.Background(), "oo")
		assert.Equal(t, data, []facility.Facility{})
		assert.Error(t, err)
	})
}
func TestGetAllFacility(t *testing.T) {
	t.Run("GetAllFacility sukses", func(t *testing.T) {
		facData.On("GetAllFacility", mock.Anything).Return(facList, nil).Once()
		data, err := facBusiness.GetAllFacility(context.Background())
		assert.Equal(t, data, facList)
		assert.Nil(t, err)
	})
	t.Run("GetAllFacility gagal", func(t *testing.T) {
		facData.On("GetAllFacility", mock.Anything).Return([]facility.Facility{}, errors.New("err")).Once()
		data, err := facBusiness.GetAllFacility(context.Background())
		assert.Equal(t, data, []facility.Facility{})
		assert.Error(t, err)
	})
}
