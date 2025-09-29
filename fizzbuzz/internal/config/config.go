package config

import (
	"os"
)

type Config struct {
	Port     string
	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string
	RedisURL string
}

func Load() *Config {
	return &Config{
		Port:     getEnv("PORT", "8080"),
		DBHost:   getEnv("DB_HOST", "db"),
		DBPort:   getEnv("DB_PORT", "5432"),
		DBUser:   getEnv("DB_USER", "postgres"),
		DBPass:   getEnv("DB_PASS", "postgres"),
		DBName:   getEnv("DB_NAME", "fizzbuzz"),
		RedisURL: getEnv("REDIS_URL", "cache:6379"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
