package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
	APIBaseURL     string
	HTTPTimeout    time.Duration
	DefaultCommand string
}

func LoadConfig() *Config {
	apiURL := getEnv("API_BASE_URL", "https://pokeapi.co/api/v2")
	timeout := getEnvDuration("HTTP_TIMEOUT", 10*time.Second)
	defaultCmd := getEnv("DEFAULT_COMMAND", "help")

	return &Config{
		APIBaseURL:     apiURL,
		HTTPTimeout:    timeout,
		DefaultCommand: defaultCmd,
	}
}

func getEnv(key string, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	d, err := time.ParseDuration(val)
	if err != nil {
		log.Printf("Invalid duration for %s: %v. Using default.\n", key, err)
		return fallback
	}
	return d
}
