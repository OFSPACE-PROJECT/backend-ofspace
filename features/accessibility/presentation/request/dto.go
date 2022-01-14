package request

import (
	accessibility "ofspace-be/features/accessibility"
)

type CreateAccessibility struct {
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func (c *CreateAccessibility) ToCore() accessibility.Core {
	return accessibility.Core{
		Name:      c.Name,
		Longitude: c.Longitude,
		Latitude:  c.Latitude,
	}
}

type UpdateAccessibility struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func (c *UpdateAccessibility) ToUpdateCore() accessibility.Core {
	return accessibility.Core{
		Id:        c.Id,
		Name:      c.Name,
		Longitude: c.Longitude,
		Latitude:  c.Latitude,
	}
}
