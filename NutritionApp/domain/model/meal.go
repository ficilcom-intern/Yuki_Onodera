package model

import (
	"errors"
	"time"
)

type Meal struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	UserID    int        `json:"user_id" gorm:"references:UserID"`
	Memo      string     `json:"memo" gorm:"default:null"`
	MealType  string     `json:"type" gorm:"default:null"`
	Carbs     float64    `json:"carbs" gorm:"default:null"`
	Fat       float64    `json:"fat" gorm:"default:null"`
	Protein   float64    `json:"protein" gorm:"default:null"`
	Calories  float64    `json:"calories" gorm:"default:null"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func NewMeal(memo string, mealType string, carbs float64, fat float64, protein float64, calories float64) (*Meal, error) {
	if mealType == "" {
		return nil, errors.New("食事の種類を入力してください")
	}

	meal := &Meal{
		Memo:     memo,
		MealType: mealType,
		Carbs:    carbs,
		Fat:      fat,
		Protein:  protein,
		Calories: calories,
	}
	return meal, nil
}

func (m *Meal) Set(memo string, mealType string, carbs float64, fat float64, protein float64, calories float64) error {
	if mealType == "" {
		return errors.New("食事の種類を入力してください")
	}

	m.Memo = memo
	m.MealType = mealType
	m.Carbs = carbs
	m.Fat = fat
	m.Protein = protein
	m.Calories = calories

	return nil
}
