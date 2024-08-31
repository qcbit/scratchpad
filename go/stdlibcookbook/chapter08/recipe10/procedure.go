// Executing stored procedures and functions
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const dbhost = "localhost"
const dbport = ":5432"
const dbname = "postgres"
const dbuser = "postgres"
const dbpwd = "postgres"

const call = "SELECT * FROM format_name($1, $2, $3)"
const callMySQL = "CALL simpleproc(?)"

type Result struct {
	Name     string
	Category int
}

func main() {
	db := createConnection()
	defer db.Close()
	r := Result{}

	if err := db.QueryRow(call, "John", "Doe", 32).Scan(&r.Name); err != nil {
		panic(err)
	}
	fmt.Printf("Result is: %+v\n", r)
}

func createConnection() *sql.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s%s/%s?sslmode=disable", dbuser, dbpwd, dbhost, dbport, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}
