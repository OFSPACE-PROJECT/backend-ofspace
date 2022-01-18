package business

import (
	"context"
	facility2 "ofspace-be/features/facility"
	"time"
)

type facilityBusiness struct {
	facilityData   facility2.Data
	contextTimeout time.Duration
}

func NewFacilityBusiness(facilityData facility2.Data, timeout time.Duration) facility2.Business {
	return &facilityBusiness{facilityData: facilityData, contextTimeout: timeout}
}

func (cb *facilityBusiness) CreateFacility(c context.Context, data facility2.Facility) (facility2.Facility, error) {
	ctx, error1 := context.WithTimeout(c, cb.contextTimeout)
	defer error1()
	data.UpdatedAt = time.Now()
	facility1, err := cb.facilityData.CreateFacility(ctx, data)
	if err != nil {
		return facility2.Facility{}, err
	}
	return facility1, nil
}

func (cb *facilityBusiness) GetFacility(c context.Context, id uint) (facility2.Facility, error) {
	ctx, error1 := context.WithTimeout(c, cb.contextTimeout)
	defer error1()
	facility1, err := cb.facilityData.GetFacility(ctx, id)
	if err != nil {
		return facility2.Facility{}, err
	}
	return facility1, nil
}

func (cb *facilityBusiness) SearchFacility(c context.Context, name string) ([]facility2.Facility, error) {
	ctx, error1 := context.WithTimeout(c, cb.contextTimeout)
	defer error1()
	facility1, err := cb.facilityData.SearchFacility(ctx, name)
	if err != nil {
		return []facility2.Facility{}, err
	}
	return facility1, nil
}

func (cb *facilityBusiness) UpdateFacility(c context.Context, data facility2.Facility) (facility2.Facility, error) {
	ctx, error1 := context.WithTimeout(c, cb.contextTimeout)
	defer error1()
	_, err := cb.facilityData.GetFacility(ctx, data.Id)
	if err != nil {
		return facility2.Facility{}, err
	}
	data.UpdatedAt = time.Now()
	up, err := cb.facilityData.UpdateFacility(ctx, data)
	if err != nil {
		return facility2.Facility{}, err
	}
	return up, nil
}

func (cb *facilityBusiness) DeleteFacility(c context.Context, id uint) (facility2.Facility, error) {

	ctx, error1 := context.WithTimeout(c, cb.contextTimeout)
	defer error1()

	_, err := cb.facilityData.GetFacility(ctx, id)
	if err != nil {
		return facility2.Facility{}, err
	}

	del, err2 := cb.facilityData.DeleteFacility(ctx, id)
	if err2 != nil {
		return facility2.Facility{}, err
	}

	return del, nil
}
