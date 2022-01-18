package response

import (
	//"ofspace-be/features/facility"
	"ofspace-be/features/unit"
	"time"
)

type Unit struct {
	Id             uint `gorm:"primaryKey"`
	UserId         uint `gorm:"not null"`
	BuildingId     uint `gorm:"not null"`
	Description    string
	UnitType       string
	Price          float32
	TotalUnit      int
	RemainingUnit  int
	UnitFacilities []Facility      `gorm:"foreignKey:UnitID;references:ID"`
	InteriorPhotos []InteriorPhoto `gorm:"foreignKey:UnitID;references:ID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Facility struct {
	Id uint
	//UnitID     uint
	//FacilityID uint
	Name string
}

type InteriorPhoto struct {
	Id          uint `gorm:"primaryKey"`
	UnitID      uint
	PhotoURL    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ToUnitFacilityResponse(c unit.Facility) Facility {
	return Facility{
		Id: c.Id,
		//UnitID: c.UnitId,
		//Id:   c.Id,
		Name: c.Name,
	}
}

func ToUnitResponse(b unit.Core) Unit {
	return Unit{
		Id:            b.Id,
		UserId:        b.UserId,
		BuildingId:    b.BuildingId,
		Description:   b.Description,
		UnitType:      b.UnitType,
		Price:         b.Price,
		TotalUnit:     b.TotalUnit,
		RemainingUnit: b.RemainingUnit,
		UpdatedAt:     time.Time{},
	}
}

func FromUnitCore(b unit.Core) Unit {
	return Unit{
		Id:             b.Id,
		UserId:         b.UserId,
		BuildingId:     b.BuildingId,
		Description:    b.Description,
		UnitType:       b.UnitType,
		Price:          b.Price,
		TotalUnit:      b.TotalUnit,
		RemainingUnit:  b.RemainingUnit,
		UnitFacilities: fromSliceUnitFacilityCore(b.UnitFacilities),
		InteriorPhotos: fromSliceInteriorCore(b.InteriorPhoto),
		UpdatedAt:      time.Time{},
	}
}
func fromSliceUnitFacilityCore(facilities []unit.Facility) (result []Facility) {
	for _, facility := range facilities {
		result = append(result, FromBuildingFacilityCore(facility))
	}
	return
}

func FromBuildingFacilityCore(c unit.Facility) Facility {
	return Facility{
		Id:   c.Id,
		Name: c.Name,
	}
}
func ToInteriorResponse(b unit.InteriorCore) InteriorPhoto {
	return InteriorPhoto{
		Id:          b.Id,
		UnitID:      b.UnitId,
		PhotoURL:    b.PhotoURL,
		Description: b.Description,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

func FromInteriorPhotoCore(e unit.InteriorCore) InteriorPhoto {
	return InteriorPhoto{
		Id:          e.Id,
		UnitID:      e.UnitId,
		PhotoURL:    e.PhotoURL,
		Description: e.Description,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

func ToListUnitCore(core []unit.Core) (response []Unit) {
	for _, build := range core {
		response = append(response, FromUnitCore(build))
	}
	return
}
func ToListInteriorCore(core []unit.InteriorCore) (response []InteriorPhoto) {
	for _, build := range core {
		response = append(response, FromInteriorPhotoCore(build))
	}
	return
}

func ToListFacilityCore(core []unit.Facility) (response []Facility) {
	for _, build := range core {
		response = append(response, FromUnitFacilityCore(build))
	}
	return
}

//func ToSliceInteriorPhotoCore(e []InteriorPhoto) []unit.InteriorCore {
//	photos := make([]unit.InteriorCore, len(e))
//
//	for i, v := range e {
//		photos[i] = toInteriorPhotoCore(&v)
//	}
//
//	return photos
//}
func fromSliceInteriorCore(photos []unit.InteriorCore) (result []InteriorPhoto) {
	for _, photo := range photos {
		result = append(result, FromInteriorPhotoCore(photo))
	}
	return
}

func toFacilityCore(b *Facility) unit.Facility {
	return unit.Facility{
		//UnitId: b.UnitID,
		Id:   b.Id,
		Name: b.Name,
	}
}

func ToSliceFacilityPhotoCore(e []Facility) []unit.Facility {
	fac := make([]unit.Facility, len(e))

	for i, v := range e {
		fac[i] = toFacilityCore(&v)
	}

	return fac
}

func FromUnitFacilityCore(c unit.Facility) Facility {
	return Facility{
		//Id:         c.Id,
		//UnitID: c.UnitId,
		Id:   c.Id,
		Name: c.Name,
	}
}
