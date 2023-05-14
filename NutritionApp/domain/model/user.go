package model

import (
	"fmt"
	"time"
	"unicode/utf8"
)

type User struct {
	ID        int        `json:"user_id" gorm:"praimaly_key"`
	Meals     []Meal     `json:"meals" gorm:"foreignKey:UserID"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (u *User) Validate() error {
	if len(u.Name) == 0 {
		return fmt.Errorf("User name is empty")
	}

	if utf8.RuneCountInString(u.Name) > 20 {
		return fmt.Errorf("User name is too long")
	}

	return nil
}
