package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv загружает переменные окружения из файла .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// GetEnv получает значение переменной окружения
func GetEnv(key string) string {
	return os.Getenv(key)
}
