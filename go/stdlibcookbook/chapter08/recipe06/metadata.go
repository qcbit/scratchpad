// Reading query result metadata
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

const sel = "SELECT * FROM post p"

func main() {
	db := createConnection()
	defer db.Close()

	rs, err := db.Query(sel)
	if err != nil {
		panic(err)
	}
	defer rs.Close()
	columns, err := rs.Columns()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected columns: %v\n", columns)

	colTypes, err := rs.ColumnTypes()
	if err != nil {
		panic(err)
	}
	for _, col := range colTypes {
		fmt.Println()
		fmt.Printf("%+v\n", col)
	}
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
