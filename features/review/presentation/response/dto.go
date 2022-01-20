package response

import (
	"ofspace-be/features/review"
	"time"
)

type Review struct {
	Id                    uint
	CustomerId            uint      `json:"customer_id"`
	Customer              User      `json:"customer,omitempty"`
	UnitId                uint      `json:"unit_id"`
	Unit                  Unit      `json:"unit,omitempty"`
	RatingAccess          float32   `json:"rating_acces"`
	RatingFacility        float32   `json:"rating_facility"`
	RatingManagement      float32   `json:"rating_management"`
	RatingQuality         float32   `json:"rating_quality"`
	CostumerOverallRating float32   `json:"costumer_overall_rating"`
	Comment               string    `json:"comment"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type User struct {
	ID   uint   `json:"customer_id,omitempty"`
	Name string `json:"customer_name,omitempty"`
}

type Unit struct {
	Id          uint   `json:"unit_id,omitempty"`
	BuildingId  uint   `json:"unit_building_id,omitempty"`
	Description string `json:"unit_description,omitempty"`
	UnitType    string `json:"unit_type,omitempty"`
}

func FromUserCore(req review.User) User {
	return User{
		ID:   req.ID,
		Name: req.Name,
	}
}

func FromUnitCore(req review.Unit) Unit {
	return Unit{
		Id:          req.Id,
		BuildingId:  req.BuildingId,
		Description: req.Description,
		UnitType:    req.UnitType,
	}
}

func FromReviewCore(req review.Core) Review {
	return Review{
		Id:                    req.Id,
		CustomerId:            req.CustomerId,
		Customer:              FromUserCore(req.Customer),
		UnitId:                req.UnitId,
		Unit:                  FromUnitCore(req.Unit),
		RatingAccess:          req.RatingAccess,
		RatingFacility:        req.RatingFacility,
		RatingManagement:      req.RatingManagement,
		RatingQuality:         req.RatingQuality,
		CostumerOverallRating: req.CostumerOverallRating,
		Comment:               req.Comment,
		CreatedAt:             req.CreatedAt,
		UpdatedAt:             req.UpdatedAt,
	}
}

func ToListReview(data []review.Core) (result []Review) {
	result = []Review{}
	for _, review := range data {
		result = append(result, FromReviewCore(review))
	}
	return
}
