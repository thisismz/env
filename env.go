package env

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	env       map[string]string
	setupOnce sync.Once
)

func GetEnv(key, def string) string {
	setupOnce.Do(setupEnvFile)
	if val, ok := env[key]; ok {
		return val
	}
	return getOsEnv(key, def)
}
func GetEnvAsInt(key string, def int) int {
	strVal := GetEnv(key, "")
	if strVal == "" {
		return def
	}
	intVal, err := strconv.Atoi(strVal)
	if err != nil {
		log.Printf("Warning: could not convert key %s value to int, defaulting to 0: %v\n", key, err)
		return def
	}
	return intVal
}

func getOsEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func setupEnvFile() {
	envFile := ".env"
	fileEnv, err := godotenv.Read(envFile)
	if err != nil {
		log.Printf("Warning: could not read env file : %s error: %v\n", envFile, err)
		return
	}
	env = fileEnv
}

func IsDevelopment() bool {
	return GetEnv("APP_ENV", "development") == "development"
}
