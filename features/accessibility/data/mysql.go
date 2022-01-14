package data

import (
	"context"
	"gorm.io/gorm"
	accessibility "ofspace-be/features/accessibility"
)

type accessibilityData struct {
	Connect *gorm.DB
}

func NewAccessibilityData(connect *gorm.DB) accessibility.Data {
	return &accessibilityData{Connect: connect}
}

func (cd *accessibilityData) CreateAccessibility(ctx context.Context, core accessibility.Core) (accessibility.Core, error) {
	accessibility1 := fromAccessibilityCore(core)
	result := cd.Connect.Create(&accessibility1)
	if result.Error != nil {
		return accessibility.Core{}, result.Error
	}
	return toAccessibilityCore(accessibility1), nil
}

func (cd *accessibilityData) GetAccessibility(ctx context.Context, id uint) (accessibility.Core, error) {
	var accessibility1 Accessibility
	result := cd.Connect.First(&accessibility1, "id= ?", id)
	if result.Error != nil {
		return accessibility.Core{}, result.Error
	}
	return toAccessibilityCore(accessibility1), nil
}

func (cd *accessibilityData) SearchAccessibility(ctx context.Context, name string) ([]accessibility.Core, error) {
	var accessibility1 []Accessibility
	result := cd.Connect.Where("name LIKE ?", "%"+name+"%").Find(&accessibility1)
	if result.Error != nil {
		return []accessibility.Core{}, result.Error
	}
	//result2 := cd.Connect.Where("name LIKE ?", "%"+name).Find(&accessibility1)
	//if result2.Error != nil {
	//	return []accessibility.Core{}, result2.Error
	//}
	return ListToCore(accessibility1), nil
}

func (cd *accessibilityData) UpdateAccessibility(ctx context.Context, core accessibility.Core) (accessibility.Core, error) {
	accessibility1 := fromAccessibilityCore(core)
	result := cd.Connect.Where("id= ?", accessibility1.Id).Updates(&Accessibility{
		Name:      accessibility1.Name,
		Longitude: accessibility1.Longitude,
		Latitude:  accessibility1.Latitude,
	})
	if result.Error != nil {
		return accessibility.Core{}, result.Error
	}
	return toAccessibilityCore(accessibility1), nil
}

func (cd *accessibilityData) RequestAccessibility(ctx context.Context, id uint, name string) (accessibility.Core, error) {
	var accessibility1 Accessibility
	result := cd.Connect.Where("id= ? && name= ?", id, name).Create(&accessibility1)
	if result.Error != nil {
		return accessibility.Core{}, result.Error
	}
	return toAccessibilityCore(accessibility1), nil
}
