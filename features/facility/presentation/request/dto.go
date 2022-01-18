package request

import (
	"ofspace-be/features/facility"
)

type CreateFacility struct {
	Name string `json:"name"`
}

func (c *CreateFacility) ToCore() facility.Facility {
	return facility.Facility{
		Name: c.Name,
	}
}

type UpdateFacility struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (c *UpdateFacility) ToUpdateCore() facility.Facility {
	return facility.Facility{
		Id:   c.Id,
		Name: c.Name,
	}
}
