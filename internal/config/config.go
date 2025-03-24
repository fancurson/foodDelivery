package config

import (
	"delivery/pkg/postgres"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Postgres postgres.Config `yaml:"POSTGRES" env:"POSTGRES"`
	RestPort int             `yaml:"REST_PORT" env:"REST_PORT" env-default:"8081"`

	GRPCPort int `yaml:"GRPC_PORT" env:"GRPC_PORT" env-default:"50051"`
}

func New() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, fmt.Errorf("error while reading config occured: %w", err)
	}
	return &cfg, nil
}
