package usecase

import (
	"math"

	"kunikida123456/NutritionApp/myerror"
)

type ProfileUsecase interface {
	CalculateBMI(height float64, weight float64) (float64, error)
	CalculateDailyNutritions(height float64, weight float64, age int) (*DailyNutritions, error)
}

type profileUsecase struct{}

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

type DailyNutritions struct {
	Carbohydrates float64
	Fat           float64
	Protein       float64
	BMR           float64
}

func (pu *profileUsecase) CalculateDailyNutritions(height float64, weight float64, age int) (*DailyNutritions, error) {
	if weight <= 0 {
		return &DailyNutritions{}, &myerror.BadRequestError{Msg: "Invalid weight"}
	}

	if height <= 0 {
		return &DailyNutritions{}, &myerror.BadRequestError{Msg: "Invalid height"}
	}

	if age <= 0 {
		return &DailyNutritions{}, &myerror.BadRequestError{Msg: "Invalid age"}
	}

	// 基礎代謝の計算（Harris-Benedict 方程式を使用）
	bmr := 0.0

	if age < 18 {
		bmr = 88.362 + (13.397 * weight) + (4.799 * height) - (5.677 * float64(age))
	} else {
		bmr = 447.593 + (9.247 * weight) + (3.098 * height) - (4.330 * float64(age))
	}

	// 糖質の目安量の計算
	carbohydrates := 0.45 * bmr / 4 // 総カロリー摂取量の45％が糖質から供給されると仮定

	// 脂質の目安量の計算
	fat := 0.25 * bmr / 9 // 総カロリー摂取量の25％が脂質から供給されると仮定

	// タンパク質の目安量の計算
	protein := 0.15 * bmr / 4 // 総カロリー摂取量の15％がタンパク質から供給されると仮定

	dailyNutritions := &DailyNutritions{
		Carbohydrates: carbohydrates,
		Fat:           fat,
		Protein:       protein,
		BMR:           bmr,
	}

	return dailyNutritions, nil
}
