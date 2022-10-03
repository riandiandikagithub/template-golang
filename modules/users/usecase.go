package users

import (
	"context"
	"payment-simulator/models"
)

type Usecase interface {
	RegisterUser(ctx context.Context, data models.User) (err error)
	GetByUsername(ctx context.Context, username string) (data []models.User, err error)
}
