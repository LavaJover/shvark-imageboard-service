package domain

type ImageRepository interface {
	CreateImage(image *Image) (*Image, error)
	GetImageByID(imageID string) (*Image, error)
	UpdateImage(imageID string, image *Image, fields []string) (*Image, error)
	DeleteImage(imageID string) (*Image, error)
}

type ReviewRepository interface {
	CreateReview(review *Review) (*Review, error)
	GetReviewByID(reviewID string) (*Review, error)
	UpdateReview(reviewID string, review *Review, fields []string) (*Review, error)
	DeleteReview(reviewID string) (*Review, error)
}