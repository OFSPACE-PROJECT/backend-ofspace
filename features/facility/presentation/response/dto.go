package response

import "ofspace-be/features/facility"

type Facility struct {
	Id   uint   `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

func FromFacilityCore(c facility.Facility) Facility {
	return Facility{
		Id:   c.Id,
		Name: c.Name,
	}
}
func ToFacilityResponse(c facility.Facility) Facility {
	return Facility{
		Id:   c.Id,
		Name: c.Name,
	}
}

func ToListFacilityCore(core []facility.Facility) (response []Facility) {
	for _, acc := range core {
		response = append(response, FromFacilityCore(acc))
	}
	return
}
