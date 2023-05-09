package infra

import (
	"kunikida123456/NutritionApp/domain/model"
	"kunikida123456/NutritionApp/domain/repository"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return &UserRepository{Conn: conn}
}

func (ur *UserRepository) Create(user *model.User) (*model.User, error) {
	if err := ur.Conn.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
