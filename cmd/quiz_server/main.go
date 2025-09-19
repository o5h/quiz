package main

import (
	"log"
	"time"

	"github.com/o5h/config"
	"github.com/o5h/quiz/pkg/context"
	"github.com/o5h/quiz/pkg/controller"
	"github.com/o5h/quiz/pkg/db"
)

var (
	version string = "dev"
	date    string = time.Now().Format("2006-01-02T15:04:05Z07:00")
)

func main() {
	context.Init(version, date)
	must(config.Load(".env/config.yaml"))
	database := InitDatabase()
	defer database.Close()
	controller.Start()
}

func InitDatabase() db.Database {
	database, err := db.Open(&db.Config{
		URL:          config.Get("database.url", "postgres://localhost/quiz?sslmode=disable"),
		MaxOpenConns: config.Get("database.max_open_conns", 10),
		MaxIdleConns: config.Get("database.max_idle_conns", 5),
	})
	must(err)
	return database
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
