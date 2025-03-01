package config

import (
	"delivery/pkg/postgres"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Postgres postgres.Configs `yaml:"POSTGRES" env:"POSTGRES"`

	GRPCPort int `yaml:"GRPC_PORT" env:"GRPC_PORT" env-default:"50051"`
}

func New() (*Config, error) {

	var cfg Config
	if err := cleanenv.ReadConfig("./config/config.yaml", &cfg); err != nil {
		return nil, fmt.Errorf("error while reading config file: %w", err)
	}

	return &cfg, nil
}
