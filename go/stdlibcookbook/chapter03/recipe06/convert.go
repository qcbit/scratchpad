// Converting between binary, octal, decimal, and hexadecimal
package main

import (
	"fmt"
	"strconv"
)

const (
	bin      = "10111"
	hex      = "1A"
	oct      = "12"
	dec      = "10"
	floatNum = 16.123557
)

func main() {
	v, _ := ConvertInt(bin, 2, 16)
	fmt.Printf("Binary: %s\thex: %s\n", bin, v)

	v, _ = ConvertInt(hex, 16, 10)
	fmt.Printf("Hex: %s\tdec: %s\n", hex, v)

	v, _ = ConvertInt(oct, 8, 16)
	fmt.Printf("oct: %s\thex: %s\n", oct, v)

	v, _ = ConvertInt(dec, 10, 8)
	fmt.Printf("Dec: %s\toct: %s\n", dec, v)
}

func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}
