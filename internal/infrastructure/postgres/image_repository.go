package postgres

import (
	"github.com/LavaJover/shvark-imageboard-service/internal/domain"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type DefaultImageRepository struct {
	ImageDB *gorm.DB
}

func ReviewModelsToDomain(reviewModels []ReviewModel) []domain.Review {
	var reviews []domain.Review
	for _, reviewModel := range reviewModels {
		reviews = append(reviews, domain.Review{
			ID: reviewModel.ID,
			ImageID: reviewModel.ImageID,
			UserID: reviewModel.UserID,
			Text: reviewModel.Text,
		})
	}

	return reviews
}

func (r *DefaultImageRepository) CreateImage(image *domain.Image) (*domain.Image, error) {
	var reviews []ReviewModel

	for _, review := range image.Reviews {
		reviews = append(reviews, ReviewModel{
			ID: review.ID,
			ImageID: review.ImageID,
			UserID: review.UserID,
			Text: review.Text,
		})
	}

	imageModel := &ImageModel{
		ID: uuid.New().String(),
		ProfileID: image.ProfileID,
		AvatarUrl: image.AvatarUrl,
		Bio: image.Bio,
		Rating: image.Rating,
		Tags: image.Tags,
		Reviews: reviews,
	}
	if err := r.ImageDB.Create(imageModel).Error; err != nil {
		return nil, err
	}

	image.ID = imageModel.ID

	return image, nil
}

func (r *DefaultImageRepository) UpdateImage(imageID string, image *domain.Image, fields []string) (*domain.Image, error) {
	var imageModel ImageModel
	if err := r.ImageDB.Where("id = ?", imageID).First(&imageModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrImageNotFound
		}
		return nil, err
	}

	for _, field := range fields {
		switch field {
		case "profile_id":
			imageModel.ProfileID = image.ProfileID
		case "avatar_url":
			imageModel.AvatarUrl = image.AvatarUrl
		case "bio":
			imageModel.Bio = image.Bio
		case "rating":
			imageModel.Rating = image.Rating
		}
	}

	if err := r.ImageDB.Save(&imageModel).Error; err != nil {
		return nil, err
	}

	return &domain.Image{
		ID: imageModel.ID,
		ProfileID: imageModel.ProfileID,
		AvatarUrl: imageModel.AvatarUrl,
		Bio: imageModel.Bio,
		Rating: imageModel.Rating,
		Tags: imageModel.Tags,
		Reviews: ReviewModelsToDomain(imageModel.Reviews),
	}, nil
}

func (r *DefaultImageRepository) DeleteImage(imageID string) (*domain.Image, error) {
	var imageModel ImageModel
	if err := r.ImageDB.Where("id = ?", imageID).First(&imageModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrImageNotFound
		}
		return nil, err
	}

	if err := r.ImageDB.Delete(&imageModel).Error; err != nil {
		return nil, err
	}

	return &domain.Image{
		ID: imageModel.ID,
		ProfileID: imageModel.ProfileID,
		AvatarUrl: imageModel.AvatarUrl,
		Bio: imageModel.Bio,
		Rating: imageModel.Rating,
		Tags: imageModel.Tags,
		Reviews: ReviewModelsToDomain(imageModel.Reviews),
	}, nil
}

func (r *DefaultImageRepository) GetImageByID(imageID string) (*domain.Image, error) {
	var imageModel ImageModel
	if err := r.ImageDB.Where("id = ?", imageID).First(&imageModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrImageNotFound
		}
		return nil, err
	}

	return &domain.Image{
		ID: imageModel.ID,
		ProfileID: imageModel.ProfileID,
		AvatarUrl: imageModel.AvatarUrl,
		Bio: imageModel.Bio,
		Rating: imageModel.Rating,
		Tags: imageModel.Tags,
		Reviews: ReviewModelsToDomain(imageModel.Reviews),
	}, nil
}