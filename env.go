package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var env map[string]string

func GetEnv(key, def string) string {

	if val, ok := env[key]; ok {
		return val
	}
	return getOsEnv(key, def)
}

func getOsEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func SetupEnvFile() {
	envFile := ".env"
	env, _ = godotenv.Read(envFile)
}

func IsDevelopment() bool {
	return GetEnv("APP_ENV", "development") == "development"
}
