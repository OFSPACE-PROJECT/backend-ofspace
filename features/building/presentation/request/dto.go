package request

import (
	"ofspace-be/features/building"
)

type CreateBuilding struct {
	UserId           uint   `json:"user_id"`
	ComplexId        uint   `json:"complex_id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ImageURL         string `json:"image_url"`
	OfficeHours      string `json:"office_hours"`
	FloorCount       string `json:"floor_count"`
	BuildingSize     string `json:"building_size"`
	AverageFloorSize string `json:"average_floor_size"`
	YearConstructed  string `json:"year_constructed"`
	Lifts            string `json:"lifts"`
	Parking          string `json:"parking"`
	Toilets          string `json:"toilets"`
}

type UpdateBuilding struct {
	Id               uint   `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ImageURL         string `json:"image_url"`
	OfficeHours      string `json:"office_hours"`
	BuildingSize     string `json:"building_size"`
	FloorCount       string `json:"floor_count"`
	AverageFloorSize string `json:"average_floor_size"`
	YearConstructed  string `json:"year_constructed"`
	Lifts            string `json:"lifts"`
	Parking          string `json:"parking"`
	Toilets          string `json:"toilets"`
	BuildingStatus   string `json:"building_status"`
}

type CreateExteriorPhoto struct {
	BuildingId  uint   `json:"building_id"`
	PhotoURL    string `json:"photo_url"`
	Description string `json:"description"`
}

type CreateFloorPhoto struct {
	BuildingId  uint   `json:"building_id"`
	PhotoURL    string `json:"photo_url"`
	Description string `json:"description"`
}

type AddFacility struct {
	BuildingId uint `json:"building_id"`
	FacilityId uint `json:"facility_id"`
}
type UpdateFacility struct {
	//Id         uint   `json:"id"`
	BuildingId uint `json:"building_id"`
	FacilityId uint `json:"facility_id"`
	//Name       string `json:"name"`
}
type UpdateExteriorPhoto struct {
	Id          uint   `json:"id"`
	BuildingId  uint   `json:"building_id"`
	PhotoURL    string `json:"photo_url"`
	Description string `json:"description"`
}
type UpdateFloorPhoto struct {
	Id          uint   `json:"id"`
	BuildingId  uint   `json:"building_id"`
	PhotoURL    string `json:"photo_url"`
	Description string `json:"description"`
}

func (c *UpdateBuilding) ToUpdateCore() building.Core {
	return building.Core{
		Id:               c.Id,
		Name:             c.Name,
		Description:      c.Description,
		OfficeHours:      c.OfficeHours,
		ImageURL:         c.ImageURL,
		BuildingSize:     c.BuildingSize,
		AverageFloorSize: c.AverageFloorSize,
		FloorCount:       c.FloorCount,
		YearConstructed:  c.YearConstructed,
		Lifts:            c.Lifts,
		Parking:          c.Parking,
		Toilets:          c.Toilets,
		BuildingStatus:   c.BuildingStatus,
	}
}
func (c *UpdateExteriorPhoto) ToUpdateExteriorCore() building.ExteriorCore {
	return building.ExteriorCore{
		Id:          c.Id,
		BuildingId:  c.BuildingId,
		PhotoURL:    c.PhotoURL,
		Description: c.Description,
	}
}
func (c *UpdateFacility) ToUpdateFacilityCore() building.Facility {
	return building.Facility{
		Id: c.FacilityId,
		//BuildingId: c.BuildingId,
		//FacilityId: c.FacilityId,
		//Name:       c.Name,
	}
}
func (c *UpdateFloorPhoto) ToUpdateFloorCore() building.FloorCore {
	return building.FloorCore{
		Id:          c.Id,
		BuildingId:  c.BuildingId,
		PhotoURL:    c.PhotoURL,
		Description: c.Description,
	}
}
func (c *CreateBuilding) ToCore() building.Core {
	return building.Core{
		UserId:           c.UserId,
		ComplexId:        c.ComplexId,
		Name:             c.Name,
		Description:      c.Description,
		ImageURL:         c.ImageURL,
		OfficeHours:      c.OfficeHours,
		BuildingSize:     c.BuildingSize,
		FloorCount:       c.FloorCount,
		AverageFloorSize: c.AverageFloorSize,
		YearConstructed:  c.YearConstructed,
		Lifts:            c.Lifts,
		Parking:          c.Parking,
		Toilets:          c.Toilets,
	}
}
func (c *CreateExteriorPhoto) ToExteriorCore() building.ExteriorCore {
	return building.ExteriorCore{
		BuildingId:  c.BuildingId,
		PhotoURL:    c.PhotoURL,
		Description: c.Description,
	}
}

func (c *CreateFloorPhoto) ToFloorCore() building.FloorCore {
	return building.FloorCore{
		BuildingId:  c.BuildingId,
		PhotoURL:    c.PhotoURL,
		Description: c.Description,
	}
}

//func (c *AddFacility) ToFacilityCore() building.Facility {
//	return building.Facility{
//		BuildingId: c.BuildingId,
//		FacilityId: c.FacilityId,
//	}
//}

func ToFacilityCore(c AddFacility) building.Facility {

	return building.Facility{
		Id: c.FacilityId,
		//FacilityId: c.FacilityId,
	}
}
