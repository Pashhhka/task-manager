package config

type Config struct {
	DBUrl string
}

func LoadConfig() *Config {
	return &Config{
		DBUrl: "postgres://postgres:postgres@localhost:5432/task_manager?sslmode=disable",
	}
}
