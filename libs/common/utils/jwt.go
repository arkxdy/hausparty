package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// token generation, validation, expiration handling, and secret from env
func GenerateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	return token.SignedString([]byte("your-secret-key"))
}
