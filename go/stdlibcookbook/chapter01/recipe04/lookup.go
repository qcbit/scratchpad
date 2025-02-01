package main

import (
	"log"
	"os"
	"fmt"
)

func main() {
	key := "DB_CONN"

	connStr, ex := os.LookupEnv(key)
	if !ex {
		log.Printf("The env variable %s is not set.\n", key)
	}
	fmt.Println(connStr)
}