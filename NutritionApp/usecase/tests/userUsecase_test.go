package usecase_test

import (
	"testing"

	"kunikida123456/NutritionApp/domain/model"
	"kunikida123456/NutritionApp/myerror"
	"kunikida123456/NutritionApp/usecase"
	"kunikida123456/NutritionApp/util"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) GetUserByEmail(email string) (*model.User, error) {
	args := m.Called(email)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *mockUserRepository) CreateUser(user *model.User) (*model.User, error) {
	args := m.Called(user)
	return args.Get(0).(*model.User), args.Error(1)
}

func TestUserUsecase_Signup(t *testing.T) {
	mockUserRepo := &mockUserRepository{}

	validName := "John Doe"
	validEmail := "john@example.com"
	validPassword := "password123"

	createdUser := &model.User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	existingUser := &model.User{
		ID:       1,
		Name:     "Existing User",
		Email:    "existing@example.com",
		Password: "hashedpassword",
	}

	mockUserRepo.On("GetUserByEmail", validEmail).Return(&model.User{}, nil)
	mockUserRepo.On("GetUserByEmail", existingUser.Email).Return(existingUser, nil)
	mockUserRepo.On("CreateUser", mock.AnythingOfType("*model.User")).Return(createdUser, nil)

	uc := usecase.NewUserUsecase(mockUserRepo)

	t.Run("Valid Signup", func(t *testing.T) {
		user, err := uc.Signup(validName, validEmail, validPassword)
		assert.NoError(t, err)
		assert.NotNil(t, user)
	})

	t.Run("User Already Exists", func(t *testing.T) {
		user, err := uc.Signup(existingUser.Name, existingUser.Email, validPassword)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "user already exist", err.(*myerror.BadRequestError).Msg)
	})

	// モックのアサーションを確認
	mockUserRepo.AssertExpectations(t)
}

func TestUserUsecase_Login(t *testing.T) {
	mockUserRepo := &mockUserRepository{}

	t.Run("Valid_Login", func(t *testing.T) {
		userUsecase := usecase.NewUserUsecase(mockUserRepo)

		email := "john@example.com"
		password := "password"

		hashedPassword, err := util.HashPassword(password)
		if err != nil {
			panic(err)
		}

		user := &model.User{ID: 1, Name: "John Doe", Email: email, Password: hashedPassword}

		mockUserRepo.On("GetUserByEmail", email).Return(user, nil)

		token, foundUser, err := userUsecase.Login(email, password)

		assert.NoError(t, err)
		assert.NotNil(t, token)
		assert.NotNil(t, foundUser)
		assert.Equal(t, user, foundUser)

		mockUserRepo.AssertCalled(t, "GetUserByEmail", email)
	})

	t.Run("Invalid_User_Not_Found", func(t *testing.T) {
		userUsecase := usecase.NewUserUsecase(mockUserRepo)

		email := "unknown@example.com"
		password := "password"

		mockUserRepo.On("GetUserByEmail", email).Return(&model.User{}, nil)

		token, foundUser, err := userUsecase.Login(email, password)

		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Nil(t, foundUser)
		assert.Equal(t, &myerror.NotFoundError{Msg: "user not found or password is invalid"}, err)

		mockUserRepo.AssertCalled(t, "GetUserByEmail", email)
	})

	t.Run("Invalid_Password_Mismatch", func(t *testing.T) {
		userUsecase := usecase.NewUserUsecase(mockUserRepo)

		email := "john@example.com"
		password := "wrong_password"

		user := &model.User{ID: 1, Name: "John Doe", Email: email, Password: "password"}

		mockUserRepo.On("GetUserByEmail", email).Return(user, nil)

		token, foundUser, err := userUsecase.Login(email, password)

		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Nil(t, foundUser)

		mockUserRepo.AssertCalled(t, "GetUserByEmail", email)
	})
}
