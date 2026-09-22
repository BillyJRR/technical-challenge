package config

import (
	"os"
)

type Config struct {
	Port           string
	NodeServiceURL string
	JWTSecret      string
}

func LoadConfig() *Config {
	return &Config{
		Port:           getEnv("PORT", "8080"),
		NodeServiceURL: getEnv("NODE_SERVICE_URL", "http://localhost:3000/api/v1/statistics"),
		JWTSecret:      getEnv("JWT_SECRET", "super-secret-key"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
