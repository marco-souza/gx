package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Hostname string
	Port     string
}

func Load() *Config {
	// load .env file
	godotenv.Load()

	return &Config{
		Hostname: env("HOST", "localhost"),
		Port:     env("PORT", "3001"),
	}
}

func env(varEnv string, defaultValue string) string {
	value := os.Getenv(varEnv)
	if value != "" {
		return value
	}
	return defaultValue
}
