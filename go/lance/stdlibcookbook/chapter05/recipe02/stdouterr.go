// Writing standard output and error
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	io.WriteString(os.Stdout, "This is a string to standard output.\n")
	io.WriteString(os.Stderr, "This is string to standard error output.\n")
	buf := []byte{0xAF, 0xFF, 0xFE}
	for i := 0; i < 200; i++ {
		if _, e := os.Stdout.Write(buf); e != nil {
			panic(e)
		}
	}
	fmt.Fprintln(os.Stdout)
}
