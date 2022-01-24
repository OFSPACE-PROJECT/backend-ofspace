package response

import (
	//"ofspace-be/features/facility"
	"ofspace-be/features/unit"
	"time"
)

type Unit struct {
	Id             uint            `gorm:"primaryKey" json:"id"`
	UserId         uint            `gorm:"not null" json:"user_id"`
	BuildingId     uint            `gorm:"not null" json:"building_id"`
	Description    string          `json:"description"`
	UnitType       string          `gorm:"default:office" json:"unit_type"`
	Price          float32         `json:"price"`
	TotalUnit      int             `json:"total_unit"`
	RemainingUnit  int             `json:"remaining_unit"`
	UnitFacilities []Facility      `gorm:"foreignKey:UnitID;references:ID"`
	InteriorPhotos []InteriorPhoto `gorm:"foreignKey:UnitID;references:ID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Facility struct {
	Id uint `json:"id"`
	//UnitID     uint
	//FacilityID uint
	Name string `json:"name"`
}

type InteriorPhoto struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	UnitID      uint      `json:"unit_id"`
	PhotoURL    string    `json:"photo_url"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToUnitFacilityResponse(c Facility) unit.Facility {
	return unit.Facility{
		Id: c.Id,
		//UnitID: c.UnitId,
		//Id:   c.Id,
		Name: c.Name,
	}
}

func ToUnitResponse(b Unit) unit.Core {
	return unit.Core{
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
func ToInteriorResponse(b InteriorPhoto) unit.InteriorCore {
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

func ToListUnitCore(core []Unit) (response []unit.Core) {
	for _, build := range core {
		response = append(response, ToUnitResponse(build))
	}
	return
}
func FromListUnitCore(core []unit.Core) (response []Unit) {
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
