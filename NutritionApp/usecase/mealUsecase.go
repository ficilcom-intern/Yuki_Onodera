package usecase

import (
	"kunikida123456/NutritionApp/domain/model"
	"kunikida123456/NutritionApp/domain/repository"
	"kunikida123456/NutritionApp/myerror"
)

type MealUsecase interface {
	Create(userID int, memo string, Type string, carbs float64, fat float64, protein float64, calories float64) (*model.Meal, error)
	Update(userID int, id int, memo string, Type string, carbs float64, fat float64, protein float64, calories float64) (*model.Meal, error)
	Delete(userID int, id int) error
	FindAll(userID int) ([]*model.Meal, error)
	FindByID(userID int,id int) (*model.Meal, error)
}

type mealUsecase struct {
	mealRepo repository.MealRepository
}

func NewMealUsecase(mealRepo repository.MealRepository) MealUsecase {
	return &mealUsecase{mealRepo: mealRepo}
}

func (mu *mealUsecase) Create(userID int, memo string, mealType string, carbs float64, fat float64, protein float64, calories float64) (*model.Meal, error) {
	meal, err := model.NewMeal(userID, memo, mealType, carbs, fat, protein, calories)
	if err != nil {
		return nil, &myerror.BadRequestError{Err: err}
	}

	createdMeal, err := mu.mealRepo.Create(meal)
	if err != nil {
		return nil, &myerror.BadRequestError{Err: err}
	}

	return createdMeal, nil
}

func (mu *mealUsecase) Update(userID int, id int, memo string, Type string, carbs float64, fat float64, protein float64, calories float64) (*model.Meal, error) {
	targetMeal, err := mu.mealRepo.FindByID(id, userID)
	if err != nil {
		return nil, &myerror.NotFoundError{Err: err}
	}

	err = targetMeal.Set(userID, memo, Type, carbs, fat, protein, calories)
	if err != nil {
		return nil, &myerror.BadRequestError{Err: err}
	}

	updatedMeal, err := mu.mealRepo.Update(targetMeal)
	if err != nil {
		return nil, &myerror.BadRequestError{Err: err}
	}

	return updatedMeal, nil
}

func (mu *mealUsecase) Delete(userID int, id int) error {
	targetMeal, err := mu.mealRepo.FindByID(userID, id)
	if err != nil {
		return &myerror.NotFoundError{Err: err}
	}

	err = mu.mealRepo.Delete(targetMeal)
	if err != nil {
		return &myerror.BadRequestError{Err: err}
	}
	return nil
}

func (mu *mealUsecase) FindAll(userID int) ([]*model.Meal, error) {
	foundMeals, err := mu.mealRepo.FindAll(userID)
	if err != nil {
		return nil, &myerror.NotFoundError{Err: err}
	}

	return foundMeals, nil
}

func (mu *mealUsecase) FindByID(userID int, id int) (*model.Meal, error) {
	foundMeal, err := mu.mealRepo.FindByID(userID, id)
	if err != nil {
		return nil, &myerror.NotFoundError{Err: err}
	}

	return foundMeal, nil
}
