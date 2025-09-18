package main

import (
	"log"
	"time"

	"github.com/o5h/quiz/pkg/db"
)

var (
	version string = "dev"
	date    string = time.Now().Format("2006-01-02T15:04:05Z07:00")
)

func main() {
	log.Println("Quiz Server", "version:", version, "build date:", date)
	database, err := db.Open(&db.Config{
		URL:          "postgres://quiz_db_user:mysecretpassword@localhost:5432/quiz_db?sslmode=disable",
		MaxOpenConns: 10,
		MaxIdleConns: 5,
	})
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer database.Close()
}
