package main

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vphatfla/naplex-go/data-migration/database/dstDB"
	"github.com/vphatfla/naplex-go/data-migration/database/srcDB"
)

// Migration service
type service struct {
	srcRepository *srcDB.Queries
	dstRepository *dstDB.Queries
	batchSize     int
}

func NewService(srcPool *pgxpool.Pool, dstPool *pgxpool.Pool, bacthSize int) *service {
	return &service{
		srcRepository: srcDB.New(srcPool),
		dstRepository: dstDB.New(dstPool),
		batchSize:     bacthSize,
	}
}
