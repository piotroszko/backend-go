package env

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	godotenv.Load(".env")
	isJwtSecretSet := os.Getenv("JWT_SECRET")

	if isJwtSecretSet == "" {
		panic("JWT_SECRET is not set")
	}
}
