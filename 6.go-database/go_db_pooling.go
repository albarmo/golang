package go_database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func GetConnection() *sql.DB {
	connStr := "host=localhost port=5432 user=albarms dbname=belajar_golang_database sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	// setup connection pooling
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
