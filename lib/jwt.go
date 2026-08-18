package lib

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"alilacream/socialx/models"

	"github.com/golang-jwt/jwt/v5"
)

// to get the userID
type Claim struct {
	PreferredUsername string `json:"preferred_username"`
	Email             string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateJWT creating a new token, to minimize code in login and register
func GenerateJWT(secret string, user *models.User) (string, error) {
	// the format is containing the payload, the header is the first param, containing the method

	claims := Claim{
		PreferredUsername: user.Username,
		Email:             user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.Itoa(int(user.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// the token is of type token, we need it's string value, NOTE: Hmac Meth needs the secret in a slice byte type
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// GetUserIDFromCookie will return -1 if there's an error
func GetUserIDFromCookie(secret string, r *http.Request) (int64, error) {
	tokenStr, err := r.Cookie("jwt")
	if err != nil {
		return -1, err
	}
	token, err := jwt.ParseWithClaims(tokenStr.Value, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return -1, err
	}
	claims, ok := token.Claims.(*Claim)
	if !ok {
		return -1, errors.New("Couldn't parse the claim Token")
	}
	userID, err := strconv.Atoi(claims.RegisteredClaims.Subject)
	if err != nil {
		return -1, err
	}
	return int64(userID), nil
}
