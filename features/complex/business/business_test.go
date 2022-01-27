package business

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	complex2 "ofspace-be/features/complex"
	mocks "ofspace-be/features/complex/mock"
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
	ctxMain      time.Duration
)

func TestMain(m *testing.M) {
	compBusiness = NewComplexBusiness(&compData, time.Hour*1)
	ctxMain = time.Hour * 1
	compValue = complex2.Core{
		Id:        1,
		Name:      "test",
		Address:   "address",
		Latitude:  "1",
		Longitude: "2",
		Buildings: nil,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	compValue2 = complex2.Core{
		Id:        2,
		Name:      "test",
		Address:   "address",
		Latitude:  "1",
		Longitude: "2",
		Buildings: nil,
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
			Buildings: nil,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}
	os.Exit(m.Run())
}

func TestCreateComplex(t *testing.T) {
	t.Run("create complex sukses", func(t *testing.T) {
		compData.On("CreateComplex", mock.Anything, mock.Anything).Return(compValue, nil).Once()
		data, err := compBusiness.CreateComplex(context.Background(), compValue)
		assert.NotEqual(t, data, err)
	})
	t.Run("create complex gagal", func(t *testing.T) {
		compData.On("CreateComplex", mock.Anything, mock.Anything).Return(compValue, errors.New("err")).Once()
		data, err := compBusiness.CreateComplex(context.Background(), compValue)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestGetComplex(t *testing.T) {
	t.Run("get complex sukses", func(t *testing.T) {
		compData.On("GetComplex", mock.Anything, mock.AnythingOfType("uint")).Return(compValue, nil).Once()
		data, err := compBusiness.GetComplex(context.Background(), 1)
		assert.Equal(t, data, compValue)
		assert.Nil(t, err)
	})
	t.Run("get complex gagal", func(t *testing.T) {
		compData.On("GetComplex", mock.Anything, mock.AnythingOfType("uint")).Return(compValue, errors.New("err")).Once()
		data, err := compBusiness.GetComplex(context.Background(), 2)
		assert.Equal(t, data, complex2.Core{})
		assert.Error(t, err)
	})
}
func TestGetAllComplex(t *testing.T) {
	t.Run("get all complex sukses", func(t *testing.T) {
		compData.On("GetAllComplex", mock.Anything).Return(compList, nil).Once()
		data, err := compBusiness.GetAllComplex(context.Background())
		assert.Equal(t, data, compList)
		assert.Nil(t, err)
	})
	t.Run("get all complex gagal", func(t *testing.T) {
		compData.On("GetAllComplex", mock.Anything).Return(compList, errors.New("err")).Once()
		data, err := compBusiness.GetAllComplex(context.Background())
		assert.Equal(t, data, []complex2.Core{})
		assert.Error(t, err)
	})
}

func TestSearchComplexByAddress(t *testing.T) {
	t.Run("SearchComplexByAddress sukses", func(t *testing.T) {
		compData.On("SearchComplex", mock.Anything, mock.AnythingOfType("string")).Return(compList, nil).Once()
		compData.On("SearchComplexByAddress", mock.Anything, mock.AnythingOfType("string")).Return(compList, nil).Once()
		data, err := compBusiness.SearchComplexByAddress(context.Background(), compList[0].Address)
		assert.Equal(t, data, compList)
		assert.Nil(t, err)
	})

	t.Run("SearchComplexByAddress gagal", func(t *testing.T) {
		compData.On("SearchComplex", mock.Anything, mock.AnythingOfType("string")).Return(compList, errors.New("err")).Once()
		compData.On("SearchComplexByAddress", mock.Anything, mock.AnythingOfType("string")).Return(compList, errors.New("err")).Once()
		data, err := compBusiness.SearchComplexByAddress(context.Background(), "address")
		assert.Equal(t, data, []complex2.Core{})
		assert.Error(t, err)
	})
}

func TestSearchComplex(t *testing.T) {
	t.Run("search complex sukses", func(t *testing.T) {
		compData.On("SearchComplex", mock.Anything, mock.AnythingOfType("string")).Return(compList, nil).Once()
		data, err := compBusiness.SearchComplex(context.Background(), "test")
		assert.Equal(t, data, compList)
		assert.Nil(t, err)
	})
	t.Run("search complex gagal", func(t *testing.T) {
		compData.On("SearchComplex", mock.Anything, mock.AnythingOfType("string")).Return(compList, errors.New("err")).Once()
		data, err := compBusiness.SearchComplex(context.Background(), "gagal")
		assert.Equal(t, data, []complex2.Core{})
		assert.Error(t, err)
	})
}

func TestUpdateComplex(t *testing.T) {
	t.Run("UpdateComplex sukses", func(t *testing.T) {
		compData.On("GetComplex", mock.Anything, mock.AnythingOfType("uint")).Return(compValue, nil).Once()
		compData.On("UpdateComplex", mock.Anything, mock.Anything).Return(compValue2, nil).Once()
		data, err := compBusiness.UpdateComplex(context.Background(), compValue2)
		assert.Equal(t, data, compValue2)
		assert.Nil(t, err)
	})
	t.Run("UpdateComplex gagal", func(t *testing.T) {
		compData.On("GetComplex", mock.Anything, mock.AnythingOfType("uint")).Return(compValue, errors.New("err")).Once()
		compData.On("UpdateComplex", mock.Anything, mock.Anything).Return(compValue2, errors.New("err")).Once()
		data, err := compBusiness.UpdateComplex(context.Background(), compValue2)
		assert.Equal(t, data, complex2.Core{})
		assert.Error(t, err)
	})
}