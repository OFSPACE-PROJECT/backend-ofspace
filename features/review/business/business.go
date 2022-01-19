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
