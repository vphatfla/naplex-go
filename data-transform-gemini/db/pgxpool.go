package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context) (*pgxpool.Pool, error) {

	config, err := pgxpool.ParseConfig("postgres://naplex_user:password@localhost:5432/app_db")
	if err != nil {
		return nil, err
	}

	config.MaxConns = 3
	config.MinConns = 2
	config.MaxConnLifetime = 1 * time.Hour
	config.MaxConnIdleTime = 30 * time.Minute
	config.HealthCheckPeriod = 1 * time.Minute
	config.ConnConfig.ConnectTimeout = 5 * time.Second

	config.AfterConnect = func(ctx context.Context, c *pgx.Conn) error {
		log.Println("Database >>> New database connection established, adding to pool")
		return nil
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}
	// defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Can not ping database by pool")
	}

	return pool, nil
}
