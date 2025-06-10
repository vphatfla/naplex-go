package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vphatfla/naplex-go/backend/internal/config"
)

func NewPool(ctx context.Context, config *config.Config) (*pgxpool.Pool, error) {
	dbConfig, err := pgxpool.ParseConfig(config.DBConfig.ToURLString())
	if err != nil {
		return nil, fmt.Errorf("Error create db config from conn string -> %v", err)
	}

	dbConfig.MaxConns = 3
	dbConfig.MinConns = 2
	dbConfig.MaxConnLifetime = 1 * time.Hour
	dbConfig.MaxConnIdleTime = 30 * time.Minute
	dbConfig.HealthCheckPeriod = 1 * time.Minute
	dbConfig.ConnConfig.ConnectTimeout = 5 * time.Second

	dbConfig.AfterConnect = func(ctx context.Context, c *pgx.Conn) error {
		log.Println("Database >>> New database connection established, adding to pool")
		return nil
	}

	pool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		return nil, fmt.Errorf("Error create db config from conn string -> %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Can not ping database by pool")
	}

	return pool, nil
}
