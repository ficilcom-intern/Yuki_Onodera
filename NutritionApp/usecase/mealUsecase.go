package usecase

import (
	"kunikida123456/NutritionApp/domain/model"
	"kunikida123456/NutritionApp/domain/repository"
	"kunikida123456/NutritionApp/util"

	"github.com/labstack/echo/v4"
)

type MealUsecase interface {
	Create(userID int, memo string, Type string, carbs float64, fat float64, protein float64, calories float64) (*model.Meal, error)
	Update(c echo.Context, id int, memo string, Type string, carbs float64, fat float64, protein float64, calories float64) (*model.Meal, error)
	Delete(c echo.Context, id int) error
	FindAll(c echo.Context) ([]*model.Meal, error)
	FindByID(c echo.Context, id int) (*model.Meal, error)
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
		return nil, err
	}

	createdMeal, err := mu.mealRepo.Create(meal)
	if err != nil {
		return nil, err
	}

	return createdMeal, nil
}

func (mu *mealUsecase) Update(c echo.Context, id int, memo string, Type string, carbs float64, fat float64, protein float64, calories float64) (*model.Meal, error) {
	uid := util.UserIDFromToken(c)
	targetMeal, err := mu.mealRepo.FindByID(id, uid)
	if err != nil {
		return nil, err
	}

	err = targetMeal.Set(uid, memo, Type, carbs, fat, protein, calories)
	if err != nil {
		return nil, err
	}

	updatedMeal, err := mu.mealRepo.Update(targetMeal)
	if err != nil {
		return nil, err
	}

	return updatedMeal, nil
}

func (mu *mealUsecase) Delete(c echo.Context, id int) error {
	uid := util.UserIDFromToken(c)

	targetMeal, err := mu.mealRepo.FindByID(id, uid)
	if err != nil {
		return err
	}

	err = mu.mealRepo.Delete(targetMeal)
	if err != nil {
		return err
	}
	return nil
}

func (mu *mealUsecase) FindAll(c echo.Context) ([]*model.Meal, error) {
	uid := util.UserIDFromToken(c)
	foundMeal, err := mu.mealRepo.FindAll(uid)
	if err != nil {
		return nil, err
	}

	return foundMeal, nil
}

func (mu *mealUsecase) FindByID(c echo.Context, id int) (*model.Meal, error) {
	uid := util.UserIDFromToken(c)

	foundMeal, err := mu.mealRepo.FindByID(id, uid)
	if err != nil {
		return nil, err
	}

	return foundMeal, nil
}
