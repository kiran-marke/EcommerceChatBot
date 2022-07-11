package domain

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

var databaseconn *pgx.Conn

func GetConnection() *pgx.Conn {
	client, err := pgx.Connect(context.Background(), "postgres://postgres:password@localhost:5432/postgres")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	//defer conn.Close(context.Background())

	return client
}
