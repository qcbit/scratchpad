// Handling transactions
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

const selOne = "SELECT id, title, content FROM post WHERE id=$1;"
const insert = "INSERT INTO post(ID, TITLE, CONTENT) VALUES(4,'Transaction Title', 'Transaction Content');"

type Post struct {
	ID      int
	Title   string
	Content string
}

func main() {
	db := createConnection()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	_, err = tx.Exec(insert)
	if err != nil {
		panic(err)
	}

	p := Post{}
	// Query in another session/transaction
	if err := db.QueryRow(selOne, 4).Scan(&p.ID, &p.Title, &p.Content); err != nil {
		fmt.Println("Got error for db.Query:" + err.Error())
	}
	fmt.Println(p)
	// Query within transaction
	if err := tx.QueryRow(selOne, 4).Scan(&p.ID, &p.Title, &p.Content); err != nil {
		fmt.Println("got error for db.Query:" + err.Error())
	}
	fmt.Println(p)
	tx.Rollback()
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
