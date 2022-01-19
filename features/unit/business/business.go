package business

import (
	"context"
	"fmt"
	"ofspace-be/features/booking"
	"ofspace-be/features/building"
	"ofspace-be/features/facility"
	"ofspace-be/features/unit"
	"ofspace-be/features/users"
	"time"
)

type unitBusiness struct {
	unitData         unit.Data
	buildingBusiness building.Business
	facilityBusiness facility.Business
	userBusiness     users.Business
	bookingBusiness  booking.Business
	contextTimeout   time.Duration
}

func NewUnitBusiness(unitData unit.Data, buildingBusiness building.Business, userBusiness users.Business, facilityBusiness facility.Business, bookingBusiness booking.Business, timeout time.Duration) unit.Business {
	return &unitBusiness{unitData: unitData, contextTimeout: timeout, buildingBusiness: buildingBusiness, userBusiness: userBusiness, facilityBusiness: facilityBusiness, bookingBusiness: bookingBusiness}
}

func (ub *unitBusiness) CreateUnit(c context.Context, data unit.Core) (unit.Core, error) {
	ctx, error1 := context.WithTimeout(c, ub.contextTimeout)
	defer error1()
	data.UpdatedAt = time.Now()
	unit1, err := ub.unitData.CreateUnit(ctx, data)
	if err != nil {
		return unit.Core{}, err
	}
	return unit1, nil
}

func (ub *unitBusiness) GetAllUnit(c context.Context, buildingId uint) ([]unit.Core, error) {
	ctx, error1 := context.WithTimeout(c, ub.contextTimeout)
	defer error1()
	data, err := ub.unitData.GetAllUnit(ctx, buildingId)
	if err != nil {
		return []unit.Core{}, err
	}
	return data, nil
}

func (ub *unitBusiness) GetUnitById(c context.Context, buildingId uint, facilityId uint) (unit.Core, error) {
	ctx, error1 := context.WithTimeout(c, ub.contextTimeout)
	defer error1()

	unit1, err := ub.unitData.GetUnitById(ctx, buildingId, facilityId)
	if err != nil {
		return unit.Core{}, err
	}
	fmt.Println(unit1.TotalUnit)
	return unit1, nil
}
func (ub *unitBusiness) GetUnitByType(c context.Context, buildingId uint, typeUnit string) (unit.Core, error) {
	ctx, error1 := context.WithTimeout(c, ub.contextTimeout)
	defer error1()
	unit1, err := ub.unitData.GetUnitByType(ctx, buildingId, typeUnit)
	if err != nil {
		return unit.Core{}, err
	}
	return unit1, nil
}
func (ub *unitBusiness) UpdateUnit(c context.Context, data unit.Core) (unit.Core, error) {
	ctx, error1 := context.WithTimeout(c, ub.contextTimeout)
	defer error1()
	_, err := ub.unitData.GetUnitById(ctx, data.Id, data.BuildingId)
	if err != nil {
		return unit.Core{}, err
	}
	data.UpdatedAt = time.Now()

	up, err4 := ub.unitData.UpdateUnit(ctx, data)
	if err4 != nil {
		return unit.Core{}, err4
	}
	return up, nil
}

//func (ub *unitBusiness) DeleteUnit(c context.Context, id uint) (unit.Facility, error) {
//
//	ctx, error1 := context.WithTimeout(c, ub.contextTimeout)
//	defer error1()
//
//	_, err := ub.unitData.GetUnitById(ctx, id)
//	if err != nil {
//		return unit.Facility{}, err
//	}
//
//	del, err2 := ub.unitData.DeleteUnit(ctx, id)
//	if err2 != nil {
//		return unit.Facility{}, err
//	}
//
//	return del, nil
//}

func (ub *unitBusiness) CreateInteriorPhoto(c context.Context, data unit.InteriorCore) (unit.InteriorCore, error) {
	ctx, error1 := context.WithTimeout(c, ub.contextTimeout)
	defer error1()
	data.UpdatedAt = time.Now()
	photo, err := ub.unitData.CreateInteriorPhoto(ctx, data)
	if err != nil {
		return unit.InteriorCore{}, err
	}
	return photo, nil

}
func (ub *unitBusiness) UpdateInteriorPhoto(c context.Context, data unit.InteriorCore) (unit.InteriorCore, error) {
	ctx, error1 := context.WithTimeout(c, ub.contextTimeout)
	defer error1()
	_, err := ub.unitData.GetInteriorPhoto(ctx, data.UnitId, data.Id)
	if err != nil {
		return unit.InteriorCore{}, err
	}
	data.UpdatedAt = time.Now()
	up, err2 := ub.unitData.UpdateInteriorPhoto(ctx, data)
	if err2 != nil {
		return unit.InteriorCore{}, err2
	}
	return up, nil
}
func (ub *unitBusiness) GetInteriorPhoto(c context.Context, UnitId uint, photoId uint) (unit.InteriorCore, error) {
	ctx, error1 := context.WithTimeout(c, ub.contextTimeout)
	defer error1()
	photo, err := ub.unitData.GetInteriorPhoto(ctx, UnitId, photoId)
	if err != nil {
		return unit.InteriorCore{}, err
	}
	return photo, nil
}

func (ub *unitBusiness) GetAllInteriorPhoto(c context.Context, UnitId uint) ([]unit.InteriorCore, error) {
	ctx, error1 := context.WithTimeout(c, ub.contextTimeout)
	defer error1()
	photo, err := ub.unitData.GetAllInteriorPhoto(ctx, UnitId)
	if err != nil {
		return []unit.InteriorCore{}, err
	}
	return photo, nil
}

func (ub *unitBusiness) DeleteInteriorPhoto(c context.Context, UnitId uint, photoId uint) (unit.InteriorCore, error) {
	ctx, error1 := context.WithTimeout(c, ub.contextTimeout)
	defer error1()
	thisPhoto, err := ub.unitData.GetInteriorPhoto(ctx, UnitId, photoId)
	if err != nil {
		return unit.InteriorCore{}, err
	}

	del, err2 := ub.unitData.DeleteInteriorPhoto(ctx, UnitId, thisPhoto.Id)
	if err2 != nil {
		return unit.InteriorCore{}, err2
	}

	return del, nil
}

//manage facility
func (ub *unitBusiness) AddFacilityToUnit(c context.Context, unitId uint, facilityId uint) (unit.Facility, error) {
	ctx, error1 := context.WithTimeout(c, ub.contextTimeout)
	defer error1()
	thisUnit, err := ub.facilityBusiness.GetFacility(ctx, facilityId)
	if err != nil {
		return unit.Facility{}, err
	}
	//data.UpdatedAt = time.Now()
	build, err2 := ub.unitData.AddFacilityToUnit(ctx, unitId, thisUnit.Id)
	if err2 != nil {
		return unit.Facility{}, err2
	}
	return build, nil

}
func (ub *unitBusiness) GetAllUnitFacility(c context.Context, unitId uint) (unit.Core, error) {
	ctx, error1 := context.WithTimeout(c, ub.contextTimeout)
	defer error1()
	unitF, err := ub.unitData.GetAllUnitFacility(ctx, unitId)
	if err != nil {
		return unit.Core{}, err
	}
	return unitF, nil
}
func (ub *unitBusiness) GetUnitFacility(c context.Context, unitId uint, facilityId uint) (unit.Facility, error) {
	ctx, error1 := context.WithTimeout(c, ub.contextTimeout)
	defer error1()
	buildFacc, err := ub.unitData.GetUnitFacility(ctx, unitId, facilityId)
	if err != nil {
		return unit.Facility{}, err
	}
	return buildFacc, nil
}
func (ub *unitBusiness) DeleteUnitFacility(c context.Context, unitId uint, facilityId uint) (unit.Core, error) {
	ctx, error1 := context.WithTimeout(c, ub.contextTimeout)
	defer error1()
	thisUnit, err := ub.unitData.GetUnitFacility(ctx, unitId, facilityId)
	if err != nil {
		return unit.Core{}, err
	}

	del, err2 := ub.unitData.DeleteUnitFacility(ctx, unitId, thisUnit.Id)
	if err2 != nil {
		return unit.Core{}, err2
	}

	return del, nil
}
