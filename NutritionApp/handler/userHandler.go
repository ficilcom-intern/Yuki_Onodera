package handler

import (
	"fmt"
	"net/http"
	"time"

	"kunikida123456/NutritionApp/usecase"

	"github.com/labstack/echo"
)

type UserHandler interface {
	Signup(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}

func (uh *userHandler) Signup(c echo.Context) error {
	type (
		request struct {
			Name     string `json:"name" binding:"required"`
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required,min=8"`
		}
		response struct {
			ID    int64  `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
		}
	)
	req := new(request)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	createdUser, err := uh.userUsecase.Signup(req.Name, req.Email, req.Password)
	fmt.Println(createdUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res := response{
		ID:    createdUser.ID,
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}

	return c.JSON(http.StatusCreated, res)
}

func (uh *userHandler) Login(c echo.Context) error {
	type (
		request struct {
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required"`
		}
		response struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		}
	)

	req := new(request)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	signedString, User, err := uh.userUsecase.Login(req.Email, req.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    signedString,
		Path:     "/",
		Domain:   "localhost",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   false,
	}
	c.SetCookie(cookie)

	res := response{
		ID:   User.ID,
		Name: User.Name,
	}

	return c.JSON(http.StatusCreated, res)
}

func (uh *userHandler) Logout(c echo.Context) error {
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, "logout")
}
