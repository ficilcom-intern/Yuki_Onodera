package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"kunikida123456/NutritionApp/domain/model"
)

type mockUserUsecase struct{}

func (m *mockUserUsecase) Signup(name, email, password string) (*model.User, error) {
	// 仮のユーザーを作成
	user := &model.User{
		ID:    1,
		Name:  name,
		Email: email,
	}

	// エラーは nil を返す
	return user, nil
}

func (m *mockUserUsecase) Login(email, password string) (string, *model.User, error) {
	// 仮のユーザーを作成
	user := &model.User{
		ID:    1,
		Name:  "John Doe",
		Email: email,
	}

	// エラーは nil を返す
	return "dummy_token", user, nil
}

func TestSignup(t *testing.T) {
	// モックのUserUsecaseを作成
	mockUsecase := &mockUserUsecase{}

	// テストケース実行
	req := httptest.NewRequest(http.MethodPost, "/users/signup", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	h := NewUserHandler(mockUsecase)
	err := h.Signup(c)

	// 結果の検証
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	// レスポンスの検証
	// ...
}
