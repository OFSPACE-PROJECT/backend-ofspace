package request

import "ofspace-be/features/review"

type CreateReview struct {
	CustomerId       uint    `json:"customer_id"`
	UnitId           uint    `json:"unit_id"`
	BuildingId       uint    `json:"building_id"`
	BookingId        uint    `json:"booking_id"`
	RatingAccess     float32 `json:"rating_acces"`
	RatingFacility   float32 `json:"rating_facility"`
	RatingManagement float32 `json:"rating_management"`
	RatingQuality    float32 `json:"rating_quality"`
	Comment          string  `json:"comment"`
}

func (r *CreateReview) ToCore() review.Core {
	return review.Core{
		CustomerId:       r.CustomerId,
		UnitId:           r.UnitId,
		BuildingId:       r.BuildingId,
		BookingId:        r.BookingId,
		RatingAccess:     r.RatingAccess,
		RatingFacility:   r.RatingFacility,
		RatingManagement: r.RatingManagement,
		RatingQuality:    r.RatingQuality,
		Comment:          r.Comment,
	}
}
