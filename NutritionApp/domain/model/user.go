package model

import (
	"time"
)

type User struct {
	ID        int       `json:"user_id" gorm:"praimaly_key"`
	Meals     []Meal    `json:"meals" gorm:"foreignKey:UserID"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
