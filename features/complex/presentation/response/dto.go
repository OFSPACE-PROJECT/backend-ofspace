package response

import complex2 "ofspace-be/features/complex"

type Complex struct {
	Id        uint   `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	Latitude  string `json:"latitude" form:"latitude"`
	Longitude string `json:"longitude" form:"longitude"`
}

func FromComplexCore(c complex2.Core) Complex {
	return Complex{
		Id:        c.Id,
		Name:      c.Name,
		Latitude:  c.Latitude,
		Longitude: c.Longitude,
	}
}
func ToComplexResponse(c complex2.Core) Complex {
	return Complex{
		Id:        c.Id,
		Name:      c.Name,
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
