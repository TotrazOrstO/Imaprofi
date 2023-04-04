package config

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Postgres   PostgresConfig
	Redis      RedisConfig
	RabbitMQ   RabbitMQConfig
	Centrifugo CentrifugoConfig
	GRPCServer GRPCServerConfig
}

type PostgresConfig struct {
	User     string `envconfig:"POSTGRES_USER"`
	Password string `envconfig:"POSTGRES_PASSWORD"`
	DB       string `envconfig:"POSTGRES_DB"`
	Host     string `envconfig:"POSTGRES_HOST"`
	Port     int    `envconfig:"POSTGRES_PORT"`
}

type RedisConfig struct {
	Host     string `envconfig:"REDIS_HOST"`
	Port     int    `envconfig:"REDIS_PORT"`
	Password string `envconfig:"REDIS_PASSWORD"`
}

type RabbitMQConfig struct {
	URL string `envconfig:"RABBITMQ_URL"`
}

type CentrifugoConfig struct {
	URL string `envconfig:"CENTRIFUGO_URL"`
}

type GRPCServerConfig struct {
	Address string `envconfig:"GRPC_SERVER"`
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	var cfg Config
	err = envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("error processing environment variables: %v", err)
	}

	err = validateConfig(cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func validateConfig(cfg Config) error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.Postgres, validation.Required),
		validation.Field(&cfg.Redis, validation.Required),
		validation.Field(&cfg.RabbitMQ, validation.Required),
		validation.Field(&cfg.Centrifugo, validation.Required),
		validation.Field(&cfg.GRPCServer, validation.Required),
	)
}
