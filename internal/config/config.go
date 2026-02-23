package config

import "os"

type Config struct {
	DBUrl string
}

func LoadConfig() *Config {
	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/task_manager?sslmode=disable"
	}

	return &Config{
		DBUrl: dbURL,
	}
}
