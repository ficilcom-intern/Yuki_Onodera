package usecase

import (
	"errors"
	"time"

	"kunikida123456/NutritionApp/domain/model"
	"kunikida123456/NutritionApp/domain/repository"
	"kunikida123456/NutritionApp/myerror"
	"kunikida123456/NutritionApp/util"
)

type UserUsecase interface {
	Signup(name, email, password string) (*model.User, error)
	Login(email, password string) (string, *model.User, error)
}

type userUsecase struct {
	repository repository.UserRepository
	timeout    time.Duration
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repository: userRepo,
		timeout:    time.Duration(2) * time.Second,
	}
}

func (uu *userUsecase) Signup(name, email, password string) (*model.User, error) {

	exsitUser, err := uu.repository.GetUserByEmail(email)
	if err != nil {
		return nil, &myerror.InternalServerError{Err: err}
	}
	if exsitUser.ID != 0 {
		return nil, &myerror.BadRequestError{Err: errors.New("user already exists")}
	}

	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, &myerror.InternalServerError{Err: err}
	}

	u := &model.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	user, err := uu.repository.CreateUser(u)
	if err != nil {
		return nil, &myerror.InternalServerError{Err: err}
	}

	return user, nil
}

func (uu *userUsecase) Login(email, password string) (string, *model.User, error) {
	user, err := uu.repository.GetUserByEmail(email)
	if err != nil {
		return "", nil, &myerror.InternalServerError{Err: err}
	}
	if user.ID == 0 {
		return "", nil, &myerror.BadRequestError{Err: errors.New("user is not exist")}
	}

	err = util.CheckPassword(user.Password, password)
	if err != nil {
		return "", nil, &myerror.BadRequestError{Err: errors.New("password is incorrect")}
	}

	signedString, err := util.GenerateSignedString(user.ID, user.Name)
	if err != nil {
		return "", nil, &myerror.InternalServerError{Err: err}
	}

	return signedString, user, nil
}
