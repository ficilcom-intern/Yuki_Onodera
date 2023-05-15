package util

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password is empty")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPassword(hashedPassword string, inputPassword string) error {
	if inputPassword == "" {
		return errors.New("password is empty")
	}
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}
