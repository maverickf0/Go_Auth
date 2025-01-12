package utils

import (
	"time"

	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
