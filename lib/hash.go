package lib

import (
	"golang.org/x/crypto/bcrypt" //nolint
)

// HashPassword hashing the password (power to the 5)
func HashPassword(plainText string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(plainText), 5)
	return string(hashedBytes), err
}

// CheckPassword checks the equality between the password and the hashed version
func CheckPassword(hashedPass, password string) bool {
	// NOTE: i miss-placed the varialbes
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password)) == nil
}
