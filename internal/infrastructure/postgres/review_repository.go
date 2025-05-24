package postgres

import (
	"github.com/LavaJover/shvark-imageboard-service/internal/domain"
	"gorm.io/gorm"
)

type DefaultReviewRepository struct {
	ReviewDB *gorm.DB
}

func (r *DefaultReviewRepository) CreateReview(review *domain.Review) (*domain.Review, error) {
	reviewModel := ReviewModel{
		ImageID: review.ImageID,
		UserID: review.UserID,
		Text: review.Text,
	}

	if err := r.ReviewDB.Create(&reviewModel).Error; err != nil {
		return nil, err
	}

	return &domain.Review{
		ID: reviewModel.ID,
		ImageID: reviewModel.ImageID,
		UserID: reviewModel.UserID,
		Text: reviewModel.Text,
	}, nil

}

func (r *DefaultReviewRepository) UpdateReview(reviewID string, review *domain.Review, fileds []string) (*domain.Review, error) {
	return nil, nil
}

func (r *DefaultReviewRepository) GetReviewByID(reviewID string) (*domain.Review, error) {
	return nil, nil
}

func (r *DefaultReviewRepository) DeleteReview(reviewID string) (*domain.Review, error) {
	return nil, nil
}