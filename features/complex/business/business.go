package business

import (
	"context"
	complex2 "ofspace-be/features/complex"
	"time"
)

type complexBusiness struct {
	complexData    complex2.Data
	contextTimeout time.Duration
}

func NewComplexBusiness(complexData complex2.Data, timeout time.Duration) complex2.Business {
	return &complexBusiness{complexData: complexData, contextTimeout: timeout}
}

func (cb *complexBusiness) CreateComplex(c context.Context, data complex2.Core) (complex2.Core, error) {
	ctx, error1 := context.WithTimeout(c, cb.contextTimeout)
	defer error1()
	data.UpdatedAt = time.Now()
	complex1, err := cb.complexData.CreateComplex(ctx, data)
	if err != nil {
		return complex2.Core{}, err
	}
	return complex1, nil
}

func (cb *complexBusiness) GetComplex(c context.Context, id uint) (complex2.Core, error) {
	ctx, error1 := context.WithTimeout(c, cb.contextTimeout)
	defer error1()
	complex1, err := cb.complexData.GetComplex(ctx, id)
	if err != nil {
		return complex2.Core{}, err
	}
	return complex1, nil
}
func (cb *complexBusiness) GetAllComplex(c context.Context) ([]complex2.Core, error) {
	ctx, error1 := context.WithTimeout(c, cb.contextTimeout)
	defer error1()
	complex1, err := cb.complexData.GetAllComplex(ctx)
	if err != nil {
		return []complex2.Core{}, err
	}
	return complex1, nil
}
func (cb *complexBusiness) SearchComplex(c context.Context, name string) ([]complex2.Core, error) {
	ctx, error1 := context.WithTimeout(c, cb.contextTimeout)
	defer error1()
	complex1, err := cb.complexData.SearchComplex(ctx, name)
	if err != nil {
		return []complex2.Core{}, err
	}
	return complex1, nil
}
func (cb *complexBusiness) SearchComplexByAddress(c context.Context, address string) ([]complex2.Core, error) {
	ctx, error1 := context.WithTimeout(c, cb.contextTimeout)
	defer error1()
	complex1, err := cb.complexData.SearchComplex(ctx, address)
	if err != nil {
		return []complex2.Core{}, err
	}
	return complex1, nil
}

func (cb *complexBusiness) UpdateComplex(c context.Context, data complex2.Core) (complex2.Core, error) {
	ctx, error1 := context.WithTimeout(c, cb.contextTimeout)
	defer error1()
	_, err := cb.complexData.GetComplex(ctx, data.Id)
	if err != nil {
		return complex2.Core{}, err
	}
	data.UpdatedAt = time.Now()
	up, err := cb.complexData.UpdateComplex(ctx, data)
	if err != nil {
		return complex2.Core{}, err
	}
	return up, nil
}

//func (cb complexBusiness) RequestComplex(c context.Context, id uint, name string) (complex2.Core, error) {
//	ctx, error1 := context.WithTimeout(c, cb.contextTimeout)
//	defer error1()
//	up, err := cb.complexData.RequestComplex(ctx, id, name)
//	if err != nil {
//		return complex2.Core{}, err
//	}
//	return up, nil
//
//}
