package data

import (
	accessibility "ofspace-be/features/accessibility"
	"time"
)

type Accessibility struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	Longitude string
	Latitude  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toAccessibilityCore(c Accessibility) accessibility.Core {
	return accessibility.Core{
		Id:        c.Id,
		Name:      c.Name,
		Longitude: c.Longitude,
		Latitude:  c.Latitude,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func fromAccessibilityCore(c accessibility.Core) Accessibility {
	return Accessibility{
		Id:        c.Id,
		Name:      c.Name,
		Longitude: c.Longitude,
		Latitude:  c.Latitude,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func ListToCore(core []Accessibility) (result []accessibility.Core) {
	for _, complex1 := range core {
		result = append(result, toAccessibilityCore(complex1))
	}
	return
}
