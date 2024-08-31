// Demonstration of set/unset env variables
package main

import (
	"log"
	"os"
)

func main() {
	connStr := os.Getenv("DB_CONN")
	log.Printf("Connection string: %s\n", connStr)
}