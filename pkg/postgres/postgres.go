package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Configs struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func newDatabase(c Configs) (*pgx.Conn, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.Username, c.Password, c.Host, c.Port, c.Database)
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, fmt.Errorf("Error while connecting db %w\n", err)
	}

	return conn, nil
}
