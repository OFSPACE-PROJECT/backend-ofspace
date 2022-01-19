package data

import (
	"context"
	"ofspace-be/features/review"

	"gorm.io/gorm"
)

type ReviewData struct {
	Connect *gorm.DB
}

func NewReviewData(connect *gorm.DB) review.Data {
	return &ReviewData{
		Connect: connect,
	}
}

func (rep *ReviewData) CreateReview(ctx context.Context, domain review.Core) (review.Core, error) {
	res := FromCore(domain)
	result := rep.Connect.Create(&res)

	if result.Error != nil {
		return review.Core{}, result.Error
	}

	return toReviewCore(res), nil
}

func (rep *ReviewData) GetAllReview(c context.Context, unit uint) ([]review.Core, error) {
	var res []Review
	result := rep.Connect.Preload("Customer").Preload("Unit").Find(&res)
	if result.Error != nil {
		return []review.Core{}, result.Error
	}
	return ToListCore(res), nil
}

func (rep *ReviewData) GetOneReview(ctx context.Context, id uint) (review.Core, error) {
	var res Review
	result := rep.Connect.Preload("Customer").Preload("Unit").First(&res, "id= ?", id)
	if result.Error != nil {
		return review.Core{}, result.Error
	}
	return toReviewCore(res), nil
}

func (rep *ReviewData) UpdateReview(ctx context.Context, domain review.Core) (review.Core, error) {
	res := FromCore(domain)
	result := rep.Connect.Where("id = ?", res.Id).Updates(&Review{Comment: res.Comment, RatingQuality: res.RatingQuality, RatingManagement: res.RatingManagement, RatingFacility: res.RatingFacility, RatingAccess: res.RatingAccess})

	if result.Error != nil {
		return review.Core{}, result.Error
	}

	return toReviewCore(res), nil
}
