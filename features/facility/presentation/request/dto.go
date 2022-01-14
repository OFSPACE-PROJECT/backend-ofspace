package request

import (
	"ofspace-be/features/facility"
)

type CreateFacility struct {
	Name string `json:"name"`
}

func (c *CreateFacility) ToCore() facility.Core {
	return facility.Core{
		Name: c.Name,
	}
}

type UpdateFacility struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (c *UpdateFacility) ToUpdateCore() facility.Core {
	return facility.Core{
		Id:   c.Id,
		Name: c.Name,
	}
}
