package go_database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"testing"
)

func TestGoDataBaseConnection(t *testing.T) {
	connStr := "user=postgres dbname=belajar_golang_database sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	username := "albarms"
	rows, err := db.Query("SELECT * FROM users WHERE username = $1", username)

	fmt.Println(rows, "rows")
	fmt.Println(err, "errors")
}
