package usecase

import (
	"kunikida123456/NutritionApp/domain/model"
	"kunikida123456/NutritionApp/domain/repository"
)

type MealUsecase interface {
	Create(memo string, Type string, carbs float64, fat float64, protein float64, calories float64) (*model.Meal, error)
	Update(id int, memo string, Type string, carbs float64, fat float64, protein float64, calories float64) (*model.Meal, error)
	Delete(id int) error
	// GetAll(user_id int) ([]*model.Meal, error)
	FindByID(id int) (*model.Meal, error)
}

type mealUsecase struct {
	mealRepo repository.MealRepository
}

func NewMealUsecase(mealRepo repository.MealRepository) MealUsecase {
	return &mealUsecase{mealRepo: mealRepo}
}

func (mu *mealUsecase) Create(memo string, mealType string, carbs float64, fat float64, protein float64, calories float64) (*model.Meal, error) {
	meal, err := model.NewMeal(memo, mealType, carbs, fat, protein, calories)
	if err != nil {
		return nil, err
	}

	createdMeal, err := mu.mealRepo.Create(meal)
	if err != nil {
		return nil, err
	}

	return createdMeal, nil
}

func (mu *mealUsecase) Update(id int, memo string, Type string, carbs float64, fat float64, protein float64, calories float64) (*model.Meal, error) {
	targetMeal, err := mu.mealRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = targetMeal.Set(memo, Type, carbs, fat, protein, calories)
	if err != nil {
		return nil, err
	}

	updatedMeal, err := mu.mealRepo.Update(targetMeal)
	if err != nil {
		return nil, err
	}

	return updatedMeal, nil
}

func (mu *mealUsecase) Delete(id int) error {
	targetMeal, err := mu.mealRepo.FindByID(id)
	if err != nil {
		return err
	}

	err = mu.mealRepo.Delete(targetMeal)
	if err != nil {
		return err
	}
	return nil
}

// func (mu *mealUsecase) GetAll(user_id int) ([]*model.Meal, error) {
// 	foundMeal, err := mu.mealRepo.FindByUser(user_id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return foundMeal, nil
// }

func (mu *mealUsecase) FindByID(id int) (*model.Meal, error) {
	foundMeal, err := mu.mealRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundMeal, nil
}
