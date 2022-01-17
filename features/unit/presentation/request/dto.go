package request

import (
	"ofspace-be/features/unit"
	"time"
)

type CreateUnit struct {
	UserId      uint    `json:"user_id"`
	BuildingId  uint    `json:"building_id"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	TotalUnit   int     `json:"total_unit"`
}

type UpdateUnit struct {
	Id            uint    `json:"id"`
	UserId        uint    `json:"user_id"`
	BuildingId    uint    `json:"building_id"`
	Description   string  `json:"description"`
	Price         float32 `json:"price"`
	TotalUnit     int     `json:"total_unit"`
	RemainingUnit int     `json:"remaining_unit"`
}

type CreateInteriorPhoto struct {
	UnitId      uint   `json:"unit_id"`
	PhotoURL    string `json:"photo_url"`
	Description string `json:"description"`
}

type AddFacility struct {
	UnitId     uint `json:"unit_id"`
	FacilityId uint `json:"facility_id"`
}
type UpdateFacility struct {
	//Id         uint   `json:"id"`
	UnitId     uint `json:"unit_id"`
	FacilityId uint `json:"facility_id"`
	//Name       string `json:"name"`
}
type UpdateInteriorPhoto struct {
	Id          uint   `json:"id"`
	UnitId      uint   `json:"unit_id"`
	PhotoURL    string `json:"photo_url"`
	Description string `json:"description"`
}

func (b *UpdateUnit) ToUpdateCore() unit.Core {
	return unit.Core{
		Id:            b.Id,
		UserId:        b.UserId,
		BuildingId:    b.BuildingId,
		Description:   b.Description,
		Price:         b.Price,
		TotalUnit:     b.TotalUnit,
		RemainingUnit: b.RemainingUnit,
		UpdatedAt:     time.Time{},
	}
}
func (c *UpdateInteriorPhoto) ToUpdateInteriorCore() unit.InteriorCore {
	return unit.InteriorCore{
		Id:          c.Id,
		UnitId:      c.UnitId,
		PhotoURL:    c.PhotoURL,
		Description: c.Description,
	}
}
func (c *UpdateFacility) ToUpdateFacilityCore() unit.Facility {
	return unit.Facility{
		Id: c.FacilityId,
		//UnitId: c.UnitId,
		//FacilityId: c.FacilityId,
		//Name:       c.Name,
	}
}

func (b *CreateUnit) ToCore() unit.Core {
	return unit.Core{
		UserId:      b.UserId,
		BuildingId:  b.BuildingId,
		Description: b.Description,
		Price:       b.Price,
		TotalUnit:   b.TotalUnit,
		CreatedAt:   time.Time{},
	}
}
func (c *CreateInteriorPhoto) ToInteriorCore() unit.InteriorCore {
	return unit.InteriorCore{
		UnitId:      c.UnitId,
		PhotoURL:    c.PhotoURL,
		Description: c.Description,
	}
}

//func (c *AddFacility) ToFacilityCore() unit.Facility {
//	return unit.Facility{
//		UnitId: c.UnitId,
//		FacilityId: c.FacilityId,
//	}
//}

func ToFacilityCore(c AddFacility) unit.Facility {

	return unit.Facility{
		Id: c.FacilityId,
		//FacilityId: c.FacilityId,
	}
}
