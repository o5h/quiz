package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

var dbInstance *database

type Config struct {
	URL          string
	MaxOpenConns int
	MaxIdleConns int
}

func Init(cfg *Config) error {
	dbInstance = &database{}
	return dbInstance.init(cfg)
}

func Close() {
	if dbInstance != nil && dbInstance.db != nil {
		dbInstance.close()
	}
}

type database struct {
	db *sql.DB
}

func (db *database) init(cfg *Config) error {
	var err error
	db.db, err = sql.Open("postgres", cfg.URL)
	if err != nil {
		return err
	}
	db.db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.db.SetMaxIdleConns(cfg.MaxIdleConns)

	err = db.initializeSchema()
	if err != nil {
		return err
	}

	return nil
}

func (db *database) initializeSchema() error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return err
	}
	goose.SetBaseFS(schemaSQLFiles)
	goose.SetLogger(log.Default())
	err = goose.Up(db.db, "sql/schema")
	if err != nil {
		return err
	}

	return nil
}

func (db *database) close() {
	if db.db != nil {
		db.db.Close()
	}
}
