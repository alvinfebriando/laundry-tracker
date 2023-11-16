package util

import (
	"os"

	"github.com/joho/godotenv"
)

func Getenv(key string) string {
	err := godotenv.Load()
	if err != nil {
		return ""
	}
	return os.Getenv(key)
}
