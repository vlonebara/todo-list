package config

import (
	"fmt"
	"os"
)

// Config содержит настройки, которые читаются один раз при запуске.
type Config struct {
	Port      string
	DBFile    string
	Password  string
	JWTSecret string
}

// Load загружает .env и возвращает готовую конфигурацию приложения.
func Load() (Config, error) {
	if err := loadEnvFile(".env"); err != nil {
		return Config{}, fmt.Errorf("load .env: %w", err)
	}
	return Config{
		Port:      envOrDefault("TODO_PORT", "7540"),
		DBFile:    envOrDefault("TODO_DBFILE", "scheduler.db"),
		Password:  os.Getenv("TODO_PASSWORD"),
		JWTSecret: os.Getenv("TODO_JWT_SECRET"),
	}, nil
}

// envOrDefault возвращает значение переменной или указанное значение по умолчанию.
func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
