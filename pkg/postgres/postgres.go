package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
)

type Config struct {
	Host     string `yaml:"POSTGRES_HOST" env:"POSTGRES_HOST" env-default:"localhost"`
	Port     string `yaml:"POSTGRES_PORT" env:"POSTGRES_PORT" env-default:"5432"`
	Username string `yaml:"POSTGRES_USER" env:"POSTGRES_USER" env-default:"root"`
	Password string `yaml:"POSTGRES_PASSWORD" env:"1" env-default:"1"`
	Database string `yaml:"POSTGRES_DB" env:"POSTGRES_DB" env-default:"postgres"`
}

func NewDB(c Config) (*pgx.Conn, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.Username, c.Password, c.Host, c.Port, c.Database)
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, fmt.Errorf("error while connecting db %w", err)
	}

	m, err := migrate.New(
		"file://db/migrations",
		dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	return conn, nil
}
