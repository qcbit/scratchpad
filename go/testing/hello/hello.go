package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "

func hello(name string) string {
	if name == "" {
		return englishHelloPrefix + "World"
	} else {
		return englishHelloPrefix + name
	}
}

func main() {
	fmt.Println(hello("Enter a name here"))
}
