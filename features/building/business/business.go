package business

import (
	"context"
	"ofspace-be/features/building"
	"ofspace-be/features/facility"
	//"ofspace-be/features/unit"
	"ofspace-be/features/users"
	"time"
)

type buildingBusiness struct {
	buildingData building.Data
	userBusiness users.Business
	//unitBusiness     unit.Business
	facilityBusiness facility.Business
	contextTimeout   time.Duration
}

func NewBuildingBusiness(
	buildingData building.Data,
	timeout time.Duration,
	userBusiness users.Business,
	//unitBusiness unit.Business,
	facilityBusiness facility.Business,
) building.Business {
	return &buildingBusiness{
		buildingData:     buildingData,
		userBusiness:     userBusiness,
		facilityBusiness: facilityBusiness,
		//unitBusiness:     unitBusiness,
		contextTimeout: timeout}
}

func (bb *buildingBusiness) CreateBuilding(c context.Context, data building.Core) (building.Core, error) {
	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
	defer error1()
	_, err := bb.userBusiness.GetUserByID(ctx, data.UserId)
	if err != nil {
		return building.Core{}, err
	}
	data.UpdatedAt = time.Now()
	build, err2 := bb.buildingData.CreateBuilding(ctx, data)
	if err2 != nil {
		return building.Core{}, err2
	}
	return build, nil

}
func (bb *buildingBusiness) GetAllBuilding(c context.Context, complexId uint) ([]building.Core, error) {
	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
	defer error1()
	data, err := bb.buildingData.GetAllBuilding(ctx, complexId)
	if err != nil {
		return []building.Core{}, err
	}
	return data, nil
}
func (bb *buildingBusiness) GetAllVerifiedBuilding(c context.Context, complexId uint, buildingStatus string) ([]building.Core, error) {
	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
	defer error1()
	data, err := bb.buildingData.GetAllVerifiedBuilding(ctx, complexId, buildingStatus)
	if err != nil {
		return []building.Core{}, err
	}
	return data, nil
}
func (bb *buildingBusiness) SearchBuildingByName(c context.Context, name string, status string) ([]building.Core, error) {
	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
	defer error1()
	status = "verified"
	build, err := bb.buildingData.SearchBuildingByName(ctx, name, status)
	if err != nil {
		return []building.Core{}, err
	}
	return build, nil
}
func (bb *buildingBusiness) GetBuildingById(c context.Context, id uint, status string) (building.Core, error) {
	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
	defer error1()
	status = "verified"
	build, err := bb.buildingData.GetBuildingById(ctx, id, status)
	if err != nil {
		return building.Core{}, err
	}
	return build, nil

}
func (bb *buildingBusiness) UpdateBuilding(c context.Context, data building.Core) (building.Core, error) {
	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
	defer error1()
	_, err := bb.buildingData.GetBuildingById(ctx, data.Id, data.BuildingStatus)
	if err != nil {
		return building.Core{}, err
	}
	data.UpdatedAt = time.Now()
	up, err2 := bb.buildingData.UpdateBuilding(ctx, data)
	if err2 != nil {
		return building.Core{}, err
	}
	return up, nil
}

//func (bb *buildingBusiness) RequestBuilding(c context.Context, id uint, name string) (building.Core, error) {
//	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
//	defer error1()
//
//}

//exterior photos
func (bb *buildingBusiness) CreateExteriorPhoto(c context.Context, buildingId uint, data building.ExteriorCore) (building.ExteriorCore, error) {
	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
	defer error1()
	data.UpdatedAt = time.Now()
	photo, err := bb.buildingData.CreateExteriorPhoto(ctx, buildingId, data)
	if err != nil {
		return building.ExteriorCore{}, err
	}
	return photo, nil

}
func (bb *buildingBusiness) UpdateExteriorPhoto(c context.Context, data building.ExteriorCore) (building.ExteriorCore, error) {
	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
	defer error1()
	_, err := bb.buildingData.GetExteriorPhoto(ctx, data.BuildingId, data.Id)
	if err != nil {
		return building.ExteriorCore{}, err
	}
	data.UpdatedAt = time.Now()
	up, err2 := bb.buildingData.UpdateExteriorPhoto(ctx, data)
	if err2 != nil {
		return building.ExteriorCore{}, err
	}
	return up, nil
}
func (bb *buildingBusiness) GetExteriorPhoto(c context.Context, BuildingId uint, photoId uint) (building.ExteriorCore, error) {
	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
	defer error1()
	photo, err := bb.buildingData.GetExteriorPhoto(ctx, BuildingId, photoId)
	if err != nil {
		return building.ExteriorCore{}, err
	}
	return photo, nil
}

func (bb *buildingBusiness) GetAllExteriorPhoto(c context.Context, BuildingId uint) ([]building.ExteriorCore, error) {
	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
	defer error1()
	photo, err := bb.buildingData.GetAllExteriorPhoto(ctx, BuildingId)
	if err != nil {
		return []building.ExteriorCore{}, err
	}
	return photo, nil
}

func (bb *buildingBusiness) DeleteExteriorPhoto(c context.Context, BuildingId uint, photoId uint) (building.ExteriorCore, error) {
	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
	defer error1()
	_, err := bb.buildingData.GetExteriorPhoto(ctx, BuildingId, photoId)
	if err != nil {
		return building.ExteriorCore{}, err
	}

	del, err2 := bb.buildingData.DeleteExteriorPhoto(ctx, BuildingId, photoId)
	if err2 != nil {
		return building.ExteriorCore{}, err2
	}

	return del, nil
}

//floor photo
func (bb *buildingBusiness) CreateFloorPhoto(c context.Context, buildingId uint, data building.FloorCore) (building.FloorCore, error) {
	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
	defer error1()
	data.UpdatedAt = time.Now()
	photo, err := bb.buildingData.CreateFloorPhoto(ctx, buildingId, data)
	if err != nil {
		return building.FloorCore{}, err
	}
	return photo, nil

}
func (bb *buildingBusiness) UpdateFloorPhoto(c context.Context, data building.FloorCore) (building.FloorCore, error) {
	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
	defer error1()
	_, err := bb.buildingData.GetFloorPhoto(ctx, data.BuildingId, data.Id)
	if err != nil {
		return building.FloorCore{}, err
	}
	data.UpdatedAt = time.Now()
	up, err2 := bb.buildingData.UpdateFloorPhoto(ctx, data)
	if err2 != nil {
		return building.FloorCore{}, err2
	}
	return up, nil
}
func (bb *buildingBusiness) GetFloorPhoto(c context.Context, BuildingId uint, photoId uint) (building.FloorCore, error) {
	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
	defer error1()
	photo, err := bb.buildingData.GetFloorPhoto(ctx, BuildingId, photoId)
	if err != nil {
		return building.FloorCore{}, err
	}
	return photo, nil
}

func (bb *buildingBusiness) GetAllFloorPhoto(c context.Context, BuildingId uint) ([]building.FloorCore, error) {
	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
	defer error1()
	photo, err := bb.buildingData.GetAllFloorPhoto(ctx, BuildingId)
	if err != nil {
		return []building.FloorCore{}, err
	}
	return photo, nil
}

func (bb *buildingBusiness) DeleteFloorPhoto(c context.Context, BuildingId uint, photoId uint) (building.FloorCore, error) {
	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
	defer error1()
	thisPhoto, err := bb.buildingData.GetFloorPhoto(ctx, BuildingId, photoId)
	if err != nil {
		return building.FloorCore{}, err
	}

	del, err2 := bb.buildingData.DeleteFloorPhoto(ctx, BuildingId, thisPhoto.Id)
	if err2 != nil {
		return building.FloorCore{}, err2
	}

	return del, nil
}

//manage facility
func (bb *buildingBusiness) AddFacilityToBuilding(c context.Context, buildingId uint, facilityId uint) (building.Core, error) {
	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
	defer error1()
	thisFacility, err := bb.facilityBusiness.GetFacility(ctx, facilityId)
	if err != nil {
		return building.Core{}, err
	}
	//data.UpdatedAt = time.Now()
	build, err2 := bb.buildingData.AddFacilityToBuilding(ctx, buildingId, thisFacility.Id)
	if err2 != nil {
		return building.Core{}, err2
	}
	return build, nil

}
func (bb *buildingBusiness) GetAllBuildingFacility(c context.Context, buildingId uint) ([]building.Facility, error) {
	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
	defer error1()
	buildFacc, err := bb.buildingData.GetAllBuildingFacility(ctx, buildingId)
	if err != nil {
		return []building.Facility{}, err
	}
	return buildFacc, nil
}
func (bb *buildingBusiness) GetBuildingFacility(c context.Context, buildingId uint, facilityId uint) (building.Facility, error) {
	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
	defer error1()
	buildFacc, err := bb.buildingData.GetBuildingFacility(ctx, buildingId, facilityId)
	if err != nil {
		return building.Facility{}, err
	}
	return buildFacc, nil
}
func (bb *buildingBusiness) DeleteFacility(c context.Context, buildingId uint, facilityId uint) (building.Facility, error) {
	ctx, error1 := context.WithTimeout(c, bb.contextTimeout)
	defer error1()
	thisFacility, err := bb.buildingData.GetBuildingFacility(ctx, buildingId, facilityId)
	if err != nil {
		return building.Facility{}, err
	}

	del, err2 := bb.buildingData.DeleteFacility(ctx, buildingId, thisFacility.Id)
	if err2 != nil {
		return building.Facility{}, err2
	}

	return del, nil
}
