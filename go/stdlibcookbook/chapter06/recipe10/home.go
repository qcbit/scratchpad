// Resolving the user home directory
package main

import (
	"fmt"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println("The user home directory: " + usr.HomeDir)
}
