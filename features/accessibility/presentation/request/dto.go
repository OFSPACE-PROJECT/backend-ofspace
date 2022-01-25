package request

import (
	accessibility "ofspace-be/features/accessibility"
)

type CreateAccessibility struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func (c *CreateAccessibility) ToCore() accessibility.Core {
	return accessibility.Core{
		Name:      c.Name,
		Address:   c.Address,
		Longitude: c.Longitude,
		Latitude:  c.Latitude,
	}
}

type UpdateAccessibility struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func (c *UpdateAccessibility) ToUpdateCore() accessibility.Core {
	return accessibility.Core{
		Id:        c.Id,
		Name:      c.Name,
		Address:   c.Address,
		Longitude: c.Longitude,
		Latitude:  c.Latitude,
	}
}
