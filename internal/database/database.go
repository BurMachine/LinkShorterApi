package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var schema = `CREATE TABLE IF NOT EXISTS links (
	original_link VARCHAR(1000) NOT NULL,
	short_link VARCHAR(10) UNIQUE NOT NULL
);`

func InitConnection(url string) (*pgx.Conn, error) {
	dsn := os.Getenv("POSTGRES_URI")
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		conn, err = pgx.Connect(context.Background(), url)
		if err != nil {
			return conn, err
		}
	}
	_, err = conn.Exec(context.Background(), schema)
	if err != nil {
		return nil, fmt.Errorf("table creating error: %v", err)
	}
	return conn, nil
}

func SelectWhereLinkIs(link string, conn *pgx.Conn) (string, string) {
	origin := ""
	short := ""
	conn.QueryRow(context.Background(), `SELECT original_link,short_link FROM links WHERE original_link=$1`, link).Scan(&origin, &short)
	return origin, short
}

func SelectWhereShortLinkIs(link string, conn *pgx.Conn) string {
	origin := ""
	conn.QueryRow(context.Background(), `SELECT original_link FROM links WHERE short_link=$1`, link).Scan(&origin)
	return origin
}

func Insert(link string, short string, conn *pgx.Conn) error {
	_, err := conn.Exec(context.Background(), `INSERT INTO links (original_link, short_link) VALUES ($1, $2)`, link, short)
	if err != nil {
		return err
	}
	return nil
}
