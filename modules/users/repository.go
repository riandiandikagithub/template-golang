package users

import (
	"context"
	"payment-simulator/models"
)

// Repository represent the article's repository contract
type Repository interface {
	// 	GetByID(ctx context.Context, id int64) (*models.Article, error)
	RegisterUser(ctx context.Context, user models.User) (err error)
	GetByUsername(ctx context.Context, username string) (result []models.User, err error)
}
