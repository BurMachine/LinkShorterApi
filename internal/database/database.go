package database

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func InitConnection(url string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return conn, err
	}

	return conn, nil
}
