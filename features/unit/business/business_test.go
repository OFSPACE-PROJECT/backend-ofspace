package business

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"ofspace-be/features/booking"
	"ofspace-be/features/building"
	"ofspace-be/features/facility"
	"ofspace-be/features/unit"
	mocks "ofspace-be/features/unit/mock"
	"ofspace-be/features/users"
	"os"
	"testing"
	"time"
)

var (
	facData          mocks.Data
	facBusiness      unit.Business
	facValue         unit.Core
	facValue2        unit.Core
	usBusiness       users.Business
	buBusiness       building.Business
	bookBusiness     booking.Business
	facilityBusiness facility.Business

	facList []unit.Core
)

func TestMain(m *testing.M) {
	facBusiness = NewUnitBusiness(&facData, buBusiness, usBusiness, facilityBusiness, bookBusiness, time.Hour*1)
	facValue = unit.Core{
		Id:             1,
		UserId:         1,
		BuildingId:     1,
		Description:    "",
		UnitType:       "",
		Price:          1,
		TotalUnit:      1,
		RemainingUnit:  1,
		Reviews:        nil,
		UnitFacilities: nil,
		InteriorPhoto:  nil,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}
	facValue2 = unit.Core{
		Id:             2,
		UserId:         2,
		BuildingId:     1,
		Description:    "",
		UnitType:       "",
		Price:          1,
		TotalUnit:      1,
		RemainingUnit:  1,
		Reviews:        nil,
		UnitFacilities: nil,
		InteriorPhoto:  nil,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}
	facList = []unit.Core{
		{
			Id:             2,
			UserId:         2,
			BuildingId:     1,
			Description:    "",
			UnitType:       "",
			Price:          1,
			TotalUnit:      1,
			RemainingUnit:  1,
			Reviews:        nil,
			UnitFacilities: nil,
			InteriorPhoto:  nil,
			CreatedAt:      time.Time{},
			UpdatedAt:      time.Time{},
		},
	}
	os.Exit(m.Run())
}

func TestCreateUnit(t *testing.T) {
	t.Run("CreateUnit sukses", func(t *testing.T) {
		facData.On("CreateUnit", mock.Anything, mock.Anything).Return(facValue, nil).Once()
		data, err := facBusiness.CreateUnit(context.Background(), facValue)
		assert.NotEqual(t, data, err)
	})
	t.Run("CreateUnit gagal", func(t *testing.T) {
		facData.On("CreateUnit", mock.Anything, mock.Anything).Return(facValue, errors.New("err")).Once()
		data, err := facBusiness.CreateUnit(context.Background(), facValue)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestDeleteInteriorPhoto(t *testing.T) {
	t.Run("UpdateInteriorPhoto sukses", func(t *testing.T) {
		facData.On("GetInteriorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(unit.InteriorCore{}, nil).Once()
		facData.On("DeleteInteriorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(unit.InteriorCore{}, nil).Once()
		data, err := facBusiness.DeleteInteriorPhoto(context.Background(), 1, 2)
		assert.NotEqual(t, data, err)
	})
	t.Run("UpdateInteriorPhoto gagal", func(t *testing.T) {
		facData.On("GetInteriorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(unit.InteriorCore{}, errors.New("err")).Once()
		facData.On("DeleteInteriorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(unit.InteriorCore{}, errors.New("err")).Once()
		data, err := facBusiness.DeleteInteriorPhoto(context.Background(), 1, 2)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestUpdateInteriorPhoto(t *testing.T) {
	t.Run("UpdateInteriorPhoto sukses", func(t *testing.T) {
		facData.On("GetInteriorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(unit.InteriorCore{}, nil).Once()
		facData.On("UpdateInteriorPhoto", mock.Anything, mock.Anything).Return(unit.InteriorCore{}, nil).Once()
		data, err := facBusiness.UpdateInteriorPhoto(context.Background(), unit.InteriorCore{})
		assert.NotEqual(t, data, err)
	})
	t.Run("UpdateInteriorPhoto gagal", func(t *testing.T) {
		facData.On("GetInteriorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(unit.InteriorCore{}, errors.New("err")).Once()
		facData.On("UpdateInteriorPhoto", mock.Anything, mock.Anything).Return(unit.InteriorCore{}, errors.New("err")).Once()
		data, err := facBusiness.UpdateInteriorPhoto(context.Background(), unit.InteriorCore{})
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestGetInteriorPhoto(t *testing.T) {
	t.Run("GetInteriorPhoto sukses", func(t *testing.T) {
		facData.On("GetInteriorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(unit.InteriorCore{}, nil).Once()
		data, err := facBusiness.GetInteriorPhoto(context.Background(), 1, 2)
		assert.NotEqual(t, data, err)
	})
	t.Run("GetInteriorPhoto gagal", func(t *testing.T) {
		facData.On("GetInteriorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(unit.InteriorCore{}, errors.New("err")).Once()
		data, err := facBusiness.GetInteriorPhoto(context.Background(), 1, 2)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestGetAllInteriorPhoto(t *testing.T) {
	t.Run("GetAllInteriorPhoto sukses", func(t *testing.T) {
		facData.On("GetAllInteriorPhoto", mock.Anything, mock.AnythingOfType("uint")).Return([]unit.InteriorCore{}, nil).Once()
		data, err := facBusiness.GetAllInteriorPhoto(context.Background(), 1)
		assert.NotEqual(t, data, err)
	})
	t.Run("GetAllInteriorPhoto gagal", func(t *testing.T) {
		facData.On("GetAllInteriorPhoto", mock.Anything, mock.AnythingOfType("uint")).Return([]unit.InteriorCore{}, errors.New("err")).Once()
		data, err := facBusiness.GetAllInteriorPhoto(context.Background(), 1)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestGetAllUnitFacility(t *testing.T) {
	t.Run("GetAllUnitFacility sukses", func(t *testing.T) {
		facData.On("GetAllUnitFacility", mock.Anything, mock.AnythingOfType("uint")).Return(unit.Core{}, nil).Once()
		data, err := facBusiness.GetAllUnitFacility(context.Background(), 1)
		assert.NotEqual(t, data, err)
	})
	t.Run("GetAllUnitFacility gagal", func(t *testing.T) {
		facData.On("GetAllUnitFacility", mock.Anything, mock.AnythingOfType("uint")).Return(unit.Core{}, errors.New("err")).Once()
		data, err := facBusiness.GetAllUnitFacility(context.Background(), 1)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestGetUnitFacility(t *testing.T) {
	t.Run("GetUnitFacility sukses", func(t *testing.T) {
		facData.On("GetUnitFacility", mock.Anything, mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(unit.Facility{}, nil).Once()
		data, err := facBusiness.GetUnitFacility(context.Background(), 1, 2)
		assert.NotEqual(t, data, err)
	})
	t.Run("GetUnitFacility gagal", func(t *testing.T) {
		facData.On("GetUnitFacility", mock.Anything, mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(unit.Facility{}, errors.New("err")).Once()
		data, err := facBusiness.GetUnitFacility(context.Background(), 1, 2)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestDeleteUnitFacility(t *testing.T) {
	t.Run("GetUnitFacility sukses", func(t *testing.T) {
		facData.On("GetUnitFacility", mock.Anything, mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(unit.Facility{}, nil).Once()
		facData.On("DeleteUnitFacility", mock.Anything, mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(unit.Core{}, nil).Once()
		data, err := facBusiness.DeleteUnitFacility(context.Background(), 1, 2)
		assert.NotEqual(t, data, err)
	})
	t.Run("GetUnitFacility gagal", func(t *testing.T) {
		facData.On("GetUnitFacility", mock.Anything, mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(unit.Facility{}, errors.New("err")).Once()
		facData.On("DeleteUnitFacility", mock.Anything, mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(unit.Core{}, errors.New("err")).Once()
		data, err := facBusiness.DeleteUnitFacility(context.Background(), 1, 2)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestCreateInteriorPhoto(t *testing.T) {
	t.Run("CreateInteriorPhoto sukses", func(t *testing.T) {
		facData.On("CreateInteriorPhoto", mock.Anything, mock.Anything).Return(unit.InteriorCore{}, nil).Once()
		data, err := facBusiness.CreateInteriorPhoto(context.Background(), unit.InteriorCore{})
		assert.NotEqual(t, data, err)
	})
	t.Run("CreateInteriorPhoto gagal", func(t *testing.T) {
		facData.On("CreateInteriorPhoto", mock.Anything, mock.Anything).Return(unit.InteriorCore{}, errors.New("err")).Once()
		data, err := facBusiness.CreateInteriorPhoto(context.Background(), unit.InteriorCore{})
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestUpdateUnit(t *testing.T) {
	t.Run("UpdateUnit sukses", func(t *testing.T) {
		facData.On("GetUnitById", mock.Anything, mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(facValue, nil).Once()
		facData.On("UpdateUnit", mock.Anything, mock.Anything).Return(facValue, nil).Once()
		data, err := facBusiness.UpdateUnit(context.Background(), facValue2)
		assert.Equal(t, data, facValue)
		assert.Nil(t, err)
	})
	t.Run("UpdateReview gagal", func(t *testing.T) {
		facData.On("GetUnitById", mock.Anything, mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(facValue, nil).Once()
		facData.On("UpdateUnit", mock.Anything, mock.Anything).Return(facValue, errors.New("err")).Once()
		data, err := facBusiness.UpdateUnit(context.Background(), facValue2)
		assert.NotEqual(t, data, facValue)
		assert.Error(t, err)
	})
}

func TestGetUnitById(t *testing.T) {
	t.Run("GetUnitById sukses", func(t *testing.T) {
		facData.On("GetUnitById", mock.Anything, mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(facValue, nil).Once()
		data, _ := facBusiness.GetUnitById(context.Background(), 1, 1)
		assert.Equal(t, data, facValue)
	})
	t.Run("GetUnitById gagal", func(t *testing.T) {
		facData.On("GetUnitById", mock.Anything, mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(facValue, errors.New("err")).Once()
		data, err := facBusiness.GetUnitById(context.Background(), 3, 1)
		assert.NotEqual(t, data, facValue)
		assert.Error(t, err)
	})
}

func TestGetAllUnit(t *testing.T) {
	t.Run("GetAllUnit sukses", func(t *testing.T) {
		facData.On("GetAllUnit", mock.Anything, mock.AnythingOfType("uint")).Return(facList, nil).Once()
		data, err := facBusiness.GetAllUnit(context.Background(), 1)
		assert.Equal(t, data, facList)
		assert.Nil(t, err)
	})
	t.Run("GetAllUnit gagal", func(t *testing.T) {
		facData.On("GetAllUnit", mock.Anything, mock.AnythingOfType("uint")).Return([]unit.Core{}, errors.New("err")).Once()
		data, err := facBusiness.GetAllUnit(context.Background(), 4)
		assert.Equal(t, data, []unit.Core{})
		assert.Error(t, err)
	})
}
