package db

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Database interface {
	Begin(context.Context) (*sql.Tx, func())
	BeginOpt(context.Context, *sql.TxOptions) (*sql.Tx, func())
	GetStatement(name string) *sql.Stmt
	Close() error
}

func Open(cfg *Config) (Database, error) {
	dbInstance := &database{}
	if err := dbInstance.init(cfg); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
		return nil, err
	}
	return dbInstance, nil
}
