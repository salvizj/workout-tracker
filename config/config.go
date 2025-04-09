package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabasePath string `env:"DATABASE_PATH"`
	Port         string `env:"PORT"`
	Domain       string `env:"DOMAIN"`
	Env          string `env:"ENV"`
	BindAddress  string `env:"BIND_ADDRESS"`
}

func LoadConfig() (Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return Config{}, fmt.Errorf("error loading .env file: %w", err)
	}

	cfg := Config{
		DatabasePath: os.Getenv("DATABASE_PATH"),
		Port:         os.Getenv("PORT"),
		Domain:       os.Getenv("DOMAIN"),
		Env:          os.Getenv("ENV"),
		BindAddress:  os.Getenv("BIND_ADDRESS"),
	}

	if cfg.DatabasePath == "" {
		return Config{}, errors.New("DATABASE_PATH is required")
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
