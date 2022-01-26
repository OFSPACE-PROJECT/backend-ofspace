package response

import (
	"ofspace-be/features/building"
	"ofspace-be/features/facility"
	review "ofspace-be/features/review/presentation/response"
	"ofspace-be/features/unit/presentation/response"
	"time"
)

type Building struct {
	Id                 uint            `gorm:"primaryKey" json:"id"`
	UserId             uint            `gorm:"not null" json:"user_id"`
	ComplexId          uint            `gorm:"not null" json:"complex_id"`
	Name               string          `json:"name"`
	Description        string          `json:"description"`
	ImageURL           string          `json:"image_url"`
	OfficeHours        string          `json:"office_hours"`
	BuildingSize       string          `json:"building_size"`
	AverageFloorSize   string          `json:"average_floor_size"`
	YearConstructed    string          `json:"year_constructed"`
	Lifts              string          `json:"lifts"`
	Parking            string          `json:"parking"`
	FloorCount         string          `json:"floor_count"`
	Toilets            string          `json:"toilets"`
	BuildingStatus     string          `gorm:"default:unverified" json:"building_status"`
	Units              []response.Unit `json:"units" gorm:"foreignKey:BuildingId"`
	Reviews            []review.Review `json:"reviews" gorm:"foreignKey:BuildingId"`
	BuildingFacilities []Facility      `gorm:"foreignKey:BuildingID;references:ID" json:"building_facilities"`
	ExteriorPhotos     []ExteriorPhoto `gorm:"foreignKey:BuildingID;references:ID" json:"exterior_photos"`
	FloorPhotos        []FloorPhoto    `gorm:"foreignKey:BuildingID;references:ID" json:"floor_photos"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
}

type Facility struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type ExteriorPhoto struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	BuildingID  uint      `json:"building_id"`
	PhotoURL    string    `json:"photo_url"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type FloorPhoto struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	BuildingID  uint      `json:"building_id"`
	PhotoURL    string    `json:"photo_url"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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
		ImageURL:         b.ImageURL,
		OfficeHours:      b.OfficeHours,
		BuildingSize:     b.BuildingSize,
		AverageFloorSize: b.AverageFloorSize,
		YearConstructed:  b.YearConstructed,
		Lifts:            b.Lifts,
		Parking:          b.Parking,
		Toilets:          b.Toilets,
		FloorCount:       b.FloorCount,
		Units:            response.FromListUnitCore(b.Units),
		Reviews:          review.ToListReview(b.Reviews),
		BuildingStatus:   b.BuildingStatus,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
	}
}

func FromBuildingCore(b building.Core) Building {
	return Building{
		Id:                 b.Id,
		UserId:             b.UserId,
		ComplexId:          b.ComplexId,
		Name:               b.Name,
		Description:        b.Description,
		ImageURL:           b.ImageURL,
		OfficeHours:        b.OfficeHours,
		BuildingSize:       b.BuildingSize,
		AverageFloorSize:   b.AverageFloorSize,
		YearConstructed:    b.YearConstructed,
		Lifts:              b.Lifts,
		Parking:            b.Parking,
		FloorCount:         b.FloorCount,
		Toilets:            b.Toilets,
		BuildingStatus:     b.BuildingStatus,
		Units:              response.FromListUnitCore(b.Units),
		Reviews:            review.ToListReview(b.Reviews),
		BuildingFacilities: fromSliceFacilityCore(b.BuildingFacilities),
		ExteriorPhotos:     fromSliceExteriorCore(b.ExteriorPhotos),
		FloorPhotos:        fromSliceFloorCore(b.FloorPhotos),
		CreatedAt:          time.Time{},
		UpdatedAt:          time.Time{},
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
func fromSliceFacilityCore(photos []building.Facility) (result []Facility) {
	for _, photo := range photos {
		result = append(result, FromBuildingFacilityCore(photo))
	}
	return
}
func toFacilityCore(b *facility.Facility) building.Facility {
	return building.Facility{
		//BuildingId: b.BuildingID,
		Id:   b.Id,
		Name: b.Name,
	}
}

func ToSliceFacilityCore(e []facility.Facility) []building.Facility {
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
