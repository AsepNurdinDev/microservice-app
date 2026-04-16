package config

import (
	"log"
	"os"
)

type Config struct {
	Port      string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	JWTSecret string
}

func Load() *Config {
	cfg := &Config{
		Port:      getEnv("AUTH_PORT", "8001"),
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBUser:    getEnv("DB_USER", "user"),
		DBPass:    getEnv("DB_PASSWORD", "password"),
		DBName:    getEnv("DB_NAME", "auth_db"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	if cfg.JWTSecret == "" {
		log.Fatal("[CONFIG] JWT_SECRET is required and cannot be empty")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}