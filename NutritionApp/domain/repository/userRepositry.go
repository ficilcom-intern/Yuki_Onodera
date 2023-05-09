package repository

import (
	"kunikida123456/NutritionApp/domain/model"
)

type UserRepository interface {
	Create(meal *model.User) (*model.User, error)
}
