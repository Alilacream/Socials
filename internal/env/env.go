package env

import (
	"os"

	"github.com/joho/godotenv"
)

func GetVar(variable string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	return os.Getenv(variable), nil
}
