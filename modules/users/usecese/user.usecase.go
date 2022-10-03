package usecase

import (
	"context"
	"payment-simulator/models"
	"payment-simulator/modules/users"
)

type userUsecase struct {
	// articleRepo    article.Repository
	userRepo users.Repository
	// contextTimeout time.Duration
}

// // NewArticleUsecase will create new an articleUsecase object representation of article.Usecase interface
func NewUserUsecase(a users.Repository) users.Usecase {
	return &userUsecase{
		userRepo: a,
	}
}

func (u *userUsecase) RegisterUser(ctx context.Context, data models.User) (err error) {

	err = u.userRepo.RegisterUser(ctx, data)

	return
}

func (u *userUsecase) GetByUsername(ctx context.Context, username string) (data []models.User, err error) {

	data, err = u.userRepo.GetByUsername(ctx, username)

	return
}

// Health
