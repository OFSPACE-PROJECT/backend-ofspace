package business

import (
	"ofspace-be/features/review"
	"time"

	"context"
)

type ReviewBusiness struct {
	reviewData     review.Data
	contextTimeout time.Duration
}

func NewReviewBusiness(reviewData review.Data, timeout time.Duration) review.Business {
	return &ReviewBusiness{reviewData: reviewData, contextTimeout: timeout}
}

func (ub *ReviewBusiness) CreateReview(c context.Context, data review.Core) (review.Core, error) {
	ctx, error := context.WithTimeout(c, ub.contextTimeout)
	defer error()

	reviewData, err := ub.reviewData.CreateReview(ctx, data)

	reviewData.CostumerOverallRating = (reviewData.RatingAccess + reviewData.RatingFacility + reviewData.RatingManagement + reviewData.RatingQuality) / 4

	if err != nil {
		return review.Core{}, err
	}
	return reviewData, nil
}

func (ub *ReviewBusiness) GetAllReview(c context.Context, unit uint) ([]review.Core, error) {

	ctx, error := context.WithTimeout(c, ub.contextTimeout)
	defer error()

	res, err := ub.reviewData.GetAllReview(ctx, unit)
	if err != nil {
		return []review.Core{}, err
	}
	for i, num := range res {
		res[i].CostumerOverallRating = (num.RatingAccess + num.RatingFacility + num.RatingManagement + num.RatingQuality) / 4
	}
	return res, nil

}
