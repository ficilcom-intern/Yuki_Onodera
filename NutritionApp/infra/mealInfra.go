package infra

import (
	"kunikida123456/NutritionApp/domain/model"
	"kunikida123456/NutritionApp/domain/repository"

	"gorm.io/gorm"
)

type MealRepository struct {
	Conn *gorm.DB
}

func NewMealRepository(conn *gorm.DB) repository.MealRepository {
	return &MealRepository{Conn: conn}
}

func (mr *MealRepository) Create(meal *model.Meal) (*model.Meal, error) {
	if err := mr.Conn.Create(&meal).Error; err != nil {
		return nil, err
	}
	return meal, nil
}

// Update mealの更新
func (mr *MealRepository) Update(meal *model.Meal) (*model.Meal, error) {
	if err := mr.Conn.Save(&meal).Error; err != nil {
		return nil, err
	}

	return meal, nil
}

// Delete mealの削除
func (mr *MealRepository) Delete(meal *model.Meal) error {
	if err := mr.Conn.Delete(&meal).Error; err != nil {
		return err
	}

	return nil
}

func (mr *MealRepository) FindByID(id int, uid int) (*model.Meal, error) {
	meal := &model.Meal{ID: id, UserID: uid}
	if err := mr.Conn.First(&meal).Error; err != nil {
		return nil, err
	}
	return meal, nil
}
