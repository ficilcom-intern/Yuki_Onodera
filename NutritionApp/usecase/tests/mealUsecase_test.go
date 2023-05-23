package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"kunikida123456/NutritionApp/domain/model"
	"kunikida123456/NutritionApp/usecase"
)

type mockMealRepo struct {
	mock.Mock
}

func (m *mockMealRepo) Create(meal *model.Meal) (*model.Meal, error) {
	args := m.Called(meal)
	return args.Get(0).(*model.Meal), args.Error(1)
}

func (m *mockMealRepo) Update(meal *model.Meal) (*model.Meal, error) {
	args := m.Called(meal)
	return args.Get(0).(*model.Meal), args.Error(1)
}

func (m *mockMealRepo) Delete(meal *model.Meal) error {
	args := m.Called(meal)
	return args.Error(0)
}

func (m *mockMealRepo) FindAll(userID int) ([]*model.Meal, error) {
	args := m.Called(userID)
	return args.Get(0).([]*model.Meal), args.Error(1)
}

func (m *mockMealRepo) FindByID(userID int, id int) (*model.Meal, error) {
	args := m.Called(userID, id)
	return args.Get(0).(*model.Meal), args.Error(1)
}

func TestMealUsecase_Create(t *testing.T) {
	mockMealRepo := &mockMealRepo{}
	mealUsecase := usecase.NewMealUsecase(mockMealRepo)

	userID := 1
	memo := "test meal"
	mealType := "breakfast"
	carbs := 20.5
	fat := 10.2
	protein := 15.7
	calories := 250.3

	createdMeal := &model.Meal{
		ID:        1,
		UserID:    userID,
		Memo:      memo,
		MealType:  mealType,
		Carbs:     carbs,
		Fat:       fat,
		Protein:   protein,
		Calories:  calories,
	}

	mockMealRepo.On("Create", mock.AnythingOfType("*model.Meal")).Return(createdMeal, nil)

	result, err := mealUsecase.Create(userID, memo, mealType, carbs, fat, protein, calories)

	assert.NoError(t, err)
	assert.Equal(t, createdMeal, result)

	mockMealRepo.AssertCalled(t, "Create", mock.AnythingOfType("*model.Meal"))
}

func TestMealUsecase_Update(t *testing.T) {
	mockMealRepo := &mockMealRepo{}
	mealUsecase := usecase.NewMealUsecase(mockMealRepo)

	userID := 1
	mealID := 1
	memo := "updated meal"
	mealType := "lunch"
	carbs := 30.7
	fat := 12.5
	protein := 18.2
	calories := 320.9

	targetMeal := &model.Meal{
		ID:        mealID,
		UserID:    userID,
		Memo:      "test meal",
		MealType:  "breakfast",
		Carbs:     20.5,
		Fat:       10.2,
		Protein:   15.7,
		Calories:  250.3,
	}

	updatedMeal := &model.Meal{
		ID:        mealID,
		UserID:    userID,
		Memo:      memo,
		MealType:  mealType,
		Carbs:     carbs,
		Fat:       fat,
		Protein:   protein,
		Calories:  calories,
	}

	mockMealRepo.On("FindByID", mealID, userID).Return(targetMeal, nil)
	mockMealRepo.On("Update", mock.AnythingOfType("*model.Meal")).Return(updatedMeal, nil)

	result, err := mealUsecase.Update(userID, mealID, memo, mealType, carbs, fat, protein, calories)

	assert.NoError(t, err)
	assert.Equal(t, updatedMeal, result)

	mockMealRepo.AssertCalled(t, "FindByID", mealID, userID)
	mockMealRepo.AssertCalled(t, "Update", mock.AnythingOfType("*model.Meal"))
}

func TestMealUsecase_Delete(t *testing.T) {
	mockMealRepo := &mockMealRepo{}
	mealUsecase := usecase.NewMealUsecase(mockMealRepo)

	userID := 1
	mealID := 1

	targetMeal := &model.Meal{
		ID:        mealID,
		UserID:    userID,
		Memo:      "test meal",
		MealType:  "breakfast",
		Carbs:     20.5,
		Fat:       10.2,
		Protein:   15.7,
		Calories:  250.3,
	}

	mockMealRepo.On("FindByID", mealID, userID).Return(targetMeal, nil)
	mockMealRepo.On("Delete", mock.AnythingOfType("*model.Meal")).Return(nil)

	err := mealUsecase.Delete(userID, mealID)

	assert.NoError(t, err)

	mockMealRepo.AssertCalled(t, "FindByID", mealID, userID)
	mockMealRepo.AssertCalled(t, "Delete", mock.AnythingOfType("*model.Meal"))
}

func TestMealUsecase_FindAll(t *testing.T) {
	mockMealRepo := &mockMealRepo{}
	mealUsecase := usecase.NewMealUsecase(mockMealRepo)

	userID := 1

	foundMeals := []*model.Meal{
		{
			ID:        1,
			UserID:    userID,
			Memo:      "meal 1",
			MealType:  "breakfast",
			Carbs:     20.5,
			Fat:       10.2,
			Protein:   15.7,
			Calories:  250.3,
		},
		{
			ID:        2,
			UserID:    userID,
			Memo:      "meal 2",
			MealType:  "lunch",
			Carbs:     30.7,
			Fat:       12.5,
			Protein:   18.2,
			Calories:  320.9,
		},
	}

	mockMealRepo.On("FindAll", userID).Return(foundMeals, nil)

	result, err := mealUsecase.FindAll(userID)

	assert.NoError(t, err)
	assert.Equal(t, foundMeals, result)

	mockMealRepo.AssertCalled(t, "FindAll", userID)
}

func TestMealUsecase_FindByID(t *testing.T) {
	mockMealRepo := &mockMealRepo{}
	mealUsecase := usecase.NewMealUsecase(mockMealRepo)

	userID := 1
	mealID := 1

	foundMeal := &model.Meal{
		ID:        mealID,
		UserID:    userID,
		Memo:      "test meal",
		MealType:  "breakfast",
		Carbs:     20.5,
		Fat:       10.2,
		Protein:   15.7,
		Calories:  250.3,
	}

	mockMealRepo.On("FindByID", mealID, userID).Return(foundMeal, nil)

	result, err := mealUsecase.FindByID(userID, mealID)

	assert.NoError(t, err)
	assert.Equal(t, foundMeal, result)

	mockMealRepo.AssertCalled(t, "FindByID", mealID, userID)
}
