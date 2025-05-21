package usecase

import "github.com/LavaJover/shvark-imageboard-service/internal/domain"

type DefaultImageUsecase struct {
	ImageRepo domain.ImageRepository
}

func (uc *DefaultImageUsecase) CreateImage(image *domain.Image) (*domain.Image, error) {
	return uc.ImageRepo.CreateImage(image)
}

func (uc *DefaultImageUsecase) GetImageByID(imageID string) (*domain.Image, error) {
	return uc.ImageRepo.GetImageByID(imageID)
}

func (uc *DefaultImageUsecase) UpdateImage(imageID string, image *domain.Image, fields []string) (*domain.Image, error){
	return uc.ImageRepo.UpdateImage(imageID, image, fields)
}

func (uc *DefaultImageUsecase) DeleteImage(imageID string) (*domain.Image, error) {
	return uc.ImageRepo.DeleteImage(imageID)
}