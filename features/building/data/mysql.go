package data

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"ofspace-be/features/building"
	facility "ofspace-be/features/facility/data"
)

type BuildingData struct {
	Connect *gorm.DB
}

func NewBuildingData(connect *gorm.DB) building.Data {
	return &BuildingData{
		Connect: connect,
	}
}

func (bd *BuildingData) CreateBuilding(ctx context.Context, core building.Core) (building.Core, error) {
	build := FromBuildingCore(core)
	result := bd.Connect.Create(&build)
	if result.Error != nil {
		return building.Core{}, result.Error
	}
	return toBuildingCore(&build), nil
}
func (bd *BuildingData) GetAllBuilding(ctx context.Context, complexId uint) ([]building.Core, error) {
	var buildings []Building
	result := bd.Connect.Find(&buildings, "complex_id= ?", complexId)
	if result.Error != nil {
		fmt.Println(result.Error)
		return []building.Core{}, result.Error
	}
	return ListBuildingToCore(buildings), nil
}

func (bd *BuildingData) GetAllVerifiedBuilding(ctx context.Context, complexId uint, buildingStatus string) ([]building.Core, error) {
	var buildings []Building
	result := bd.Connect.Where("building_status= 'verified'").Find(&buildings, "complex_id= ?", complexId)
	if result.Error != nil {
		fmt.Println(result.Error)
		return []building.Core{}, result.Error
	}
	return ListBuildingToCore(buildings), nil
}

func (bd *BuildingData) SearchBuildingByName(ctx context.Context, name string, status string) ([]building.Core, error) {
	var buildings []Building
	status = "verified"
	result := bd.Connect.Where("name LIKE ? && building_status= ?", "%"+name+"%", status).Find(&buildings)
	if result.Error != nil {
		fmt.Println(result.Error)
		return []building.Core{}, result.Error
	}
	return ListBuildingToCore(buildings), nil
}

func (bd *BuildingData) GetBuildingById(ctx context.Context, id uint) (building.Core, error) {
	var build Building
	result := bd.Connect.Preload("BuildingFacilities").Preload("ExteriorPhotos").Preload("FloorPhotos").First(&build, "id= ?", id)
	if result.Error != nil {
		fmt.Println(result.Error)
		return building.Core{}, result.Error
	}
	return toBuildingCore(&build), nil
}

func (bd *BuildingData) UpdateBuilding(ctx context.Context, core building.Core) (building.Core, error) {
	build := FromBuildingCore(core)
	result := bd.Connect.Where("id= ?", build.Id).Updates(&Building{BuildingStatus: build.BuildingStatus, Name: build.Name, BuildingSize: build.BuildingSize,
		AverageFloorSize: build.AverageFloorSize, Toilets: build.Toilets, Lifts: build.Lifts, Description: build.Description, YearConstructed: build.YearConstructed, Parking: build.Parking, OfficeHours: build.OfficeHours})
	if result.Error != nil {
		fmt.Println(result.Error)
		return building.Core{}, result.Error
	}
	//result = bd.Connect.Save(&build)
	//if result.Error != nil {
	//	fmt.Println(result.Error)
	//	return building.Core{}, result.Error
	//}
	return toBuildingCore(&build), nil
}

func (bd *BuildingData) CreateExteriorPhoto(ctx context.Context, core building.ExteriorCore) (building.ExteriorCore, error) {
	photo := FromExteriorPhotoCore(core)
	result := bd.Connect.Preload("Building").Where("building_id= ?", photo.BuildingID).Create(&photo)
	if result.Error != nil {
		fmt.Println(result.Error)
		return building.ExteriorCore{}, result.Error
	}
	return toExteriorPhotoCore(&photo), nil
}

func (bd *BuildingData) UpdateExteriorPhoto(ctx context.Context, core building.ExteriorCore) (building.ExteriorCore, error) {
	photo := FromExteriorPhotoCore(core)
	result := bd.Connect.Where("id= ?", photo.Id).Updates(&ExteriorPhoto{
		Id:          photo.Id,
		BuildingID:  photo.BuildingID,
		PhotoURL:    photo.PhotoURL,
		Description: photo.Description,
	})
	if result.Error != nil {
		fmt.Println(result.Error)
		return building.ExteriorCore{}, result.Error
	}
	return toExteriorPhotoCore(&photo), nil
}

func (bd *BuildingData) GetAllExteriorPhoto(ctx context.Context, BuildingId uint) ([]building.ExteriorCore, error) {
	var photos []ExteriorPhoto
	result := bd.Connect.Where("building_id= ?", BuildingId).Find(&photos)
	if result.Error != nil {
		fmt.Println(result.Error)
		return []building.ExteriorCore{}, result.Error
	}
	return ToSliceExteriorPhotoCore(photos), nil
}

func (bd *BuildingData) GetExteriorPhoto(ctx context.Context, buildingId uint, photoId uint) (building.ExteriorCore, error) {
	var photo ExteriorPhoto
	result := bd.Connect.First(&photo, "building_id= ? && id= ?", buildingId, photoId)
	if result.Error != nil {
		fmt.Println(result.Error)
		return building.ExteriorCore{}, result.Error
	}
	return toExteriorPhotoCore(&photo), nil
}

func (bd *BuildingData) DeleteExteriorPhoto(ctx context.Context, buildingId uint, photoId uint) (building.ExteriorCore, error) {
	var photo ExteriorPhoto
	result := bd.Connect.Where("building_id= ?", buildingId).Delete(&photo, "id= ?", photoId)

	if result.Error != nil {
		return building.ExteriorCore{}, result.Error
	}

	return toExteriorPhotoCore(&photo), nil
}

func (bd *BuildingData) CreateFloorPhoto(ctx context.Context, core building.FloorCore) (building.FloorCore, error) {
	photo := FromFloorPhotoCore(core)
	result := bd.Connect.Preload("Building").Where("building_id= ?", photo.BuildingID).Create(&photo)
	if result.Error != nil {
		fmt.Println(result.Error)
		return building.FloorCore{}, result.Error
	}
	return toFloorPhotoCore(&photo), nil
}

func (bd *BuildingData) UpdateFloorPhoto(ctx context.Context, core building.FloorCore) (building.FloorCore, error) {
	photo := FromFloorPhotoCore(core)
	result := bd.Connect.Where("id= ?", photo.Id).Updates(&FloorPhoto{
		Id:          photo.Id,
		BuildingID:  photo.BuildingID,
		PhotoURL:    photo.PhotoURL,
		Description: photo.Description,
	})
	if result.Error != nil {
		fmt.Println(result.Error)
		return building.FloorCore{}, result.Error
	}
	return toFloorPhotoCore(&photo), nil
}

func (bd *BuildingData) GetAllFloorPhoto(ctx context.Context, BuildingId uint) ([]building.FloorCore, error) {
	var photos []FloorPhoto
	result := bd.Connect.Where("building_id= ?", BuildingId).Find(&photos)
	if result.Error != nil {
		fmt.Println(result.Error)
		return []building.FloorCore{}, result.Error
	}
	return ToSliceFloorPhotoCore(photos), nil
}

func (bd *BuildingData) GetFloorPhoto(ctx context.Context, buildingId uint, photoId uint) (building.FloorCore, error) {
	var photo FloorPhoto
	result := bd.Connect.First(&photo, "building_id= ? && id= ?", buildingId, photoId)
	if result.Error != nil {
		fmt.Println(result.Error)
		return building.FloorCore{}, result.Error
	}
	return toFloorPhotoCore(&photo), nil
}

func (bd *BuildingData) DeleteFloorPhoto(ctx context.Context, buildingId uint, photoId uint) (building.FloorCore, error) {
	var photo FloorPhoto
	result := bd.Connect.Where("building_id= ?", buildingId).Delete(&photo, "id= ?", photoId)
	if result.Error != nil {
		return building.FloorCore{}, result.Error
	}

	return toFloorPhotoCore(&photo), nil
}

func (bd *BuildingData) AddFacilityToBuilding(c context.Context, facilityId uint, buildingId uint) (building.Facility, error) {
	newFacility := BuildingFacility{
		FacilityID: facilityId,
		BuildingID: buildingId,
	}
	thisFacility := facility.Facility{Id: facilityId}
	fmt.Println(facilityId, buildingId, "test")
	err := bd.Connect.Preload("Building").Create(&newFacility).Error
	if err != nil {
		return building.Facility{}, err
	}
	return toFacilityCore(&thisFacility), nil
}

func (bd *BuildingData) GetAllBuildingFacility(c context.Context, buildingId uint) (building.Core, error) {
	var buildings Building
	result := bd.Connect.Debug().Preload("BuildingFacilities").Find(&buildings, "id", buildingId)

	if result.Error != nil {
		fmt.Println(result.Error)
		return building.Core{}, result.Error
	}
	return toBuildingCore(&buildings), nil
}

//SELECT buildings.id, buildings.name, facilities.id AS facility_id, facilities.name AS facility_name
//FROM buildings

func (bd *BuildingData) GetBuildingFacility(c context.Context, buildingId uint, facilityId uint) (building.Facility, error) {
	buildings := FromBuildingCore(building.Core{})
	result := bd.Connect.Raw("SELECT buildings.id, facilities.id AS facility_id, facilities.name AS name FROM buildings JOIN building_facilities on (buildings.id=building_facilities.building_id) JOIN facilities on (facilities.id=building_facilities.facility_id) AND building_facilities.building_id= ? AND building_facilities.facility_id= ?", buildingId, facilityId).First(&buildings.BuildingFacilities)
	if result.Error != nil {
		fmt.Println(result.Error)
		return building.Facility{}, result.Error
	}
	return toFacilityCore(&buildings.BuildingFacilities[0]), nil
}
func (bd *BuildingData) DeleteFacility(c context.Context, buildingId uint, facilityId uint) (building.Core, error) {
	builds := FromBuildingCore(building.Core{})
	var facility BuildingFacility
	result := bd.Connect.Debug().Delete(&facility, "building_id= ? && facility_id= ?", buildingId, facilityId)
	if result.Error != nil {
		return building.Core{}, result.Error
	}
	return builds.ToBuildingCore(), nil
}
