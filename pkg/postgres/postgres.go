package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Configs struct {
	Host     string `yaml:"POSTGRES_HOST" env:"POSTGRES_HOST" env-default:"localhost"`
	Port     string `yaml:"POSTGRES_PORT" env:"POSTGRES_PORT" env-default:"5432"`
	Username string `yaml:"POSTGRES_USER" env:"POSTGRES_USER" env-default:"root"`
	Password string `yaml:"POSTGRES_PASSWORD" env:"POSTGRES_PASSWORD" env-default:"1"`
	Database string `yaml:"POSTGRES_DB" env:"POSTGRES_DB" env-default:"postgres"`
}

func NewDB(c Configs) (*pgx.Conn, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.Username, c.Password, c.Host, c.Port, c.Database)
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, fmt.Errorf("error while connecting db %w", err)
	}

	return conn, nil
}
