package util

import (
  "os"
  "strconv"
  "time"

  "github.com/golang-jwt/jwt/v4"
)

type MyJWTClaims struct {
  ID       string `json:"id"`
  Username string `json:"username"`
  jwt.RegisteredClaims
}

func getJWTSecret() []byte {
  return []byte(os.Getenv("JWT_SECRET_KEY"))
}

func GenerateSignedString(userId int64, username string) (string, error) {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
    ID:       strconv.Itoa(int(userId)),
    Username: username,
    RegisteredClaims: jwt.RegisteredClaims{
      Issuer:    strconv.Itoa(int(userId)),
      ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
    },
  })

  return token.SignedString(getJWTSecret())
}
