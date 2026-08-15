package lib

import (
	"time"

	"alilacream/socialx/models"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT creating a new token, to minimize code in login and register
func GenerateJWT(secret string, user *models.User) (string, error) {
	// the format is containing the payload, the header is the first param, containing the method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":                user.ID, // Use id, to be unique
		"preferred_username": user.Username,
		"email":              user.Email,
		"exp":                time.Now().Add(24 * time.Hour).Unix(),
		"iat":                time.Now().Unix(),
	})
	// the token is of type token, we need it's string value, NOTE: Hmac Meth needs the secret in a slice byte type
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
