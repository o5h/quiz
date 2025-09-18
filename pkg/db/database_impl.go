package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/pressly/goose/v3"
)

type database struct {
	db      *sql.DB
	stmtMap map[string]*sql.Stmt
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

	err = db.initStatements()
	return err
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

func (db *database) initStatements() error {
	db.stmtMap = make(map[string]*sql.Stmt)
	//TODO: prepare statements here and add to stmtMap
	return nil
}

func (db *database) Close() error {
	if db.db != nil {
		return db.db.Close()
	}
	return nil
}

func (db *database) Begin(ctx context.Context) (*sql.Tx, func()) {
	tx, err := db.db.BeginTx(ctx, nil)
	must(err)
	return tx, func() { commit(tx) }
}

func (db *database) BeginOpt(ctx context.Context, op *sql.TxOptions) (*sql.Tx, func()) {
	tx, err := db.db.BeginTx(ctx, op)
	must(err)
	return tx, func() { commit(tx) }
}

func (db *database) GetStatement(name string) *sql.Stmt {
	if db.stmtMap == nil {
		db.stmtMap = make(map[string]*sql.Stmt)
	}
	return db.stmtMap[name]
}

func commit(tx *sql.Tx) {
	if r := recover(); r != nil {
		log.Println("Rollback recovery Tx", r)
		tx.Rollback()
		panic(r)
	} else {
		if err := tx.Commit(); err != nil {
			log.Println("Rollback commit Tx", err)
			tx.Rollback()
		}
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
