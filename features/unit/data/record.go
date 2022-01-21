package data

import (
	facility "ofspace-be/features/facility/data"
	"ofspace-be/features/unit"
	"time"
)

type Unit struct {
	Id             uint `gorm:"primaryKey"`
	UserId         uint `gorm:"not null"`
	BuildingId     uint `gorm:"not null" json:"building_id"`
	Description    string
	UnitType       string
	Price          float32
	TotalUnit      int
	RemainingUnit  int
	UnitFacilities []facility.Facility `gorm:"many2many:unit_facilities;"`
	InteriorPhotos []InteriorPhoto     `gorm:"foreignKey:UnitID;references:Id"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type UnitFacility struct {
	UnitID     uint
	FacilityID uint
}

type Facility struct {
	ID   uint
	Name string `gorm:"unique;"`
}

type InteriorPhoto struct {
	Id          uint `gorm:"primaryKey"`
	UnitID      uint
	PhotoURL    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func toUnitCore(b *Unit) unit.Core {
	return unit.Core{
		Id:         b.Id,
		UserId:     b.UserId,
		BuildingId: b.BuildingId,
		//Building:
		Description:   b.Description,
		UnitType:      b.UnitType,
		Price:         b.Price,
		TotalUnit:     b.TotalUnit,
		RemainingUnit: b.RemainingUnit,
		//UnitFacilities: (b.UnitFacilities),
		InteriorPhoto: ToSliceInteriorPhotoCore(b.InteriorPhotos),
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	}
}

func FromUnitCore(b unit.Core) Unit {
	return Unit{
		Id:            b.Id,
		UserId:        b.UserId,
		BuildingId:    b.BuildingId,
		Description:   b.Description,
		UnitType:      b.UnitType,
		Price:         b.Price,
		TotalUnit:     b.TotalUnit,
		RemainingUnit: b.RemainingUnit,
		//UnitFacilities: FromUnitFacilityCore(b.UnitFacilities),
		InteriorPhotos: fromSliceInteriorCore(b.InteriorPhoto),
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}
}

func toInteriorPhotoCore(b *InteriorPhoto) unit.InteriorCore {
	return unit.InteriorCore{
		Id:          b.Id,
		UnitId:      b.UnitID,
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

func FromUnitFacilityCore(e unit.Facility) Facility {
	return Facility{
		ID:   e.Id,
		Name: e.Name,
	}
}

func ToSliceInteriorPhotoCore(e []InteriorPhoto) []unit.InteriorCore {
	photos := make([]unit.InteriorCore, len(e))

	for i, v := range e {
		photos[i] = toInteriorPhotoCore(&v)
	}

	return photos
}
func fromSliceInteriorCore(photos []unit.InteriorCore) (result []InteriorPhoto) {
	for _, photo := range photos {
		result = append(result, FromInteriorPhotoCore(photo))
	}
	return
}
func FromSliceUnitCore(units []unit.Core) (result []Unit) {
	for _, unit1 := range units {
		result = append(result, FromUnitCore(unit1))
	}
	return
}

func ListUnitToCore(buildings []Unit) (result []unit.Core) {
	for _, build := range buildings {
		result = append(result, toUnitCore(&build))
	}
	return
}

func toFacilityCore(b *facility.Facility) unit.Facility {
	return unit.Facility{
		Id:   b.Id,
		Name: b.Name,
		//FacilityId: b.FacilityID,
	}
}

//func ToSliceUnitFacilityCore(e []unit.Facility) []unit.Facility {
//	fac := make([]unit.Facility, len(e))
//
//	for i, v := range e {
//		fac[i] = toFacilityCore(&v)
//	}
//
//	return fac
//}

func FromSliceUnitFacilityCore(e []facility.Facility) []unit.Facility {
	fac := make([]unit.Facility, len(e))

	for i, v := range e {
		fac[i] = toFacilityCore(&v)
	}

	return fac
}
