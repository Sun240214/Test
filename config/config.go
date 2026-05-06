package config

import "os"

// Config contains application-level settings loaded from the environment.
type Config struct {
	Port string
}

// Load reads configuration from environment variables and applies defaults.
func Load() Config {
	return Config{
		Port: getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
