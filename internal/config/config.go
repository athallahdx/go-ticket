package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	BaseURL   string
	Port      string
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	JWTSecret string
}

func getEnv(key, defaultValue string) string {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		return defaultValue
	}
	return val
}

func validate(cfg *Config) {
	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}

	if cfg.DBHost == "" || cfg.DBUser == "" || cfg.DBName == "" {
		log.Fatal("Database configuration is incomplete")
	}
}

func LoadConfig() *Config {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: .env file not found, using OS environment variables")
		}
	}

	cfg := &Config{
		BaseURL:   getEnv("BASE_URL", "http://localhost:8080"),
		Port:      getEnv("PORT", "8080"),
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBUser:    getEnv("DB_USER", "root"),
		DBPass:    getEnv("DB_PASS", ""),
		DBName:    getEnv("DB_NAME", "go_ticket"),
		JWTSecret: getEnv("JWT_SECRET", ""),
	}

	validate(cfg)

	return cfg
}
