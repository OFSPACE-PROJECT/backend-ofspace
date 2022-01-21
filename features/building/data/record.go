package data

import (
	"ofspace-be/features/building"
	facility2 "ofspace-be/features/facility"
	unit "ofspace-be/features/unit/data"

	//"ofspace-be/features/facility"
	facility "ofspace-be/features/facility/data"
	"time"
)

type Building struct {
	Id                 uint        `gorm:"primaryKey"`
	UserId             uint        `gorm:"not null"`
	ComplexId          uint        `gorm:"not null" json:"complex_id"`
	Units              []unit.Unit `gorm:"foreignKey:BuildingId"`
	Name               string
	Description        string
	OfficeHours        string
	BuildingSize       string
	AverageFloorSize   string
	YearConstructed    string
	Lifts              string
	Parking            string
	Toilets            string
	BuildingStatus     string              `gorm:"default:unverified" json:"building_status"`
	BuildingFacilities []facility.Facility `gorm:"many2many:building_facilities;"`
	ExteriorPhotos     []ExteriorPhoto     `gorm:"foreignKey:BuildingID;references:Id"`
	FloorPhotos        []FloorPhoto        `gorm:"foreignKey:BuildingID;references:Id"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Facility struct {
	ID   uint
	Name string
}

type BuildingFacility struct {
	BuildingID uint
	FacilityID uint
}

type ExteriorPhoto struct {
	Id          uint `gorm:"primaryKey"`
	BuildingID  uint
	PhotoURL    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type FloorPhoto struct {
	Id          uint `gorm:"primaryKey"`
	BuildingID  uint
	PhotoURL    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var a = building.Core{}

func toBuildingCore(b *Building) building.Core {
	return building.Core{
		Id:                 b.Id,
		UserId:             b.UserId,
		ComplexId:          b.ComplexId,
		Name:               b.Name,
		Description:        b.Description,
		OfficeHours:        b.OfficeHours,
		BuildingSize:       b.BuildingSize,
		AverageFloorSize:   b.AverageFloorSize,
		YearConstructed:    b.YearConstructed,
		Lifts:              b.Lifts,
		Parking:            b.Parking,
		Toilets:            b.Toilets,
		BuildingStatus:     b.BuildingStatus,
		Units:              unit.ListUnitToCore(b.Units),
		BuildingFacilities: ToSliceFacilityCore(b.BuildingFacilities),
		ExteriorPhotos:     ToSliceExteriorPhotoCore(b.ExteriorPhotos),
		FloorPhotos:        ToSliceFloorPhotoCore(b.FloorPhotos),
		CreatedAt:          time.Time{},
		UpdatedAt:          time.Time{},
	}
}

func toFacilityCore(b *facility.Facility) building.Facility {
	return building.Facility{
		Id:   b.Id,
		Name: b.Name,
		//FacilityId: b.FacilityID,
	}
}

func FromBuildingCore(b building.Core) Building {
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
		Units:            unit.FromSliceUnitCore(b.Units),
		//BuildingFacilities: FromBuildingFacilityCore(c.BuildingFacilities,
		ExteriorPhotos: fromSliceExteriorCore(b.ExteriorPhotos),
		FloorPhotos:    fromSliceFloorCore(b.FloorPhotos),
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}
}

func toExteriorPhotoCore(b *ExteriorPhoto) building.ExteriorCore {
	return building.ExteriorCore{
		Id:          b.Id,
		BuildingId:  b.BuildingID,
		PhotoURL:    b.PhotoURL,
		Description: b.Description,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

func FromExteriorPhotoCore(e building.ExteriorCore) ExteriorPhoto {
	return ExteriorPhoto{
		Id:          e.Id,
		BuildingID:  e.BuildingId,
		PhotoURL:    e.PhotoURL,
		Description: e.Description,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

func FromBuildingFacilityCore(e facility2.Facility) facility.Facility {
	return facility.Facility{
		Id:   e.Id,
		Name: e.Name,
	}
}

func (c *Building) ToBuildingCore() building.Core {
	return building.Core{
		Id:                 c.Id,
		UserId:             c.UserId,
		ComplexId:          c.ComplexId,
		Name:               c.Name,
		Description:        c.Description,
		OfficeHours:        c.OfficeHours,
		BuildingSize:       c.BuildingSize,
		AverageFloorSize:   c.AverageFloorSize,
		YearConstructed:    c.YearConstructed,
		Lifts:              c.Lifts,
		Parking:            c.Parking,
		Toilets:            c.Toilets,
		BuildingStatus:     c.BuildingStatus,
		BuildingFacilities: ToSliceFacilityCore(c.BuildingFacilities),
		ExteriorPhotos:     ToSliceExteriorPhotoCore(c.ExteriorPhotos),
		FloorPhotos:        ToSliceFloorPhotoCore(c.FloorPhotos),
		CreatedAt:          time.Time{},
		UpdatedAt:          time.Time{},
	}
}

//func (c *building.Core) FromBuildingCore() Building {
//	return Building{
//		Id:                 c.Id,
//		UserId:             c.UserId,
//		ComplexId:          c.ComplexId,
//		Name:               c.Name,
//		Description:        c.Description,
//		OfficeHours:        c.OfficeHours,
//		BuildingSize:       c.BuildingSize,
//		AverageFloorSize:   c.AverageFloorSize,
//		YearConstructed:    c.YearConstructed,
//		Lifts:              c.Lifts,
//		Parking:            c.Parking,
//		Toilets:            c.Toilets,
//		BuildingStatus:     c.BuildingStatus,
//		BuildingFacilities: ToSliceFacilityCore(c.BuildingFacilities),
//		ExteriorPhotos:     ToSliceExteriorPhotoCore(c.ExteriorPhotos),
//		FloorPhotos:        ToSliceFloorPhotoCore(c.FloorPhotos),
//		CreatedAt:          time.Time{},
//		UpdatedAt:          time.Time{},
//	}
//}

func ToListBuildingCore(data []Building) (result []building.Core) {
	result = []building.Core{}
	for _, builds := range data {
		result = append(result, builds.ToBuildingCore())
	}
	return
}
func FromListBuildingCore(data []building.Core) (result []Building) {
	result = []Building{}
	for _, builds := range data {
		result = append(result, FromBuildingCore(builds))
	}
	return
}

func ToListFacilityBuildingCore(data []Building) (result []building.Core) {
	result = []building.Core{}
	for _, builds := range data {
		result = append(result, builds.ToBuildingCore())
	}
	return
}

//func FromBuildingFacilityCore(e building.Facility) Facility {
//	return Facility{Name: e.Name, ID: e.Id}
//}
func toFloorPhotoCore(b *FloorPhoto) building.FloorCore {
	return building.FloorCore{
		Id:          b.Id,
		BuildingId:  b.BuildingID,
		PhotoURL:    b.PhotoURL,
		Description: b.Description,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}
func FromFloorPhotoCore(e building.FloorCore) FloorPhoto {
	return FloorPhoto{
		Id:          e.Id,
		BuildingID:  e.BuildingId,
		PhotoURL:    e.PhotoURL,
		Description: e.Description,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

func ToSliceExteriorPhotoCore(e []ExteriorPhoto) []building.ExteriorCore {
	photos := make([]building.ExteriorCore, len(e))

	for i, v := range e {
		photos[i] = toExteriorPhotoCore(&v)
	}

	return photos
}
func fromSliceExteriorCore(photos []building.ExteriorCore) (result []ExteriorPhoto) {
	for _, photo := range photos {
		result = append(result, FromExteriorPhotoCore(photo))
	}
	return
}

func ToSliceFloorPhotoCore(e []FloorPhoto) []building.FloorCore {
	photos := make([]building.FloorCore, len(e))

	for i, v := range e {
		photos[i] = toFloorPhotoCore(&v)
	}

	return photos
}

func fromSliceFloorCore(photos []building.FloorCore) (result []FloorPhoto) {
	for _, photo := range photos {
		result = append(result, FromFloorPhotoCore(photo))
	}
	return
}

func ListBuildingToCore(buildings []Building) (result []building.Core) {
	for _, build := range buildings {
		result = append(result, toBuildingCore(&build))
	}
	return
}

//func toFacilityCore(b *Facility) building.Facility {
//	return building.Facility{
//		Id: b.ID,
//		//BuildingId: b.BuildingID,
//		//FacilityId: b.FacilityID,
//	}
//}

func ToSliceFacilityCore(e []facility.Facility) []building.Facility {
	fac := make([]building.Facility, len(e))

	for i, v := range e {
		fac[i] = toFacilityCore(&v)
	}

	return fac
}

//func ToSliceFacilityPhotoCore(e []Facility) []building.Facility {
//	fac := make([]building.Facility, len(e))
//
//	for i, v := range e {
//		fac[i] = toFacilityCore(&v)
//	}
//
//	return fac
//}
