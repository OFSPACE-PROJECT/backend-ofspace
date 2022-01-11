package request

import (
	complex2 "ofspace-be/features/complex"
)

type CreateComplex struct {
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
