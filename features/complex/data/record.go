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

type Building struct {
	Id               uint `gorm:"primaryKey"`
	UserId           uint
	ComplexId        uint
	Name             string
	Description      string
	OfficeHours      string
	BuildingSize     string
	AverageFloorSize string
	YearConstructed  string
	Lifts            string
	Parking          string
	Toilets          string
	BuildingStatus   string `gorm:"default:unverified" json:"building_status"`
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

func toBuildingCore(b *Building) complex2.Building {
	return complex2.Building{
		Id:               b.Id,
		UserId:           b.UserId,
		ComplexId:        b.ComplexId,
		Name:             b.Name,
		Description:      b.Description,
		OfficeHours:      b.OfficeHours,
		BuildingSize:     b.BuildingSize,
		AverageFloorSize: b.AverageFloorSize,
		YearConstructed:  b.YearConstructed,
		Lifts:            b.Lifts,
		Parking:          b.Parking,
		Toilets:          b.Toilets,
		BuildingStatus:   b.BuildingStatus,
	}
}

func ToListBuildingCore(data []Building) (result []complex2.Building) {
	result = []complex2.Building{}
	for _, builds := range data {
		result = append(result, toBuildingCore(&builds))
	}
	return
}
func FromListBuildingCore(data []complex2.Building) (result []Building) {
	result = []Building{}
	for _, builds := range data {
		result = append(result, FromBuildingCore(&builds))
	}
	return
}

func FromBuildingCore(b *complex2.Building) Building {
	return Building{
		Id:               b.Id,
		UserId:           b.UserId,
		ComplexId:        b.ComplexId,
		Name:             b.Name,
		Description:      b.Description,
		OfficeHours:      b.OfficeHours,
		BuildingSize:     b.BuildingSize,
		AverageFloorSize: b.AverageFloorSize,
		YearConstructed:  b.YearConstructed,
		Lifts:            b.Lifts,
		Parking:          b.Parking,
		Toilets:          b.Toilets,
		BuildingStatus:   b.BuildingStatus,
	}
}
