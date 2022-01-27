package business

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"ofspace-be/features/building"
	mocks "ofspace-be/features/building/mock"
	"ofspace-be/features/facility"
	"ofspace-be/features/users"
	"os"
	"testing"
	"time"
)

var (
	facData          mocks.Data
	facValue         building.Core
	facValue2        building.Core
	usBusiness       users.Business
	facBusiness      building.Business
	facilityBusiness facility.Business

	facList []building.Core
)

func TestMain(m *testing.M) {
	facBusiness = NewBuildingBusiness(&facData, time.Hour*1, usBusiness, facilityBusiness)
	facValue = building.Core{
		Id:                 1,
		UserId:             1,
		ComplexId:          0,
		Units:              nil,
		Reviews:            nil,
		Name:               "",
		Description:        "",
		FloorCount:         "",
		ImageURL:           "",
		OfficeHours:        "",
		BuildingSize:       "",
		AverageFloorSize:   "",
		YearConstructed:    "",
		Lifts:              "",
		Parking:            "",
		Toilets:            "",
		BuildingStatus:     "",
		BuildingFacilities: nil,
		ExteriorPhotos:     nil,
		FloorPhotos:        nil,
		CreatedAt:          time.Time{},
		UpdatedAt:          time.Time{},
	}
	facValue2 = building.Core{
		Id:                 1,
		UserId:             1,
		ComplexId:          0,
		Units:              nil,
		Reviews:            nil,
		Name:               "",
		Description:        "",
		FloorCount:         "",
		ImageURL:           "",
		OfficeHours:        "",
		BuildingSize:       "",
		AverageFloorSize:   "",
		YearConstructed:    "",
		Lifts:              "",
		Parking:            "",
		Toilets:            "",
		BuildingStatus:     "",
		BuildingFacilities: nil,
		ExteriorPhotos:     nil,
		FloorPhotos:        nil,
		CreatedAt:          time.Time{},
		UpdatedAt:          time.Time{},
	}
	facList = []building.Core{
		{
			Id:                 1,
			UserId:             1,
			ComplexId:          0,
			Units:              nil,
			Reviews:            nil,
			Name:               "",
			Description:        "",
			FloorCount:         "",
			ImageURL:           "",
			OfficeHours:        "",
			BuildingSize:       "",
			AverageFloorSize:   "",
			YearConstructed:    "",
			Lifts:              "",
			Parking:            "",
			Toilets:            "",
			BuildingStatus:     "",
			BuildingFacilities: nil,
			ExteriorPhotos:     nil,
			FloorPhotos:        nil,
			CreatedAt:          time.Time{},
			UpdatedAt:          time.Time{},
		},
	}
	os.Exit(m.Run())
}

func TestCreateBuilding(t *testing.T) {
	t.Run("CreateBuilding sukses", func(t *testing.T) {
		facData.On("CreateBuilding", mock.Anything, mock.Anything).Return(facValue, nil).Once()
		data, err := facBusiness.CreateBuilding(context.Background(), facValue)
		assert.NotEqual(t, data, err)
	})
	t.Run("CreateBuilding gagal", func(t *testing.T) {
		facData.On("CreateBuilding", mock.Anything, mock.Anything).Return(facValue, errors.New("err")).Once()
		data, err := facBusiness.CreateBuilding(context.Background(), facValue)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestDeleteExteriorPhoto(t *testing.T) {
	t.Run("DeleteExteriorPhoto sukses", func(t *testing.T) {
		facData.On("GetExteriorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(building.ExteriorCore{}, nil).Once()
		facData.On("DeleteExteriorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(building.ExteriorCore{}, nil).Once()
		data, err := facBusiness.DeleteExteriorPhoto(context.Background(), 1, 2)
		assert.NotEqual(t, data, err)
	})
	t.Run("DeleteExteriorPhoto gagal", func(t *testing.T) {
		facData.On("GetExteriorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(building.ExteriorCore{}, errors.New("err")).Once()
		facData.On("DeleteInteriorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(building.ExteriorCore{}, errors.New("err")).Once()
		data, err := facBusiness.DeleteExteriorPhoto(context.Background(), 1, 2)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestDeleteFloorPhoto(t *testing.T) {
	t.Run("DeleteFloorPhoto sukses", func(t *testing.T) {
		facData.On("GetFloorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(building.FloorCore{}, nil).Once()
		facData.On("DeleteFloorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(building.FloorCore{}, nil).Once()
		data, err := facBusiness.DeleteFloorPhoto(context.Background(), 1, 2)
		assert.NotEqual(t, data, err)
	})
	t.Run("DeleteFloorPhoto gagal", func(t *testing.T) {
		facData.On("GetFloorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(building.FloorCore{}, errors.New("err")).Once()
		facData.On("DeleteFloorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(building.FloorCore{}, errors.New("err")).Once()
		data, err := facBusiness.DeleteFloorPhoto(context.Background(), 1, 2)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestGetExteriorPhoto(t *testing.T) {
	t.Run("GetExteriorPhoto sukses", func(t *testing.T) {
		facData.On("GetExteriorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(building.ExteriorCore{}, nil).Once()
		data, err := facBusiness.GetExteriorPhoto(context.Background(), 1, 2)
		assert.NotEqual(t, data, err)
	})
	t.Run("GetExteriorPhoto gagal", func(t *testing.T) {
		facData.On("GetExteriorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(building.ExteriorCore{}, errors.New("err")).Once()
		data, err := facBusiness.GetExteriorPhoto(context.Background(), 1, 2)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestGetFloorPhoto(t *testing.T) {
	t.Run("GetFloorPhoto sukses", func(t *testing.T) {
		facData.On("GetFloorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(building.FloorCore{}, nil).Once()
		data, err := facBusiness.GetFloorPhoto(context.Background(), 1, 2)
		assert.NotEqual(t, data, err)
	})
	t.Run("GetFloorPhoto gagal", func(t *testing.T) {
		facData.On("GetFloorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(building.FloorCore{}, errors.New("err")).Once()
		data, err := facBusiness.GetFloorPhoto(context.Background(), 1, 2)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}
func TestGetAllBuilding(t *testing.T) {
	t.Run("GetAllBuilding sukses", func(t *testing.T) {
		facData.On("GetAllBuilding", mock.Anything, mock.AnythingOfType("uint")).Return([]building.Core{}, nil).Once()
		data, err := facBusiness.GetAllBuilding(context.Background(), 1)
		assert.NotEqual(t, data, err)
	})
	t.Run("GetAllBuilding gagal", func(t *testing.T) {
		facData.On("GetAllBuilding", mock.Anything, mock.AnythingOfType("uint")).Return([]building.Core{}, errors.New("err")).Once()
		data, err := facBusiness.GetAllBuilding(context.Background(), 1)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestSearchBuildingByName(t *testing.T) {
	t.Run("SearchBuildingByName sukses", func(t *testing.T) {
		facData.On("SearchBuildingByName", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]building.Core{}, nil).Once()
		data, err := facBusiness.SearchBuildingByName(context.Background(), "a", "a")
		assert.NotEqual(t, data, err)
	})
	t.Run("SearchBuildingByName gagal", func(t *testing.T) {
		facData.On("SearchBuildingByName", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]building.Core{}, errors.New("err")).Once()
		data, err := facBusiness.SearchBuildingByName(context.Background(), "a", "a")
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}
func TestGetAllVerifiedBuilding(t *testing.T) {
	t.Run("GetAllVerifiedBuilding sukses", func(t *testing.T) {
		facData.On("GetAllVerifiedBuilding", mock.Anything, mock.AnythingOfType("uint"), mock.AnythingOfType("string")).Return([]building.Core{}, nil).Once()
		data, err := facBusiness.GetAllVerifiedBuilding(context.Background(), 1, "verified")
		assert.NotEqual(t, data, err)
	})
	t.Run("GetAllVerifiedBuilding gagal", func(t *testing.T) {
		facData.On("GetAllVerifiedBuilding", mock.Anything, mock.AnythingOfType("uint"), mock.AnythingOfType("string")).Return([]building.Core{}, errors.New("err")).Once()
		data, err := facBusiness.GetAllVerifiedBuilding(context.Background(), 1, "verified")
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestGetAllExteriorPhoto(t *testing.T) {
	t.Run("GetAllExteriorPhoto sukses", func(t *testing.T) {
		facData.On("GetAllExteriorPhoto", mock.Anything, mock.AnythingOfType("uint")).Return([]building.ExteriorCore{}, nil).Once()
		data, err := facBusiness.GetAllExteriorPhoto(context.Background(), 1)
		assert.NotEqual(t, data, err)
	})
	t.Run("GetAllExteriorPhoto gagal", func(t *testing.T) {
		facData.On("GetAllExteriorPhoto", mock.Anything, mock.AnythingOfType("uint")).Return([]building.ExteriorCore{}, errors.New("err")).Once()
		data, err := facBusiness.GetAllExteriorPhoto(context.Background(), 1)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestGetAllFloorPhoto(t *testing.T) {
	t.Run("GetAllFloorPhoto sukses", func(t *testing.T) {
		facData.On("GetAllFloorPhoto", mock.Anything, mock.AnythingOfType("uint")).Return([]building.FloorCore{}, nil).Once()
		data, err := facBusiness.GetAllFloorPhoto(context.Background(), 1)
		assert.NotEqual(t, data, err)
	})
	t.Run("GetAllFloorPhoto gagal", func(t *testing.T) {
		facData.On("GetAllFloorPhoto", mock.Anything, mock.AnythingOfType("uint")).Return([]building.FloorCore{}, errors.New("err")).Once()
		data, err := facBusiness.GetAllFloorPhoto(context.Background(), 1)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestGetAllBuildingFacility(t *testing.T) {
	t.Run("GetAllBuildingFacility sukses", func(t *testing.T) {
		facData.On("GetAllBuildingFacility", mock.Anything, mock.AnythingOfType("uint")).Return(building.Core{}, nil).Once()
		data, err := facBusiness.GetAllBuildingFacility(context.Background(), 1)
		assert.NotEqual(t, data, err)
	})
	t.Run("GetAllBuildingFacility gagal", func(t *testing.T) {
		facData.On("GetAllBuildingFacility", mock.Anything, mock.AnythingOfType("uint")).Return(building.Core{}, errors.New("err")).Once()
		data, err := facBusiness.GetAllBuildingFacility(context.Background(), 1)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestGetBuildingFacility(t *testing.T) {
	t.Run("GetBuildingFacility sukses", func(t *testing.T) {
		facData.On("GetBuildingFacility", mock.Anything, mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(building.Facility{}, nil).Once()
		data, err := facBusiness.GetBuildingFacility(context.Background(), 1, 2)
		assert.NotEqual(t, data, err)
	})
	t.Run("GetBuildingFacility gagal", func(t *testing.T) {
		facData.On("GetBuildingFacility", mock.Anything, mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(building.Facility{}, errors.New("err")).Once()
		data, err := facBusiness.GetBuildingFacility(context.Background(), 1, 2)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestDeleteBuildingFacility(t *testing.T) {
	t.Run("DeleteBuildingFacility sukses", func(t *testing.T) {
		facData.On("GetBuildingFacility", mock.Anything, mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(building.Facility{}, nil).Once()
		facData.On("DeleteFacility", mock.Anything, mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(building.Core{}, nil).Once()
		data, err := facBusiness.DeleteFacility(context.Background(), 1, 2)
		assert.NotEqual(t, data, err)
	})
	t.Run("DeleteFacility gagal", func(t *testing.T) {
		facData.On("GetBuildingFacility", mock.Anything, mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(building.Facility{}, errors.New("err")).Once()
		facData.On("DeleteFacility", mock.Anything, mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(building.Core{}, errors.New("err")).Once()
		data, err := facBusiness.DeleteFacility(context.Background(), 1, 2)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestCreateExteriorPhoto(t *testing.T) {
	t.Run("CreateExteriorPhoto sukses", func(t *testing.T) {
		facData.On("CreateExteriorPhoto", mock.Anything, mock.Anything).Return(building.ExteriorCore{}, nil).Once()
		data, err := facBusiness.CreateExteriorPhoto(context.Background(), building.ExteriorCore{})
		assert.NotEqual(t, data, err)
	})
	t.Run("CreateExteriorPhoto gagal", func(t *testing.T) {
		facData.On("CreateExteriorPhoto", mock.Anything, mock.Anything).Return(building.ExteriorCore{}, errors.New("err")).Once()
		data, err := facBusiness.CreateExteriorPhoto(context.Background(), building.ExteriorCore{})
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}
func TestCreateFloorPhoto(t *testing.T) {
	t.Run("CreateFloorPhoto sukses", func(t *testing.T) {
		facData.On("CreateFloorPhoto", mock.Anything, mock.Anything).Return(building.FloorCore{}, nil).Once()
		data, err := facBusiness.CreateFloorPhoto(context.Background(), building.FloorCore{})
		assert.NotEqual(t, data, err)
	})
	t.Run("CreateFloorPhoto gagal", func(t *testing.T) {
		facData.On("CreateFloorPhoto", mock.Anything, mock.Anything).Return(building.FloorCore{}, errors.New("err")).Once()
		data, err := facBusiness.CreateFloorPhoto(context.Background(), building.FloorCore{})
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestUpdateExteriorPhoto(t *testing.T) {
	t.Run("UpdateExteriorPhoto sukses", func(t *testing.T) {
		facData.On("GetExteriorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(building.ExteriorCore{}, nil).Once()
		facData.On("UpdateExteriorPhoto", mock.Anything, mock.Anything).Return(building.ExteriorCore{}, nil).Once()
		data, err := facBusiness.UpdateExteriorPhoto(context.Background(), building.ExteriorCore{})
		assert.NotEqual(t, data, err)
	})
	t.Run("UpdateExteriorPhoto gagal", func(t *testing.T) {
		facData.On("GetExteriorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(building.ExteriorCore{}, nil).Once()
		facData.On("UpdateExteriorPhoto", mock.Anything, mock.Anything).Return(building.ExteriorCore{}, errors.New("err")).Once()
		_, err := facBusiness.UpdateExteriorPhoto(context.Background(), building.ExteriorCore{})
		assert.Nil(t, err)
	})
}

//
//func TestUpdateFloorPhoto(t *testing.T) {
//	t.Run("UpdateFloorPhoto sukses", func(t *testing.T) {
//		facData.On("GetFloorPhoto", mock.Anything, mock.AnythingOfType("uint")).Return(building.FloorCore{}, nil).Once()
//		facData.On("UpdateFloorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(building.FloorCore{}, nil).Once()
//		data, err := facBusiness.UpdateFloorPhoto(context.Background(), building.FloorCore{})
//		assert.NotEqual(t, data, err)
//	})
//	t.Run("UpdateFloorPhoto gagal", func(t *testing.T) {
//		facData.On("GetFloorPhoto", mock.Anything, mock.AnythingOfType("uint")).Return(building.FloorCore{}, nil).Once()
//		facData.On("UpdateFloorPhoto", mock.Anything, mock.Anything, mock.Anything).Return(building.FloorCore{}, errors.New("err")).Once()
//		data, err := facBusiness.UpdateFloorPhoto(context.Background(), building.FloorCore{})
//		assert.NotEqual(t, data, err)
//		assert.Error(t, err)
//	})
//}

func TestUpdateBuilding(t *testing.T) {
	t.Run("UpdateBuilding sukses", func(t *testing.T) {
		facData.On("GetBuildingById", mock.Anything, mock.AnythingOfType("uint")).Return(facValue, nil).Once()
		facData.On("UpdateBuilding", mock.Anything, mock.Anything).Return(facValue, nil).Once()
		data, err := facBusiness.UpdateBuilding(context.Background(), facValue2)
		assert.Equal(t, data, facValue)
		assert.Nil(t, err)
	})
	t.Run("UpdateBuilding gagal", func(t *testing.T) {
		facData.On("GetBuildingById", mock.Anything, mock.AnythingOfType("uint")).Return(facValue, nil).Once()
		facData.On("UpdateBuilding", mock.Anything, mock.Anything).Return(facValue, errors.New("err")).Once()
		data, _ := facBusiness.UpdateBuilding(context.Background(), facValue)
		assert.NotEqual(t, data, facValue)
	})
}

func TestGetBuildingById(t *testing.T) {
	t.Run("GetBuildingById sukses", func(t *testing.T) {
		facData.On("GetBuildingById", mock.Anything, mock.AnythingOfType("uint")).Return(facValue, nil).Once()
		data, _ := facBusiness.GetBuildingById(context.Background(), 1)
		assert.Equal(t, data, facValue)
	})
	t.Run("GetBuildingById gagal", func(t *testing.T) {
		facData.On("GetBuildingById", mock.Anything, mock.AnythingOfType("uint")).Return(facValue, errors.New("err")).Once()
		data, err := facBusiness.GetBuildingById(context.Background(), 1)
		assert.NotEqual(t, data, facValue)
		assert.Error(t, err)
	})
}
