package business

import (
	"context"
	accessibility "ofspace-be/features/accessibility"
	"time"
)

type accessibilityBusiness struct {
	accessibilityData accessibility.Data
	contextTimeout    time.Duration
}

func NewAccessibilityBusiness(accessibilityData accessibility.Data, timeout time.Duration) accessibility.Business {
	return &accessibilityBusiness{accessibilityData: accessibilityData, contextTimeout: timeout}
}

func (cb *accessibilityBusiness) CreateAccessibility(c context.Context, data accessibility.Core) (accessibility.Core, error) {
	ctx, error1 := context.WithTimeout(c, cb.contextTimeout)
	defer error1()
	data.UpdatedAt = time.Now()
	accessibility1, err := cb.accessibilityData.CreateAccessibility(ctx, data)
	if err != nil {
		return accessibility.Core{}, err
	}
	return accessibility1, nil
}

func (cb *accessibilityBusiness) GetAccessibility(c context.Context, id uint) (accessibility.Core, error) {
	ctx, error1 := context.WithTimeout(c, cb.contextTimeout)
	defer error1()
	accessibility1, err := cb.accessibilityData.GetAccessibility(ctx, id)
	if err != nil {
		return accessibility.Core{}, err
	}
	return accessibility1, nil
}

func (cb *accessibilityBusiness) SearchAccessibility(c context.Context, name string) ([]accessibility.Core, error) {
	ctx, error1 := context.WithTimeout(c, cb.contextTimeout)
	defer error1()
	accessibility1, err := cb.accessibilityData.SearchAccessibility(ctx, name)
	if err != nil {
		return []accessibility.Core{}, err
	}
	return accessibility1, nil
}
func (cb *accessibilityBusiness) SearchAccessibilityByAddress(c context.Context, address string) ([]accessibility.Core, error) {
	ctx, error1 := context.WithTimeout(c, cb.contextTimeout)
	defer error1()
	accessibility1, err := cb.accessibilityData.SearchAccessibilityByAddress(ctx, address)
	if err != nil {
		return []accessibility.Core{}, err
	}
	return accessibility1, nil
}

func (cb *accessibilityBusiness) UpdateAccessibility(c context.Context, data accessibility.Core) (accessibility.Core, error) {
	ctx, error1 := context.WithTimeout(c, cb.contextTimeout)
	defer error1()
	_, err := cb.accessibilityData.GetAccessibility(ctx, data.Id)
	if err != nil {
		return accessibility.Core{}, err
	}
	data.UpdatedAt = time.Now()
	up, err := cb.accessibilityData.UpdateAccessibility(ctx, data)
	if err != nil {
		return accessibility.Core{}, err
	}
	return up, nil
}

func (cb accessibilityBusiness) RequestAccessibility(c context.Context, id uint, name string) (accessibility.Core, error) {
	ctx, error1 := context.WithTimeout(c, cb.contextTimeout)
	defer error1()
	up, err := cb.accessibilityData.RequestAccessibility(ctx, id, name)
	if err != nil {
		return accessibility.Core{}, err
	}
	return up, nil

}
