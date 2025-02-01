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

const sel = `SELECT title, content FROM post;
SELECT 1234 NUM;`

const selOne = "SELECT title, content FROM post WHERE ID = $1;"

type Post struct {
	Name sql.NullString
	Text sql.NullString
}

func main() {
	db := createConnection()
	defer db.Close()

	rs, err := db.Query(sel)
	if err != nil {
		panic(err)
	}
	defer rs.Close()

	posts := []Post{}
	for rs.Next() {
		if rs.Err() != nil {
			panic(rs.Err())
		}
		p := Post{}
		if err := rs.Scan(&p.Name, &p.Text); err != nil {
			panic(err)
		}
		posts = append(posts, p)
	}

	var num int
	if rs.NextResultSet() {
		for rs.Next() {
			if rs.Err() != nil {
				panic(rs.Err())
			}
			rs.Scan(&num)
		}
	}

	fmt.Printf("Retrieved posts: %+v\n", posts)
	fmt.Printf("Retrieved number: %d\n", num)

	row := db.QueryRow(selOne, 100)
	or := Post{}
	if err := row.Scan(&or.Name, &or.Text); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	fmt.Printf("Retireved one post: %+v\n", or)
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
