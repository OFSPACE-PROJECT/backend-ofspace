package data

import (
	facility "ofspace-be/features/facility"
	"time"
)

type Facility struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toFacilityCore(c Facility) facility.Core {
	return facility.Core{
		Id:        c.Id,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func fromFacilityCore(c facility.Core) Facility {
	return Facility{
		Id:        c.Id,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func ListToCore(core []Facility) (result []facility.Core) {
	for _, complex1 := range core {
		result = append(result, toFacilityCore(complex1))
	}
	return
}
