// Validating the connection
package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

//const dbhost = "localhost"
const dbhost = "localhost"
const dbport = ":5432"

func main() {
	connStr := "postgres://postgres:postgres@" + dbhost + dbport + "/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Pink OK.")
	ctx, _ := context.WithTimeout(context.Background(), time.Nanosecond)
	if err = db.PingContext(ctx); err != nil {
		fmt.Println("Error: " + err.Error())
	}

	conn, err := db.Conn(context.Background())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	if err = conn.PingContext(context.Background()); err != nil {
		panic(err)
	}
	fmt.Println("Connection Pink OK.")
}
