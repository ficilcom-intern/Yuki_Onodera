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
