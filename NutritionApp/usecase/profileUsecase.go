package usecase

import (
	"math"

	"kunikida123456/NutritionApp/myerror"
)

type ProfileUsecase interface {
	CalculateBMI(height float64, weight float64) (float64, error)
}

type profileUsecase struct {
}

func NewProfileUsecase() ProfileUsecase {
	return &profileUsecase{}
}

func (pu *profileUsecase) CalculateBMI(weight float64, height float64) (float64, error) {
	// BMIの計算
	if weight <= 0 {
		return 0, &myerror.BadRequestError{Msg: "Invalid weight"}
	}

	if height <= 0 {
		return 0, &myerror.BadRequestError{Msg: "Invalid height"}
	}

	bmi := weight / math.Pow(height, 2)
	return math.Round(bmi*10) / 10, nil
}
