// Package utils ...
package utils

import (
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv is to add the environment variable from the .env file.
func LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		Logger.Panic("Error loading .env file variable")
	}

	return nil
}

// GetEnv is to add the environment variable.
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
