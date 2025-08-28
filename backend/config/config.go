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
		Payment  PaymentConfig
	}

	HTTPConfig struct {
		Port string `env:"HTTP_PORT"`
	}

	PostgresConfig struct {
		Host     string `env:"DB_HOST"`
		Username string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Name     string `env:"DB_NAME"`
		Port     string `env:"DB_PORT"`
		SslMode  string `env:"DB_SSL_MODE"`
	}

	AuthConfig struct {
		SecretKey string `env:"JWT_SECRET_KEY"`
	}

	PaymentConfig struct {
		PaymentProcessorURL string `env:"PAYMENT_PROCESSOR_URL"`
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
