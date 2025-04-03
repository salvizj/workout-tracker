package config

import (
	"os"
)

type Config struct {
	DATABASE_URL string
	PORT         string
}

func LoadConfig() Config {
	return Config{
		DATABASE_URL: os.Getenv("DATABASE_URL"),
		PORT:         os.Getenv("PORT"),
	}
}
