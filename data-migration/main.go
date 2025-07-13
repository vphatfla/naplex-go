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
	BATCH_SIZE = 20
	NUM_WORKER = 3
)

func NewPool(ctx context.Context, dbString string) (*pgxpool.Pool, error) {
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

func main() {
	log.Print("Data Migration")

	if err := godotenv.Load(); err != nil {
		log.Panicf("Error loading env variables %v", err)
	}

	config := LoadConfig()

	srcPool, err := NewPool(context.Background(), config.SrcDBConfig.ToURLString())
	if err != nil {
		log.Panicf("Can not init srcDBPool: %v", err)
	}

	dstPool, err := NewPool(context.Background(), config.DstDBConfig.ToURLString())
	if err != nil {
		log.Panicf("Can not init dstDBPool: %v", err)
	}

	s := NewService(srcPool, dstPool, BATCH_SIZE, NUM_WORKER)

	if err := s.StartMigration(); err != nil {
		log.Panicf("Service migration start error: %v", err)
	}
}
