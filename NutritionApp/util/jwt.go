package util

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/subosito/gotenv"
)	

type MyJWTClaims struct {
	ID   int
	Name string
	jwt.RegisteredClaims
}

var Config = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(MyJWTClaims)
	},
	SigningKey: getJWTSecret(),
}

func getJWTSecret() []byte {
	err := gotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return []byte(os.Getenv("JWT_SECRET_KEY"))
}

func GenerateSignedString(userId int, name string) (string, error) {
	claims := &MyJWTClaims{
		userId,
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJWTSecret())
}

func UserIDFromToken(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*MyJWTClaims)
	return claims.ID
}
