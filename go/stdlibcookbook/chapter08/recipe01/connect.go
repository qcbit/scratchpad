// Connecting the database
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const db_host = "localhost:5432"

func main() {
	connStr := "postgres://postgres:postgres@" + db_host + "/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Ping OK")
}
