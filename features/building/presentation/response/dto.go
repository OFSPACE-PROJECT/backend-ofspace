package response

import (
	"ofspace-be/features/building"
	"time"
)

type Building struct {
	Id                 uint `gorm:"primaryKey"`
	UserId             uint `gorm:"not null"`
	ComplexId          uint `gorm:"not null"`
	Name               string
	Description        string
	OfficeHours        string
	BuildingSize       string
	AverageFloorSize   string
	YearConstructed    string
	Lifts              string
	Parking            string
	Toilets            string
	BuildingStatus     string          `gorm:"default:unverified"`
	BuildingFacilities []Facility      `gorm:"foreignKey:BuildingID;references:ID"`
	ExteriorPhotos     []ExteriorPhoto `gorm:"foreignKey:BuildingID;references:ID"`
	FloorPhotos        []FloorPhoto    `gorm:"foreignKey:BuildingID;references:ID"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Facility struct {
	Id   uint
	Name string
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

func ToFacilityResponse(c building.Facility) Facility {
	return Facility{
		Id: c.Id,
		//BuildingID: c.BuildingId,
		//Id:   c.Id,
		Name: c.Name,
	}
}
func ToBuildingResponse(b building.Core) Building {
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
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
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
		//BuildingFacilities: b.BuildingFacilities,
		ExteriorPhotos: fromSliceExteriorCore(b.ExteriorPhotos),
		FloorPhotos:    fromSliceFloorCore(b.FloorPhotos),
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}
}
func ToExteriorResponse(b building.ExteriorCore) ExteriorPhoto {
	return ExteriorPhoto{
		Id:          b.Id,
		BuildingID:  b.BuildingId,
		PhotoURL:    b.PhotoURL,
		Description: b.Description,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}
func ToFloorResponse(b building.FloorCore) FloorPhoto {
	return FloorPhoto{
		Id:          b.Id,
		BuildingID:  b.BuildingId,
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

func ToListBuildingCore(core []building.Core) (response []Building) {
	for _, build := range core {
		response = append(response, FromBuildingCore(build))
	}
	return
}
func ToListExteriorCore(core []building.ExteriorCore) (response []ExteriorPhoto) {
	for _, build := range core {
		response = append(response, FromExteriorPhotoCore(build))
	}
	return
}
func ToListFloorCore(core []building.FloorCore) (response []FloorPhoto) {
	for _, build := range core {
		response = append(response, FromFloorPhotoCore(build))
	}
	return
}
func ToListFacilityCore(core []building.Facility) (response []Facility) {
	for _, build := range core {
		response = append(response, FromBuildingFacilityCore(build))
	}
	return
}

//func ToSliceExteriorPhotoCore(e []ExteriorPhoto) []building.ExteriorCore {
//	photos := make([]building.ExteriorCore, len(e))
//
//	for i, v := range e {
//		photos[i] = toExteriorPhotoCore(&v)
//	}
//
//	return photos
//}
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

//func ToSliceBuildingCore(e []Building) []building.Core {
//	photos := make([]building.FloorCore, len(e))
//
//	for i, v := range e {
//		photos[i] = toFloorPhotoCore(&v)
//	}
//
//	return photos
//}
func fromSliceFloorCore(photos []building.FloorCore) (result []FloorPhoto) {
	for _, photo := range photos {
		result = append(result, FromFloorPhotoCore(photo))
	}
	return
}

func toFacilityCore(b *Facility) building.Facility {
	return building.Facility{
		//BuildingId: b.BuildingID,
		Id:   b.Id,
		Name: b.Name,
	}
}

func ToSliceFacilityPhotoCore(e []Facility) []building.Facility {
	fac := make([]building.Facility, len(e))

	for i, v := range e {
		fac[i] = toFacilityCore(&v)
	}

	return fac
}

func FromBuildingFacilityCore(c building.Facility) Facility {
	return Facility{
		//Id:         c.Id,
		//BuildingID: c.BuildingId,
		Id:   c.Id,
		Name: c.Name,
	}
}
