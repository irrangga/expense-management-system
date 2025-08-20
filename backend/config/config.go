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
		Auth     AuthConfig
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

	AuthConfig struct {
		SecretKey string `env:"JWT_SECRET_KEY"`
	}
)

var Cfg = &Config{}

// NewConfig returns app config.
func NewConfig() error {
	godotenv.Load()

	if err := env.Parse(Cfg); err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	return nil
}
