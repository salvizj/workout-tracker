package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_URL string
	PORT         string
}

func LoadConfig() (Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return Config{}, fmt.Errorf("Error loading .env file: %v", err)
	}

	dbURL := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	if dbURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is not set")
	}

	if port == "" {
		return Config{}, fmt.Errorf("PORT is not set")
	}

	return Config{
		DATABASE_URL: dbURL,
		PORT:         port,
	}, nil
}
