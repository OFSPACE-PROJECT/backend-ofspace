package data

import (
	facility "ofspace-be/features/facility"
	"time"
)

type Facility struct {
	Id   uint `gorm:"primaryKey"`
	Name string
	//Buildings []building.Building `gorm:"many2many:building_facilities;"`
	//Units     []data.Unit         `gorm:"many2many:unit_facilities;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toFacilityCore(c Facility) facility.Facility {
	return facility.Facility{
		Id:        c.Id,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func fromFacilityCore(c facility.Facility) Facility {
	return Facility{
		Id:        c.Id,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func ListToCore(core []Facility) (result []facility.Facility) {
	for _, complex1 := range core {
		result = append(result, toFacilityCore(complex1))
	}
	return
}
