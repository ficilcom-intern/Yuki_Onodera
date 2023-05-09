package model

import (
    "time"
)

type User struct {
	ID   int64  `json:"user_id" gorm:"praimaly_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Height   int    `json:"height"`
	Weight   int    `json:"weight"`
    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"` 
}
