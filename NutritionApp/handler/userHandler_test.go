package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"kunikida123456/NutritionApp/domain/model"
	"kunikida123456/NutritionApp/myerror"
)

// モックのUserUsecase
type mockUserUsecase struct {
	mock.Mock
}

// Signupメソッドのモック
func (m *mockUserUsecase) Signup(name, email, password string) (*model.User, error) {
	if email == "" || password == "" {
		return nil,  &myerror.BadRequestError{Msg: "empty password"}
	}
	args := m.Called(name, email, password)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *mockUserUsecase) Login(email, password string) (string, *model.User, error) {
	args := m.Called(email, password)
	return args.String(0), args.Get(1).(*model.User), args.Error(2)
}

func TestSignup(t *testing.T) {
	// モックのUserUsecaseを作成
	mockUsecase := &mockUserUsecase{}

	// モックの戻り値を設定（成功の場合）
	mockUsecase.On("Signup", "John Doe", "john@example.com", "password123").Return(&model.User{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
		Password: "password123",
	}, nil)

	// テストケース実行
	reqBody := `{
		"name": "John Doe",
		"email": "john@example.com",
		"password": "password123"
	}`
	req := httptest.NewRequest("POST", "/users/signup", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	h := NewUserHandler(mockUsecase)
	err := h.Signup(c)

	// 結果の検証
	assert.NoError(t, err)
	fmt.Println(rec.Code)
	assert.Equal(t, http.StatusCreated, rec.Code)

	expectedResponse := `{"id":1,"name":"John Doe","email":"john@example.com"}`
	assert.JSONEq(t, expectedResponse, rec.Body.String())
}
func TestSignup_EmptyPassword(t *testing.T) {
	// モックのUserUsecaseを作成
	mockUsecase := &mockUserUsecase{}

	mockUsecase.On("Signup", "John Doe", "john@example.com").Return(nil, &myerror.BadRequestError{Msg: "empty password"})

	// テストケース実行
	reqBody := `{
		"name": "John Doe",
		"email": "john@example.com",
		"password": ""
	}`
	req := httptest.NewRequest("POST", "/users/signup", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	h := NewUserHandler(mockUsecase)
	err := h.Signup(c)

	// 結果の検証
	assert.EqualError(t, err, "email and password are required")
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
