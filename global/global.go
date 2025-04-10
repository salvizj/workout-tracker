package global

import (
	"database/sql"
)

type Config struct {
	DatabasePath string `env:"DATABASE_PATH"`
	Port         string `env:"PORT"`
	Domain       string `env:"DOMAIN"`
	Env          string `env:"ENV"`
	BindAddress  string `env:"BIND_ADDRESS"`
	JwtSecret    string `env:"JWT_SECRET"`
}

var (
	DB     *sql.DB
	CONFIG Config
)
