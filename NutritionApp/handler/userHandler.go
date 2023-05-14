package handler

import (
	"net/http"
	"time"

	"kunikida123456/NutritionApp/usecase"

	"github.com/labstack/echo/v4"
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

type userSignUprequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
type userSignupResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Singup godoc
// @ID postUsersSignup
// @Description サインアップ
// @Accept   json
// @Produce  json
// @Router   /users/signup [post]
// @Param    body body userSignupRequest true "サインイン情報"
// @Success  200 {object} userSignupResponse
// @Failure  400 {object} myerror.BadRequestError
// @Failure  401 {object} myerror.UnauthorizedError
// @Failure  404 {object} myerror.NotFoundError
//	@Failure  500 {object} myerror.InternalServerError
func (uh *userHandler) Signup(c echo.Context) error {
	req := new(userSignUprequest)

	if err := c.Bind(&req); err != nil {
		return err
	}

	createdUser, err := uh.userUsecase.Signup(req.Name, req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res := userSignupResponse{
		ID:    createdUser.ID,
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}

	return c.JSON(http.StatusCreated, res)
}

type userLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type userLoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token`
}

// Login godoc
// @ID postUsersLogin
// @Description ログイン
// @Accept   json
// @Produce  json
// @Router   /users/lgin [post]
// @Param    body body userLoginRequest true "ログイン情報"
// @Success  200 {object} userLoginResponse
// @Failure  400 {object} myerror.BadRequestError
// @Failure  401 {object} myerror.UnauthorizedError
// @Failure  404 {object} myerror.NotFoundError
//
//	@Failure  500 {object} myerror.InternalServerError
func (uh *userHandler) Login(c echo.Context) error {
	req := new(userLoginRequest)

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	signedString, User, err := uh.userUsecase.Login(req.Email, req.Password)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	res := userLoginResponse{
		ID:    User.ID,
		Name:  User.Name,
		Token: signedString,
	}

	return c.JSON(http.StatusCreated, res)
}

type userLogoutResponse struct {
	Token string `json:"token`
}

// Logout godoc
// @ID postUsersLogout
// @Description ログアウト
// @Accept   json
// @Produce  json
// @Router   /users/logout [post]
// @Success  200 {object} userLoginResponse
// @Failure  400 {object} myerror.BadRequestError
// @Failure  401 {object} myerror.UnauthorizedError
//	@Failure  500 {object} myerror.InternalServerError
func (uh *userHandler) Logout(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*MyJWTClaims)
	claims["exp"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	res := userLogoutResponse{
		Token: refreshToken,
	}
	return c.JSON(http.StatusOK, res)
}
