package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
)

type Database struct {
	Client *sqlx.DB
}

func NewDatabase() (*Database, error) {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)

	dbConn, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return &Database{}, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Database{Client: dbConn}, nil
}

func (db *Database) Ping(ctx context.Context) error {
	return db.Client.PingContext(ctx)
}
