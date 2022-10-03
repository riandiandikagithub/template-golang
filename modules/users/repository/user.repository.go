package repository

import (
	"context"
	"fmt"
	"payment-simulator/models"
	"payment-simulator/modules/users"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"gorm.io/gorm"
)

type userRepository struct {
	Conn *gorm.DB
}

// // NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewUserRepository(Conn *gorm.DB) users.Repository {
	return &userRepository{Conn}
}

func (m *userRepository) GetByUsername(ctx context.Context, username string) (result []models.User, err error) {
	db := m.Conn.WithContext(ctx)
	db.Debug()
	db.Where("username  = ?", username)
	err = db.Find(&result).Error // query
	if err != nil {
		otelzap.Ctx(ctx).Error("Error running query <GetByUsername> username : " + username)

		return result, err
	}
	return result, nil
}

func (m *userRepository) RegisterUser(ctx context.Context, user models.User) (err error) {
	db := m.Conn.WithContext(ctx)
	db.Debug()
	err = db.Create(&user).Error

	if err != nil {
		otelzap.Ctx(ctx).Error("Error running query <RegisterUser> data : " + fmt.Sprintf("%#v", user))

		return err
	}
	return nil
}
