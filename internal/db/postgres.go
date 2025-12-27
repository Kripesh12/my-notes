package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kripesh12/my-notes/internal/env"
)

var DB *pgxpool.Pool

func Connect() error {
	host := env.GetString("DB_HOST", "localhost")
	port := env.GetString("DB_PORT", "8080")
	dbUser := env.GetString("DB_USER", "postgres")
	dbPassword := env.GetString("DB_PASSWORD", "dsdsdwr")
	dbName := env.GetString("DB_NAME", "todo")

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPassword, host, port, dbName)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return err
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal("failed to connect to database: ", err.Error())
	}

	DB = pool
	log.Println("connected to database")

	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
