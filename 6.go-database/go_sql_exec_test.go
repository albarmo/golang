package go_database

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}()

	sqlScript := `INSERT INTO users(id,username,password,email) VALUES (2,"elang","elang","elangmaina@mantab.com")`

	ctx := context.Background()
	_, err := db.ExecContext(ctx, sqlScript)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")

}
