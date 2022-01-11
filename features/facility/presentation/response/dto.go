package response

import "ofspace-be/features/facility"

type Facility struct {
	Id   uint   `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

func FromFacilityCore(c facility.Core) Facility {
	return Facility{
		Id:   c.Id,
		Name: c.Name,
	}
}
func ToFacilityResponse(c facility.Core) Facility {
	return Facility{
		Id:   c.Id,
		Name: c.Name,
	}
}

func ToListFacilityCore(core []facility.Core) (response []Facility) {
	for _, acc := range core {
		response = append(response, FromFacilityCore(acc))
	}
	return
}
