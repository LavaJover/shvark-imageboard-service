package grpcapi

import (
	"context"

	"github.com/LavaJover/shvark-imageboard-service/internal/domain"
	"github.com/LavaJover/shvark-imageboard-service/internal/usecase"
	imagepb "github.com/LavaJover/shvark-imageboard-service/proto/gen"
)

type ImageHandler struct {
	imagepb.UnimplementedImageServiceServer
	Uc usecase.DefaultImageUsecase
}

func (h *ImageHandler) CreateImage(ctx context.Context, r *imagepb.CreateImageRequest) (*imagepb.CreateImageResponse, error) {
	image := &domain.Image{
		ProfileID: r.ProfileId,
		AvatarUrl: r.AvatarUrl,
		Bio: r.Bio,
		Rating: r.Rating,
	}

	respImage, err := h.Uc.CreateImage(image)
	if err != nil {
		return nil, err
	}

	return &imagepb.CreateImageResponse{
		ImageId: respImage.ID,
	}, nil
}

func (h *ImageHandler) UpdateImage(ctx context.Context, r *imagepb.UpdateImageRequest) (*imagepb.UpdateImageResponse, error) {
	imageID := r.ImageId
	image := &domain.Image{
		ID: r.Image.ImageId,
		ProfileID: r.Image.ProfileId,
		AvatarUrl: r.Image.AvatarUrl,
		Bio: r.Image.AvatarUrl,
		Rating: r.Image.Rating,
	}
	fields := r.UpdateMask.Paths

	respImage, err := h.Uc.UpdateImage(imageID, image, fields)
	if err != nil {
		return nil, err
	}

	return &imagepb.UpdateImageResponse{
		ImageId: respImage.ID,
		ProfileId: respImage.ProfileID,
		AvatarUrl: respImage.AvatarUrl,
		Bio: respImage.Bio,
		Rating: respImage.Rating,
	}, nil
}

func (h *ImageHandler) DeleteImage(ctx context.Context, r *imagepb.DeleteImageRequest) (*imagepb.DeleteImageResponse, error) {
	imageID := r.ImageId
	respImage, err := h.Uc.DeleteImage(imageID)
	if err != nil {
		return nil, err
	}

	return &imagepb.DeleteImageResponse{
		ImageId: respImage.ID,
		ProfileId: respImage.ProfileID,
		AvatarUrl: respImage.AvatarUrl,
		Bio: respImage.Bio,
		Rating: respImage.Rating,
	}, nil
}

func (h *ImageHandler) GetImageByID(ctx context.Context, r *imagepb.GetImageByIDRequest) (*imagepb.GetImageByIDResponse, error) {
	imageID := r.ImageId
	respImage, err := h.Uc.GetImageByID(imageID)
	if err != nil {
		return nil, err
	}

	return &imagepb.GetImageByIDResponse{
		ImageId: respImage.ID,
		ProfileId: respImage.ProfileID,
		AvatarUrl: respImage.AvatarUrl,
		Bio: respImage.Bio,
		Rating: respImage.Rating,
	}, nil
}