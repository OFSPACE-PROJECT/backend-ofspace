package response

import accessibility "ofspace-be/features/accessibility"

type Accessibility struct {
	Id        uint   `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	Address   string `json:"address" form:"address"`
	Latitude  string `json:"latitude" form:"latitude"`
	Longitude string `json:"longitude" form:"longitude"`
}

func FromAccessibilityCore(c accessibility.Core) Accessibility {
	return Accessibility{
		Id:        c.Id,
		Name:      c.Name,
		Address:   c.Address,
		Latitude:  c.Latitude,
		Longitude: c.Longitude,
	}
}
func ToAccessibilityResponse(c accessibility.Core) Accessibility {
	return Accessibility{
		Id:        c.Id,
		Name:      c.Name,
		Address:   c.Address,
		Longitude: c.Longitude,
		Latitude:  c.Latitude,
	}
}

func ToListAccCore(core []accessibility.Core) (response []Accessibility) {
	for _, acc := range core {
		response = append(response, FromAccessibilityCore(acc))
	}
	return
}
