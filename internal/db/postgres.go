package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() error {

	connString :=
		"postgres://admin:admin@postgres:5432/orchestrator?sslmode=disable"

	fmt.Println("Connecting to:", connString)

	conn, err := pgxpool.New(
		context.Background(),
		connString,
	)

	if err != nil {
		return err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return err
	}

	DB = conn

	fmt.Println("Database connected successfully")

	return nil
}
