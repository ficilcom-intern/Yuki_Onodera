package handler

import (
	"net/http"

	"kunikida123456/NutritionApp/usecase"

	"github.com/labstack/echo/v4"
)

type ProfileHandler struct {
	profileUsecase usecase.ProfileUsecase
}

func NewProfileHandler(profileUsecase usecase.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{
		profileUsecase: profileUsecase,
	}
}

type calculateBMIRequest struct {
	height float64 `json:"height"`
	weight float64 `json:"weight"`
}

type calculateBMIResponse struct {
	bmi float64 `json:"bmi"`
}

func (h *ProfileHandler) CalculateBMI(c echo.Context) error {
	req := new(calculateBMIRequest)
	if err := c.Bind(&req); err != nil {
		return err
	}

	bmi, err := h.profileUsecase.CalculateBMI(req.height, req.weight)
	if err != nil {
		return err
	}

	res := calculateBMIResponse{
		bmi: bmi,
	}

	return c.JSON(http.StatusOK, res)
}

type calculateDailyNutritionsRequest struct {
	height float64 `json:"height"`
	weight float64 `json:"weight"`
	age    int     `json:"age"`
}

type calculateDailyNutritionsResponse struct {
	carbohydrates float64 `json:"carbohydrates"`
	protein       float64 `json:"protein"`
	fat           float64 `json:"fat"`
	bmr           float64 `json:"bmr"`
}

func (h *ProfileHandler) CalculateDailyNutritions(c echo.Context) error {
	req := new(calculateDailyNutritionsRequest)
	if err := c.Bind(&req); err != nil {
		return err
	}

	nutritions, err := h.profileUsecase.CalculateDailyNutritions(req.height, req.weight, req.age)
	if err != nil {
		return err
	}

	res := calculateDailyNutritionsResponse{
		carbohydrates: nutritions.Carbohydrates,
		protein:       nutritions.Protein,
		fat:           nutritions.Fat,
		bmr:           nutritions.BMR,
	}

	return c.JSON(http.StatusOK, res)
}
