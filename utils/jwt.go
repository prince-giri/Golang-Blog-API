package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey = []byte("supersecretkey") // 🔒 you can keep in env file

func GenerateJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(SecretKey)
}
