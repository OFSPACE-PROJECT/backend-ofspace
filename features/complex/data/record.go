package data

import (
	complex2 "ofspace-be/features/complex"
	"time"
)

type Complex struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	Longitude string
	Latitude  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toComplexCore(c Complex) complex2.Core {
	return complex2.Core{
		Id:        c.Id,
		Name:      c.Name,
		Longitude: c.Longitude,
		Latitude:  c.Latitude,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func fromComplexCore(c complex2.Core) Complex {
	return Complex{
		Id:        c.Id,
		Name:      c.Name,
		Longitude: c.Longitude,
		Latitude:  c.Latitude,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func ListToCore(core []Complex) (result []complex2.Core) {
	for _, complex1 := range core {
		result = append(result, toComplexCore(complex1))
	}
	return
}
