package data

import (
	"context"
	"gorm.io/gorm"
	facility "ofspace-be/features/facility"
)

type facilityData struct {
	Connect *gorm.DB
}

func NewFacilityData(connect *gorm.DB) facility.Data {
	return &facilityData{Connect: connect}
}

func (cd *facilityData) CreateFacility(ctx context.Context, core facility.Facility) (facility.Facility, error) {
	facility1 := fromFacilityCore(core)
	result := cd.Connect.Create(&facility1)
	if result.Error != nil {
		return facility.Facility{}, result.Error
	}
	return toFacilityCore(facility1), nil
}

func (cd *facilityData) GetFacility(ctx context.Context, id uint) (facility.Facility, error) {
	var facility1 Facility
	result := cd.Connect.First(&facility1, "id= ?", id)
	if result.Error != nil {
		return facility.Facility{}, result.Error
	}
	return toFacilityCore(facility1), nil
}

func (cd *facilityData) GetAllFacility(ctx context.Context) ([]facility.Facility, error) {
	var facility1 []Facility
	result := cd.Connect.Find(&facility1)
	if result.Error != nil {
		return []facility.Facility{}, result.Error
	}
	return ListToCore(facility1), nil
}

func (cd *facilityData) SearchFacility(ctx context.Context, name string) ([]facility.Facility, error) {
	var facility1 []Facility
	result := cd.Connect.Where("name LIKE ?", "%"+name+"%").Find(&facility1)
	if result.Error != nil {
		return []facility.Facility{}, result.Error
	}
	return ListToCore(facility1), nil
}

func (cd *facilityData) UpdateFacility(ctx context.Context, core facility.Facility) (facility.Facility, error) {
	facility1 := fromFacilityCore(core)
	result := cd.Connect.Where("id= ?", facility1.Id).Updates(&Facility{
		Name: facility1.Name,
	})
	if result.Error != nil {
		return facility.Facility{}, result.Error
	}
	return toFacilityCore(facility1), nil
}

func (cd *facilityData) DeleteFacility(ctx context.Context, id uint) (facility.Facility, error) {
	var fac Facility
	result := cd.Connect.Delete(&fac, "id= ?", id)

	if result.Error != nil {
		return facility.Facility{}, result.Error
	}

	return toFacilityCore(fac), nil
}
