package main

import (
	"fmt"
)

func printSlice[T any](slice []T) string {
  var retstr string
	for i, k := range slice {
		retstr += fmt.Sprintf("%d: %s", i, printGeneric[T](k))
	}
  return retstr
}

func printGeneric[T any](t T) string {
	return fmt.Sprintf("hello - %v", t)
}

func main() {
	fmt.Println("hello")
}
