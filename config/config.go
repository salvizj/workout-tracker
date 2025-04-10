package config

import (
	"errors"
	"fmt"
	"os"

	"workout_tracker/global"

	"github.com/joho/godotenv"
)

func LoadConfig() (global.Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return global.Config{}, fmt.Errorf("error loading .env file: %w", err)
	}

	cfg := global.Config{
		DatabasePath: os.Getenv("DATABASE_PATH"),
		Port:         os.Getenv("PORT"),
		Domain:       os.Getenv("DOMAIN"),
		Env:          os.Getenv("ENV"),
		BindAddress:  os.Getenv("BIND_ADDRESS"),
		JwtSecret:    os.Getenv("JWT_SECRET"),
	}

	if cfg.DatabasePath == "" {
		return global.Config{}, errors.New("DATABASE_PATH is required")
	}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	if cfg.Domain == "" {
		cfg.Domain = "localhost"
	}
	if cfg.Env == "" {
		cfg.Env = "development"
	}

	return cfg, nil
}
