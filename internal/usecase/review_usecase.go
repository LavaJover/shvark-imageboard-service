package usecase

import "github.com/LavaJover/shvark-imageboard-service/internal/domain"

type DefaultReviewUsecase struct {
	ReviewRepo domain.ReviewRepository
}

func (uc *DefaultReviewUsecase) CreateReview(review *domain.Review) (*domain.Review, error) {
	return uc.ReviewRepo.CreateReview(review)
}

func (uc *DefaultReviewUsecase) GetReviewByID(reviewID string) (*domain.Review, error) {
	return uc.ReviewRepo.GetReviewByID(reviewID)
}

func (uc *DefaultReviewUsecase) UpdateReview(reviewID string, review *domain.Review, fields []string) (*domain.Review, error){
	return uc.ReviewRepo.UpdateReview(reviewID, review, fields)
}

func (uc *DefaultReviewUsecase) DeleteReview(reviewID string) (*domain.Review, error) {
	return uc.ReviewRepo.DeleteReview(reviewID)
}