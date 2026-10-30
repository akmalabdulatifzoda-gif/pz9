package config

import (
	"os"
	"strconv"
)

// Config — настройки приложения из переменных окружения.
type Config struct {
	DB_DSN     string // строка подключения к PostgreSQL
	BcryptCost int    // сложность bcrypt, по умолчанию 12
	Addr       string // адрес HTTP-сервера
}

func Load() Config {
	cost := 12
	if v := os.Getenv("BCRYPT_COST"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n >= 4 && n <= 31 {
			cost = n
		}
	}
	addr := os.Getenv("APP_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	return Config{
		DB_DSN:     os.Getenv("DB_DSN"), // postgres://user:pass@localhost:5432/pz9?sslmode=disable
		BcryptCost: cost,
		Addr:       addr,
	}
}
