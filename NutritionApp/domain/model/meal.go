package model

import (
	"errors"
	"time"
)

type Meal struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	User      User       `json:"user_id" gorm:"references:UserID"`
	Memo      string     `json:"memo" gorm:"default:null"`
	Type      string     `json:"type" gorm:"default:null"`
	Carbs     float64    `json:"carbs" gorm:"default:null"`
	Fat       float64    `json:"fat" gorm:"default:null"`
	Protein   float64    `json:"protein" gorm:"default:null"`
	Calories  float64    `json:"calories" gorm:"default:null"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func NewMeal(memo string, Type string, carbs float64, fat float64, protein float64, calories float64) (*Meal, error) {
	if Type == "" {
		return nil, errors.New("食事の種類を入力してください")
	}

	meal := &Meal{
		Memo:     memo,
		Type:     Type,
		Carbs:    carbs,
		Fat:      fat,
		Protein:  protein,
		Calories: calories,
	}
	return meal, nil
}

func (m *Meal) Set(memo string, Type string, carbs float64, fat float64, protein float64, calories float64) error {
	if Type == "" {
		return errors.New("食事の種類を入力してください")
	}

	m.Memo = memo
	m.Type = Type
	m.Carbs = carbs
	m.Fat = fat
	m.Protein = protein
	m.Calories = calories

	return nil
}
