package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

const (
	BATCH_SIZE  = 20
	NUM_WORKER  = 3
	LOG_DIR     = "/log"
	SRC_DB_NAME = "Source Postgres DB"
	DST_DB_NAME = "Destination Postgres DB"
)

func NewPool(ctx context.Context, dbString string, name string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dbString)
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
		log.Printf("Database %s >>> New database connection established, adding to pool", name)
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

func main() {
	log.Print("Data Migration")

	if err := godotenv.Load(); err != nil {
		log.Panicf("Error loading env variables %v", err)
	}

	config := LoadConfig()

	srcPool, err := NewPool(context.Background(), config.SrcDBConfig.ToURLString(), SRC_DB_NAME)
	if err != nil {
		log.Panicf("Can not init srcDBPool: %v", err)
	}

	dstPool, err := NewPool(context.Background(), config.DstDBConfig.ToURLString(), DST_DB_NAME)
	if err != nil {
		log.Panicf("Can not init dstDBPool: %v", err)
	}

	s, err := NewService(srcPool, dstPool, BATCH_SIZE, NUM_WORKER, LOG_DIR)
	if err != nil {
		log.Panicf("Error initializing the migration service %v", err)
	}

	if err := s.StartMigration(); err != nil {
		log.Panicf("Service migration start error: %v", err)
	}
}
