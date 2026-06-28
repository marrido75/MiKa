package config

import (
	"log"
	"os"
)

type Config struct {
	DBPath    string
	JWTSecret string
	ServerPort string
}

func Load() *Config {
	cfg := &Config{
		DBPath:     getEnv("DB_PATH", "mika.db"),
		JWTSecret:  getEnv("JWT_SECRET", ""),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
