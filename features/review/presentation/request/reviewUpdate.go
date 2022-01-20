package request

import "ofspace-be/features/review"

type ReviewUpdate struct {
	Id               uint    `json:"id"`
	CustomerId       uint    `json:"customer_id"`
	UnitId           uint    `json:"unit_id"`
	BookingId        uint    `json:"booking_id"`
	RatingAccess     float32 `json:"rating_acces"`
	RatingFacility   float32 `json:"rating_facility"`
	RatingManagement float32 `json:"rating_management"`
	RatingQuality    float32 `json:"rating_quality"`
	Comment          string  `json:"comment"`
}

func (r *ReviewUpdate) ToCoreUpdate() review.Core {
	return review.Core{
		Id:               r.Id,
		CustomerId:       r.CustomerId,
		UnitId:           r.UnitId,
		BookingId:        r.BookingId,
		RatingAccess:     r.RatingAccess,
		RatingFacility:   r.RatingFacility,
		RatingManagement: r.RatingManagement,
		RatingQuality:    r.RatingQuality,
		Comment:          r.Comment,
	}
}
