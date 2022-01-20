package data

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	facility "ofspace-be/features/facility/data"
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

func (bd *UnitData) GetUnitById(ctx context.Context, facilityId uint, buildingId uint) (unit.Core, error) {
	var build Unit
	result := bd.Connect.First(&build, "id= ? && building_id= ?", facilityId, buildingId)
	if result.Error != nil {
		fmt.Println(result.Error)
		return unit.Core{}, result.Error
	}
	return toUnitCore(&build), nil
}
func (bd *UnitData) GetUnitByType(ctx context.Context, buildingId uint, unitType string) (unit.Core, error) {
	var build Unit
	result := bd.Connect.First(&build, "unit_type= ? && building_id= ?", unitType, buildingId)
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
	result := bd.Connect.Where("unit_id= ?", photo.UnitID).Create(&photo)
	if result.Error != nil {
		fmt.Println(result.Error)
		return unit.InteriorCore{}, result.Error
	}
	return toInteriorPhotoCore(&photo), nil
}

func (bd *UnitData) UpdateInteriorPhoto(ctx context.Context, core unit.InteriorCore) (unit.InteriorCore, error) {
	photo := FromInteriorPhotoCore(core)
	result := bd.Connect.Debug().Where("id= ?", photo.Id).Updates(&InteriorPhoto{
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
	result := bd.Connect.Debug().First(&photo, "unit_id= ? && id= ?", unitId, photoId)
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
	thisFacility := facility.Facility{Id: facilityId}
	err := bd.Connect.Create(&newFacility).Error
	if err != nil {
		return unit.Facility{}, err
	}
	return toFacilityCore(&thisFacility), nil
}

func (bd *UnitData) GetAllUnitFacility(c context.Context, unitId uint) (unit.Core, error) {
	var units Unit
	result := bd.Connect.Debug().Preload("UnitFacilities").Find(&units, "id", unitId)

	if result.Error != nil {
		fmt.Println(result.Error)
		return unit.Core{}, result.Error
	}
	return toUnitCore(&units), nil
}
func (bd *UnitData) GetUnitFacility(c context.Context, unitId uint, facilityId uint) (unit.Facility, error) {
	units := FromUnitCore(unit.Core{})
	result := bd.Connect.Debug().Raw("SELECT units.id, facilities.id AS facility_id, facilities.name AS name FROM units JOIN unit_facilities on (units.id=unit_facilities.unit_id) JOIN facilities on (facilities.id=unit_facilities.facility_id) AND unit_facilities.unit_id= ? AND unit_facilities.facility_id= ?", unitId, facilityId).First(&units.UnitFacilities)
	if result.Error != nil {
		fmt.Println(result.Error)
		return unit.Facility{}, result.Error
	}
	return toFacilityCore(&units.UnitFacilities[0]), nil
}
func (bd *UnitData) DeleteUnitFacility(c context.Context, unitId uint, facilityId uint) (unit.Core, error) {
	facs := FromUnitCore(unit.Core{})
	var facility1 UnitFacility
	result := bd.Connect.Debug().Delete(&facility1, "unit_id= ? && facility_id= ?", unitId, facilityId)
	if result.Error != nil {
		return unit.Core{}, result.Error
	}
	return toUnitCore(&facs), nil
}
