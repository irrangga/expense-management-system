package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		HTTP     HTTPConfig
		Postgres PostgresConfig
	}

	HTTPConfig struct {
		Port string `env:"HTTP_PORT"`
	}

	PostgresConfig struct {
		Host     string `env:"POSTGRES_HOST"`
		Username string `env:"POSTGRES_USER"`
		Password string `env:"POSTGRES_PASSWORD"`
		Name     string `env:"POSTGRES_DB"`
		Port     string `env:"POSTGRES_PORT"`
		SslMode  string `env:"POSTGRES_SSL_MODE"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	godotenv.Load()

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
