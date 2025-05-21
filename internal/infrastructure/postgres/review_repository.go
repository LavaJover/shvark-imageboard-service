package postgres

import (
	"github.com/LavaJover/shvark-imageboard-service/internal/domain"
	"gorm.io/gorm"
)

type DefaultReviewRepository struct {
	ReviewDB *gorm.DB
}

func (r DefaultReviewRepository) CreateReview(review *domain.Review) (*domain.Review, error) {
	return nil, nil
}

func (r DefaultReviewRepository) UpdateReview(reviewID string, review *domain.Review, fileds []string) (*domain.Review, error) {
	return nil, nil
}

func (r DefaultReviewRepository) GetReviewByID(reviewID string) (*domain.Review, error) {
	return nil, nil
}

func (r *DefaultReviewRepository) DeleteReview(reviewID string) (*domain.Review, error) {
	return nil, nil
}