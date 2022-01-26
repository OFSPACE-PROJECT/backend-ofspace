package data

import (
	building "ofspace-be/features/building/data"
	complex2 "ofspace-be/features/complex"
	"time"
)

type Complex struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	Address   string
	Longitude string
	Latitude  string
	Buildings []building.Building `gorm:"foreignKey:ComplexId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toComplexCore(c Complex) complex2.Core {
	return complex2.Core{
		Id:        c.Id,
		Name:      c.Name,
		Address:   c.Address,
		Longitude: c.Longitude,
		Latitude:  c.Latitude,
		Buildings: building.ToListBuildingCore(c.Buildings),
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func fromComplexCore(c complex2.Core) Complex {
	return Complex{
		Id:        c.Id,
		Name:      c.Name,
		Address:   c.Address,
		Longitude: c.Longitude,
		Latitude:  c.Latitude,
		Buildings: building.FromListBuildingCore(c.Buildings),
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
func toSliceComplexCore(complexes []Complex) (result []complex2.Core) {
	for _, complex1 := range complexes {
		result = append(result, toComplexCore(complex1))
	}
	return
}
