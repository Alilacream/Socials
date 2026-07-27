package env

import (
	"os"

	"github.com/joho/godotenv"
)

func GetVar(variable string) string {
	err := godotenv.Load()
	if err != nil {
		return ""
	}
	return os.Getenv(variable)
}
