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
