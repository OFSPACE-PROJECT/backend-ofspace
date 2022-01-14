package data

import (
	"ofspace-be/features/building"
	"time"
)

type Building struct {
	Id                 uint `gorm:"primaryKey"`
	UserId             uint `gorm:"not null"`
	ComplexId          uint `gorm:"not null" json:"complex_id"`
	Name               string
	Description        string
	OfficeHours        string
	BuildingSize       string
	AverageFloorSize   string
	YearConstructed    string
	Lifts              string
	Parking            string
	Toilets            string
	BuildingStatus     string          `gorm:"default:unverified" json:"building_status"`
	BuildingFacilities []*Facility     `gorm:"many2many:building_facilities;"`
	ExteriorPhotos     []ExteriorPhoto `gorm:"foreignKey:BuildingID;references:Id"`
	FloorPhotos        []FloorPhoto    `gorm:"foreignKey:BuildingID;references:Id"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type BuildingFacility struct {
	BuildingID uint `gorm:"primaryKey"`
	FacilityID uint `gorm:"primaryKey"`
}

type Facility struct {
	ID   uint
	Name string `gorm:"unique;"`
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

func toBuildingCore(b *Building) building.Core {
	return building.Core{
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
		ExteriorPhotos: ToSliceExteriorPhotoCore(b.ExteriorPhotos),
		FloorPhotos:    ToSliceFloorPhotoCore(b.FloorPhotos),
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
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

func FromBuildingFacilityCore(e building.Facility) BuildingFacility {
	return BuildingFacility{
		BuildingID: e.Id,
		FacilityID: e.Id,
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

func toFacilityCore(b *BuildingFacility) building.Facility {
	return building.Facility{
		Id: b.FacilityID,
		//FacilityId: b.FacilityID,
	}
}

func ToSliceFacilityPhotoCore(e []BuildingFacility) []building.Facility {
	fac := make([]building.Facility, len(e))

	for i, v := range e {
		fac[i] = toFacilityCore(&v)
	}

	return fac
}
