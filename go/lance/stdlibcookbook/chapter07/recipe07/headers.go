// Reading and writing HTTP headers
package main

import (
	"fmt"
	"net/http"
)

func main() {
	header := http.Header{}
	fmt.Println("Setting Auth-X in header")
	header.Set("Auth-X", "abcdef1234")
	fmt.Println(header)
	fmt.Println("header.Add(\"Auth-X\", \"defghijk\"")
	header.Add("Auth-X", "defghijk")
	fmt.Println(header)
	resSlice := header["Auth-X"]
	fmt.Printf("header[\"Auth-X\"] = %s\n", resSlice)
	resFirst := header.Get("Auth-X")
	fmt.Printf("header.Get(\"Auth-X\") = %s\n", resFirst)
	fmt.Println("Setting Auth-X to \"newvalue\"")
	header.Set("Auth-X", "newvalue")
	fmt.Println(header)
	fmt.Println("Deleting Auth-X")
	header.Del("Auth-X")
	fmt.Println(header)
}
