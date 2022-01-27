package business

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	complex2 "ofspace-be/features/accessibility"
	mocks "ofspace-be/features/accessibility/mock"
	"os"
	"testing"
	"time"
)

var (
	compData     mocks.Data
	compBusiness complex2.Business
	compValue    complex2.Core
	compValue2   complex2.Core
	compList     []complex2.Core
)

func TestMain(m *testing.M) {
	compBusiness = NewAccessibilityBusiness(&compData, time.Hour*1)
	compValue = complex2.Core{
		Id:        1,
		Name:      "test",
		Address:   "address",
		Latitude:  "1",
		Longitude: "2",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	compValue2 = complex2.Core{
		Id:        2,
		Name:      "test",
		Address:   "address",
		Latitude:  "1",
		Longitude: "2",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	compList = []complex2.Core{
		{
			Id:        1,
			Name:      "test",
			Address:   "test",
			Latitude:  "1",
			Longitude: "2",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}
	os.Exit(m.Run())
}

func TestCreateAccessibility(t *testing.T) {
	t.Run("create complex sukses", func(t *testing.T) {
		compData.On("CreateAccessibility", mock.Anything, mock.Anything).Return(compValue, nil).Once()
		data, err := compBusiness.CreateAccessibility(context.Background(), compValue)
		assert.NotEqual(t, data, err)
	})
	t.Run("create complex gagal", func(t *testing.T) {
		compData.On("CreateAccessibility", mock.Anything, mock.Anything).Return(compValue, errors.New("err")).Once()
		data, err := compBusiness.CreateAccessibility(context.Background(), compValue)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestGetAccessibility(t *testing.T) {
	t.Run("get complex sukses", func(t *testing.T) {
		compData.On("GetAccessibility", mock.Anything, mock.AnythingOfType("uint")).Return(compValue, nil).Once()
		data, err := compBusiness.GetAccessibility(context.Background(), 1)
		assert.Equal(t, data, compValue)
		assert.Nil(t, err)
	})
	t.Run("get complex gagal", func(t *testing.T) {
		compData.On("GetAccessibility", mock.Anything, mock.AnythingOfType("uint")).Return(compValue, errors.New("err")).Once()
		data, err := compBusiness.GetAccessibility(context.Background(), 2)
		assert.Equal(t, data, complex2.Core{})
		assert.Error(t, err)
	})
}

//func TestGetAllAccessibility(t *testing.T) {
//	t.Run("get all complex sukses", func(t *testing.T) {
//		compData.On("GetAllAccessibility", mock.Anything).Return(compList, nil).Once()
//		data, err := compBusiness.GetAllAccessibility(context.Background())
//		assert.Equal(t, data, compList)
//		assert.Nil(t, err)
//	})
//	t.Run("get all complex gagal", func(t *testing.T) {
//		compData.On("GetAllAccessibility", mock.Anything).Return(compList, errors.New("err")).Once()
//		data, err := compBusiness.GetAllAccessibility(context.Background())
//		assert.Equal(t, data, []complex2.Core{})
//		assert.Error(t, err)
//	})
//}

func TestSearchAccessibilityByAddress(t *testing.T) {
	t.Run("SearchAccessibilityByAddress sukses", func(t *testing.T) {
		compData.On("SearchAccessibility", mock.Anything, mock.AnythingOfType("string")).Return(compList, nil).Once()
		compData.On("SearchAccessibilityByAddress", mock.Anything, mock.AnythingOfType("string")).Return(compList, nil).Once()
		data, err := compBusiness.SearchAccessibilityByAddress(context.Background(), compList[0].Address)
		assert.Equal(t, data, compList)
		assert.Nil(t, err)
	})

	t.Run("SearchAccessibilityByAddress gagal", func(t *testing.T) {
		compData.On("SearchAccessibility", mock.Anything, mock.AnythingOfType("string")).Return(compList, errors.New("err")).Once()
		compData.On("SearchAccessibilityByAddress", mock.Anything, mock.AnythingOfType("string")).Return(compList, errors.New("err")).Once()
		data, err := compBusiness.SearchAccessibilityByAddress(context.Background(), "address")
		assert.Equal(t, data, []complex2.Core{})
		assert.Error(t, err)
	})
}

func TestSearchAccessibility(t *testing.T) {
	t.Run("search complex sukses", func(t *testing.T) {
		compData.On("SearchAccessibility", mock.Anything, mock.AnythingOfType("string")).Return(compList, nil).Once()
		data, err := compBusiness.SearchAccessibility(context.Background(), "test")
		assert.Equal(t, data, compList)
		assert.Nil(t, err)
	})
	t.Run("search complex gagal", func(t *testing.T) {
		compData.On("SearchAccessibility", mock.Anything, mock.AnythingOfType("string")).Return(compList, errors.New("err")).Once()
		data, err := compBusiness.SearchAccessibility(context.Background(), "gagal")
		assert.Equal(t, data, []complex2.Core{})
		assert.Error(t, err)
	})
}

func TestUpdateAccessibility(t *testing.T) {
	t.Run("UpdateAccessibility sukses", func(t *testing.T) {
		compData.On("GetAccessibility", mock.Anything, mock.AnythingOfType("uint")).Return(compValue, nil).Once()
		compData.On("UpdateAccessibility", mock.Anything, mock.Anything).Return(compValue2, nil).Once()
		data, err := compBusiness.UpdateAccessibility(context.Background(), compValue2)
		assert.Equal(t, data, compValue2)
		assert.Nil(t, err)
	})
	t.Run("UpdateAccessibility gagal", func(t *testing.T) {
		compData.On("GetAccessibility", mock.Anything, mock.AnythingOfType("uint")).Return(compValue, errors.New("err")).Once()
		compData.On("UpdateAccessibility", mock.Anything, mock.Anything).Return(compValue2, errors.New("err")).Once()
		data, err := compBusiness.UpdateAccessibility(context.Background(), compValue2)
		assert.Equal(t, data, complex2.Core{})
		assert.Error(t, err)
	})
}
