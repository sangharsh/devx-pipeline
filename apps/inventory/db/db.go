package db

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pool *pgxpool.Pool
	once sync.Once
)

func InitDB(ctx context.Context, connectionString string) error {
	var err error
	once.Do(func() {
		pool, err = pgxpool.New(ctx, connectionString)
		// You can set pool configuration here if needed
		// config.MaxConns = 25
	})
	return err
}

func GetPool() *pgxpool.Pool {
	return pool
}

func CloseDB() {
	if pool != nil {
		pool.Close()
	}
}
