package domain

import "errors"

var (
	ErrImageNotFound = errors.New("image not found")
	ErrReviewNotFound = errors.New("review not found")
)