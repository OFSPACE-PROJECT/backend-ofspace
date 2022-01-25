package response

import (
	building "ofspace-be/features/building/presentation/response"
	complex2 "ofspace-be/features/complex"
)

type Complex struct {
	Id        uint                `json:"id" form:"id"`
	Name      string              `json:"name" form:"name"`
	Buildings []building.Building `json:"buildings" form:"buildings"`
	Address   string              `json:"address" form:"address"`
	Latitude  string              `json:"latitude" form:"latitude"`
	Longitude string              `json:"longitude" form:"longitude"`
}

func FromComplexCore(c complex2.Core) Complex {
	return Complex{
		Id:        c.Id,
		Name:      c.Name,
		Buildings: building.ToListBuildingCore(c.Buildings),
		Address:   c.Address,
		Latitude:  c.Latitude,
		Longitude: c.Longitude,
	}
}
func ToComplexResponse(c complex2.Core) Complex {
	return Complex{
		Id:        c.Id,
		Name:      c.Name,
		Buildings: building.ToListBuildingCore(c.Buildings),
		Address:   c.Address,
		Longitude: c.Longitude,
		Latitude:  c.Latitude,
	}
}

func ToListCompCore(core []complex2.Core) (response []Complex) {
	for _, comp := range core {
		response = append(response, FromComplexCore(comp))
	}
	return
}
