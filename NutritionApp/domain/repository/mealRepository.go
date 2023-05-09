package repository

import (
	"kunikida123456/NutritionApp/domain/model"
)

type MealRepository interface {
	FindByID(id int) (*model.Meal, error)
	FindByUser(user_id int) ([]*model.Meal, error)
	Create(meal *model.Meal) (*model.Meal, error)
	Update(meal *model.Meal) (*model.Meal, error)
	Delete(meal *model.Meal) error
}