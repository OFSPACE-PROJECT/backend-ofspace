package data

import (
	"context"
	complex2 "ofspace-be/features/complex"

	"gorm.io/gorm"
)

type complexData struct {
	Connect *gorm.DB
}

func NewComplexData(connect *gorm.DB) complex2.Data {
	return &complexData{Connect: connect}
}

func (cd *complexData) CreateComplex(ctx context.Context, core complex2.Core) (complex2.Core, error) {
	complex1 := fromComplexCore(core)
	result := cd.Connect.Create(&complex1)
	if result.Error != nil {
		return complex2.Core{}, result.Error
	}
	return toComplexCore(complex1), nil
}

func (cd *complexData) GetComplex(ctx context.Context, id uint) (complex2.Core, error) {
	var complex1 Complex
	result := cd.Connect.Preload("Buildings").First(&complex1, "id= ?", id)
	if result.Error != nil {
		return complex2.Core{}, result.Error
	}
	return toComplexCore(complex1), nil
}
func (cd *complexData) GetAllComplex(ctx context.Context) ([]complex2.Core, error) {
	var complex1 []Complex
	result := cd.Connect.Preload("Buildings").Find(&complex1)
	if result.Error != nil {
		return []complex2.Core{}, result.Error
	}
	return toSliceComplexCore(complex1), nil
}
func (cd *complexData) SearchComplex(ctx context.Context, name string) ([]complex2.Core, error) {
	var complex1 []Complex
	result := cd.Connect.Preload("Buildings").Where("name LIKE ?", "%"+name+"%").Find(&complex1)
	if result.Error != nil {
		return []complex2.Core{}, result.Error
	}
	return ListToCore(complex1), nil
}
func (cd *complexData) SearchComplexByAddress(ctx context.Context, address string) ([]complex2.Core, error) {
	var complex1 []Complex
	result := cd.Connect.Preload("Buildings").Where("address LIKE ?", "%"+address+"%").Find(&complex1)
	if result.Error != nil {
		return []complex2.Core{}, result.Error
	}
	return ListToCore(complex1), nil
}

func (cd *complexData) UpdateComplex(ctx context.Context, core complex2.Core) (complex2.Core, error) {
	complex1 := fromComplexCore(core)
	result := cd.Connect.Where("id= ?", complex1.Id).Updates(&Complex{
		Name:      complex1.Name,
		Address:   complex1.Address,
		Longitude: complex1.Longitude,
		Latitude:  complex1.Latitude,
	})
	if result.Error != nil {
		return complex2.Core{}, result.Error
	}
	return toComplexCore(complex1), nil
}

func (cd *complexData) RequestComplex(ctx context.Context, id uint, name string) (complex2.Core, error) {
	var complex1 Complex
	result := cd.Connect.Where("id= ? && name= ?", id, name).Create(&complex1)
	if result.Error != nil {
		return complex2.Core{}, result.Error
	}
	return toComplexCore(complex1), nil
}
