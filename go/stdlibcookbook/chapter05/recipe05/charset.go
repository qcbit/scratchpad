// Read/write charset
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	encoder := charmap.Windows1252.NewEncoder()
	s, e := encoder.String("This is a sample text with runes Š")
	if e != nil {
		panic(e)
	}
	ioutil.WriteFile("example.txt", []byte(s), os.ModePerm)

	// decoder
	f, e := os.Open("example.txt")
	if e != nil {
		panic(e)
	}
	defer f.Close()
	decoder := charmap.Windows1252.NewDecoder()
	reader := decoder.Reader(f)
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
