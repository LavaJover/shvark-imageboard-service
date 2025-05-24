package grpcapi

import (
	"context"

	"github.com/LavaJover/shvark-imageboard-service/internal/domain"
	"github.com/LavaJover/shvark-imageboard-service/internal/usecase"
	imagepb "github.com/LavaJover/shvark-imageboard-service/proto/gen"
)

type ReviewHandler struct {
	imagepb.UnimplementedReviewServiceServer
	Uc *usecase.DefaultReviewUsecase
}

func (h *ReviewHandler) CreateReview(ctx context.Context, r *imagepb.CreateReviewRequest) (*imagepb.CreateReviewResponse, error) {
	review := domain.Review{
		ImageID: r.ImageId,
		UserID: r.ImageId,
		Text: r.Text,
	}
	reviewResp, err := h.Uc.CreateReview(&review)
	if err != nil {
		return nil, err
	}

	return &imagepb.CreateReviewResponse{
		ReviewId: reviewResp.ID,
	}, nil
}

func (h *ReviewHandler) UpdateReview(ctx context.Context, r *imagepb.UpdateReviewRequest) (*imagepb.UpdateReviewResponse, error) {
	reviewID := r.ReviewId
	review := domain.Review{
		ID: r.Review.Id,
		ImageID: r.Review.ImageId,
		UserID: r.Review.UserId,
		Text: r.Review.Text,
	}

	reviewResp, err := h.Uc.UpdateReview(reviewID, &review, r.UpdateMask.Paths)
	if err != nil {
		return nil, err
	}

	return &imagepb.UpdateReviewResponse{
		Id: reviewResp.ID,
		ImageId: reviewResp.ImageID,
		UserId: reviewResp.UserID,
		Text: reviewResp.Text,
	}, nil
}

func (h *ReviewHandler) DeleteReview(ctx context.Context, r *imagepb.DeleteReviewRequest) (*imagepb.DeleteReviewResponse, error) {
	reviewID := r.ReviewId
	reviewResp, err := h.Uc.DeleteReview(reviewID)
	if err != nil {
		return nil, err
	}

	return &imagepb.DeleteReviewResponse{
		Id: reviewResp.ID,
		ImageId: reviewResp.ImageID,
		UserId: reviewResp.UserID,
		Text: reviewResp.Text,
	}, nil
}

func (h *ReviewHandler) GetReviewByID(ctx context.Context, r *imagepb.GetReviewByIDRequest) (*imagepb.GetReviewByIDResponse, error) {
	reviewID := r.ReviewId
	reviewResp, err := h.Uc.GetReviewByID(reviewID)
	if err != nil {
		return nil, err
	}

	return &imagepb.GetReviewByIDResponse{
		Id: reviewResp.ID,
		ImageId: reviewResp.ImageID,
		UserId: reviewResp.UserID,
		Text: reviewResp.Text,
	}, nil
}