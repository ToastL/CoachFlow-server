package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect(url string) {
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Fatalf("Unable to parse database URL: %v\n", err)
	}

	cfg.MaxConns = 10
	cfg.MinConns = 2
	
	DB, err = pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		log.Fatalf("Unable to create DB pool: %v\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := DB.Ping(ctx); err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	log.Println("Connected to the database successfully")
}