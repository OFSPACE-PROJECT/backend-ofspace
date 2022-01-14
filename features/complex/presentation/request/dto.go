package request

import (
	complex2 "ofspace-be/features/complex"
)

type CreateComplex struct {
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type UpdateComplex struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func (c *CreateComplex) ToCore() complex2.Core {
	return complex2.Core{
		Name:      c.Name,
		Longitude: c.Longitude,
		Latitude:  c.Latitude,
	}
}

func (c *UpdateComplex) ToUpdateCore() complex2.Core {
	return complex2.Core{
		Id:        c.Id,
		Name:      c.Name,
		Longitude: c.Longitude,
		Latitude:  c.Latitude,
	}
}
