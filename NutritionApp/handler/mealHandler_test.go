package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"kunikida123456/NutritionApp/domain/model"
	"kunikida123456/NutritionApp/util"
)

type mockMealUsecase struct {
	mock.Mock
}

func (m *mockMealUsecase) Create(userId int, memo string, mealType string, carbs, fat, protein, calories float64) (*model.Meal, error) {
	args := m.Called(userId, memo, mealType, carbs, fat, protein, calories)
	return args.Get(0).(*model.Meal), args.Error(1)
}

func (m *mockMealUsecase) FindByID(userId int, id int) (*model.Meal, error) {
	args := m.Called(userId, id)
	return args.Get(0).(*model.Meal), args.Error(1)
}

func (m *mockMealUsecase) Update(userId int, id int, memo string, mealType string, carbs, fat, protein, calories float64) (*model.Meal, error) {
	args := m.Called(userId, id, memo, mealType, carbs, fat, protein, calories)
	return args.Get(0).(*model.Meal), args.Error(1)
}

func (m *mockMealUsecase) Delete(userId int, id int) error {
	args := m.Called(userId, id)
	return args.Error(0)
}

func (m *mockMealUsecase) FindAll(userId int) ([]*model.Meal, error) {
	args := m.Called(userId)
	return args.Get(0).([]*model.Meal), args.Error(1)
}

func generateTokenForTest() (token *jwt.Token) {
	// トークンの生成
	claims := &util.MyJWTClaims{
		1,
		"onodera",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

func TestPostMeal(t *testing.T) {
	mockUsecase := &mockMealUsecase{}
	mockUsecase.On("Create", 1, "Test Meal", "Breakfast", 30.5, 15.2, 20.3, 350.7).Return(&model.Meal{
		ID:       1,
		Memo:     "Test Meal",
		MealType: "Breakfast",
		Carbs:    30.5,
		Fat:      15.2,
		Protein:  20.3,
		Calories: 350.7,
	}, nil)

	// テストリクエストの作成
	reqBody := `{
        "memo": "Test Meal",
        "mealtype": "Breakfast",
        "carbs": 30.5,
        "fat": 15.2,
        "protein": 20.3,
        "calories": 350.7
    }`
	token := generateTokenForTest()

	req := httptest.NewRequest(http.MethodPost, "/meals", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	c.Set("user", token)

	h := NewMealHandler(mockUsecase)

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
	err := h.Post(c)

	// レスポンスのアサーション
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.JSONEq(t, expectedRes, rec.Body.String())
}

func TestGetMeal(t *testing.T) {
	mockUsecase := &mockMealUsecase{}
	mockMeal := &model.Meal{
		ID:       1,
		Memo:     "Test Meal",
		MealType: "Breakfast",
		Carbs:    30.5,
		Fat:      15.2,
		Protein:  20.3,
		Calories: 350.7,
	}
	mockUsecase.On("FindByID", 1, 1).Return(mockMeal, nil)

	token := generateTokenForTest()

	// テストリクエストの作成
	req := httptest.NewRequest(http.MethodGet, "/meals/1", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetPath("/meals/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	c.Set("user", token)

	h := NewMealHandler(mockUsecase)

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
	err := h.Get(c)

	// レスポンスのアサーション
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, expectedRes, rec.Body.String())
}

func TestPutMeal(t *testing.T) {
	mockUsecase := &mockMealUsecase{}
	mockMeal := &model.Meal{
		ID:       1,
		Memo:     "Test Meal",
		MealType: "Breakfast",
		Carbs:    30.5,
		Fat:      15.2,
		Protein:  20.3,
		Calories: 350.7,
	}
	mockUsecase.On("Update", 1, 1, "Put Meal", "Lunch", 40.0, 20.0, 30.0, 500.0).Return(mockMeal, nil)

	token := generateTokenForTest()

	// テストリクエストの作成
	reqBody := `{
        "memo": "Put Meal",
        "mealtype": "Lunch",
        "carbs": 40.0,
        "fat": 20.0,
        "protein": 30.0,
        "calories": 500.0
    }`
	req := httptest.NewRequest(http.MethodPut, "/meals/1", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetPath("/meals/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	c.Set("user", token)

	h := NewMealHandler(mockUsecase)

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
	err := h.Put(c)

	// レスポンスのアサーション
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, expectedRes, rec.Body.String())
}

func TestDeleteMeal(t *testing.T) {
	mockUsecase := &mockMealUsecase{}
	mockUsecase.On("Delete", 1, 1).Return(nil)

	token := generateTokenForTest()

	// テストリクエストの作成
	req := httptest.NewRequest(http.MethodDelete, "/meals/1", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetPath("/meals/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	c.Set("user", token)

	h := NewMealHandler(mockUsecase)

	err := h.Delete(c)

	// レスポンスのアサーション
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestGetAllMeals(t *testing.T) {
	mockUsecase := &mockMealUsecase{}
	mockMeals := []*model.Meal{
		{
			ID:       1,
			Memo:     "Meal 1",
			MealType: "Breakfast",
			Carbs:    30.5,
			Fat:      15.2,
			Protein:  20.3,
			Calories: 350.7,
		},
		{
			ID:       2,
			Memo:     "Meal 2",
			MealType: "Lunch",
			Carbs:    40.5,
			Fat:      10.2,
			Protein:  25.3,
			Calories: 450.7,
		},
	}
	mockUsecase.On("FindAll", 1).Return(mockMeals, nil)

	token := generateTokenForTest()

	// テストリクエストの作成
	req := httptest.NewRequest(http.MethodGet, "/meals", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	c.Set("user", token)

	h := NewMealHandler(mockUsecase)

	// レスポンスボディのアサーション
	expectedRes := `[
		{
			"id": 1,
			"memo": "Meal 1",
			"mealtype": "Breakfast",
			"carbs": 30.5,
			"fat": 15.2,
			"protein": 20.3,
			"calories": 350.7
		},
		{
			"id": 2,
			"memo": "Meal 2",
			"mealtype": "Lunch",
			"carbs": 40.5,
			"fat": 10.2,
			"protein": 25.3,
			"calories": 450.7
		}
	]`
	err := h.GetAll(c)

	// レスポンスのアサーション
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, expectedRes, rec.Body.String())
}
