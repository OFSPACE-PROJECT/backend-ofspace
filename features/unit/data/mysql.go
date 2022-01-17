package data

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"ofspace-be/features/unit"
	"time"
)

type UnitData struct {
	Connect *gorm.DB
}

func NewUnitData(connect *gorm.DB) unit.Data {
	return &UnitData{
		Connect: connect,
	}
}

func (bd *UnitData) CreateUnit(ctx context.Context, core unit.Core) (unit.Core, error) {
	build := FromUnitCore(core)
	result := bd.Connect.Create(&build)
	if result.Error != nil {
		return unit.Core{}, result.Error
	}
	return toUnitCore(&build), nil
}
func (bd *UnitData) GetAllUnit(ctx context.Context, buildingId uint) ([]unit.Core, error) {
	var units []Unit
	result := bd.Connect.Find(&units, "building_id= ?", buildingId)
	if result.Error != nil {
		fmt.Println(result.Error)
		return []unit.Core{}, result.Error
	}
	return ListUnitToCore(units), nil
}

func (bd *UnitData) GetAllVerifiedUnit(ctx context.Context, buildingId uint, unitStatus string) ([]unit.Core, error) {
	var units []Unit
	result := bd.Connect.Where("unit_status= 'verified'").Find(&units, "building_id= ?", buildingId)
	if result.Error != nil {
		fmt.Println(result.Error)
		return []unit.Core{}, result.Error
	}
	return ListUnitToCore(units), nil
}

func (bd *UnitData) SearchUnitByName(ctx context.Context, name string, status string) ([]unit.Core, error) {
	var units []Unit
	status = "verified"
	result := bd.Connect.Where("name LIKE ? && unit_status= ?", "%"+name+"%", status).Find(&units)
	if result.Error != nil {
		fmt.Println(result.Error)
		return []unit.Core{}, result.Error
	}
	return ListUnitToCore(units), nil
}

func (bd *UnitData) GetUnitById(ctx context.Context, id uint) (unit.Core, error) {
	var build Unit
	result := bd.Connect.First(&build, "id= ?", id)
	if result.Error != nil {
		fmt.Println(result.Error)
		return unit.Core{}, result.Error
	}
	return toUnitCore(&build), nil
}
func (bd *UnitData) GetUnitByType(ctx context.Context, unitType string) (unit.Core, error) {
	var build Unit
	result := bd.Connect.First(&build, "unit_type= ?", unitType)
	if result.Error != nil {
		fmt.Println(result.Error)
		return unit.Core{}, result.Error
	}
	return toUnitCore(&build), nil
}
func (bd *UnitData) UpdateUnit(ctx context.Context, core unit.Core) (unit.Core, error) {
	build := FromUnitCore(core)
	result := bd.Connect.Where("id= ?", build.Id).Updates(&Unit{
		Id:             build.Id,
		UserId:         build.UserId,
		BuildingId:     build.BuildingId,
		Description:    build.Description,
		Price:          build.Price,
		TotalUnit:      build.TotalUnit,
		RemainingUnit:  build.RemainingUnit,
		UnitFacilities: build.UnitFacilities,
		InteriorPhotos: build.InteriorPhotos,
		UpdatedAt:      time.Time{},
	})
	if result.Error != nil {
		fmt.Println(result.Error)
		return unit.Core{}, result.Error
	}
	//result = bd.Connect.Save(&build)
	//if result.Error != nil {
	//	fmt.Println(result.Error)
	//	return unit.Core{}, result.Error
	//}
	return toUnitCore(&build), nil
}

func (bd *UnitData) CreateInteriorPhoto(ctx context.Context, core unit.InteriorCore) (unit.InteriorCore, error) {
	photo := FromInteriorPhotoCore(core)
	result := bd.Connect.Preload("Unit").Where("unit_id= ?", photo.UnitID).Create(&photo)
	if result.Error != nil {
		fmt.Println(result.Error)
		return unit.InteriorCore{}, result.Error
	}
	return toInteriorPhotoCore(&photo), nil
}

func (bd *UnitData) UpdateInteriorPhoto(ctx context.Context, core unit.InteriorCore) (unit.InteriorCore, error) {
	photo := FromInteriorPhotoCore(core)
	result := bd.Connect.Where("id= ?", photo.Id).Updates(&InteriorPhoto{
		Id:          photo.Id,
		UnitID:      photo.UnitID,
		PhotoURL:    photo.PhotoURL,
		Description: photo.Description,
	})
	if result.Error != nil {
		fmt.Println(result.Error)
		return unit.InteriorCore{}, result.Error
	}
	return toInteriorPhotoCore(&photo), nil
}

func (bd *UnitData) GetAllInteriorPhoto(ctx context.Context, UnitId uint) ([]unit.InteriorCore, error) {
	var photos []InteriorPhoto
	result := bd.Connect.Where("unit_id= ?", UnitId).Find(&photos)
	if result.Error != nil {
		fmt.Println(result.Error)
		return []unit.InteriorCore{}, result.Error
	}
	return ToSliceInteriorPhotoCore(photos), nil
}

func (bd *UnitData) GetInteriorPhoto(ctx context.Context, unitId uint, photoId uint) (unit.InteriorCore, error) {
	var photo InteriorPhoto
	result := bd.Connect.First(&photo, "unit_id= ? && id= ?", unitId, photoId)
	if result.Error != nil {
		fmt.Println(result.Error)
		return unit.InteriorCore{}, result.Error
	}
	return toInteriorPhotoCore(&photo), nil
}

func (bd *UnitData) DeleteInteriorPhoto(ctx context.Context, unitId uint, photoId uint) (unit.InteriorCore, error) {
	var photo InteriorPhoto
	result := bd.Connect.Where("unit_id= ?", unitId).Delete(&photo, "id= ?", photoId)

	if result.Error != nil {
		return unit.InteriorCore{}, result.Error
	}

	return toInteriorPhotoCore(&photo), nil
}

func (bd *UnitData) AddFacilityToUnit(c context.Context, facilityId uint, unitId uint) (unit.Facility, error) {
	newFacility := UnitFacility{
		FacilityID: facilityId,
		UnitID:     unitId,
	}

	err := bd.Connect.Preload("Unit").Where("facility_id= ? && unit_id= ?", facilityId, unitId).Create(&newFacility).Error
	if err != nil {
		return unit.Facility{}, err
	}
	return toFacilityCore(&newFacility), nil
}

func (bd *UnitData) GetAllUnitFacility(c context.Context, unitId uint) ([]unit.Facility, error) {
	var facilities []UnitFacility
	result := bd.Connect.Find(&facilities, "id= ?", unitId)
	if result.Error != nil {
		fmt.Println(result.Error)
		return []unit.Facility{}, result.Error
	}
	return ToSliceFacilityPhotoCore(facilities), nil
}
func (bd *UnitData) GetUnitFacility(c context.Context, unitId uint, facilityId uint) (unit.Facility, error) {
	var facility UnitFacility
	result := bd.Connect.Where("id= ?", unitId).First(&facility, "id= ?", facilityId)
	if result.Error != nil {
		fmt.Println(result.Error)
		return unit.Facility{}, result.Error
	}
	return toFacilityCore(&facility), nil
}
func (bd *UnitData) DeleteUnitFacility(c context.Context, unitId uint, facilityId uint) (unit.Facility, error) {
	var facility UnitFacility
	result := bd.Connect.Where("unit_id= ?", unitId).Delete(&facility, "id= ?", facilityId)
	if result.Error != nil {
		return unit.Facility{}, result.Error
	}

	return toFacilityCore(&facility), nil
}
