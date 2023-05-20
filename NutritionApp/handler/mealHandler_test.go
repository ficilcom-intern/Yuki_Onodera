package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"kunikida123456/NutritionApp/domain/model"
)

type mockMealUsecase struct {
    mock.Mock
}

func (m *mockMealUsecase) Create(userID int, memo string, mealType string, carbs, fat, protein, calories float64) (*model.Meal, error) {
    args := m.Called(userID, memo, mealType, carbs, fat, protein, calories)
    return args.Get(0).(*model.Meal), args.Error(1)
}

func (m *mockMealUsecase) FindByID(id int) (*model.Meal, error) {
    args := m.Called(id)
    return args.Get(0).(*model.Meal), args.Error(1)
}

func (m *mockMealUsecase) Update(id int, memo string, mealType string, carbs, fat, protein, calories float64) (*model.Meal, error) {
    args := m.Called(id, memo, mealType, carbs, fat, protein, calories)
    return args.Get(0).(*model.Meal), args.Error(1)
}

func (m *mockMealUsecase) Delete(id int) error {
    args := m.Called(id)
    return args.Error(0)
}

func (m *mockMealUsecase) FindAll() ([]*model.Meal, error) {
    args := m.Called()
    return args.Get(0).([]*model.Meal), args.Error(1)
}

func TestPostMeal(t *testing.T) {
	// テスト対象のハンドラを作成
	// モックの MealUsecase を渡して NewMealHandler を呼び出す
	mockUsecase := &mockMealUsecase{}
	mockUsecase.On("Create", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
        Return(&model.Meal{}, nil)
	handler := NewMealHandler(mockUsecase)

	// テスト用の Echo インスタンスを作成
	e := echo.New()

	// テストリクエストの作成
	reqBody := `{
		"memo": "Test Meal",
		"mealtype": "Breakfast",
		"carbs": 30.5,
		"fat": 15.2,
		"protein": 20.3,
		"calories": 350.7
	}`
	req := httptest.NewRequest(http.MethodPost, "/meals", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// テストハンドラの実行
	err := handler.Post(c)

	// レスポンスのアサーション
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	// レスポンスボディのアサーション
	expectedRes := `{
		"id": 1,
		"memo": "Test Meal",
		"mealtype": "Breakfast",
		"carbs": 30.5,
		"fat": 15.2,
		"protein": 20.3,
		"calories": 350.7
	}`
	assert.JSONEq(t, expectedRes, rec.Body.String())
}
