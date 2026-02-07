package domain

import (
	"context"
)

// ReviewRepository defines the interface for review data access.
type ReviewRepository interface {
	// ListReviewsByProduct returns all reviews for a specific product.
	ListReviewsByProduct(ctx context.Context, productID int) ([]Review, error)

	// CreateReview creates a new review.
	CreateReview(ctx context.Context, review Review) (*Review, error)

	// GetReviewByProductAndUser checks if a review already exists for a product by a user.
	GetReviewByProductAndUser(ctx context.Context, productID, userID int) (*Review, error)
}
