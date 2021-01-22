// Formatting numbers
package main

import "fmt"

var integer int64 = 32500
var floatNum float64 = 22000.456

func main() {
	fmt.Printf("Common way to print decimal number: %d\n", integer)
	fmt.Printf("Show the sign: %+d\n", integer)
	fmt.Printf("Hexidecimal: %X\n", integer)
	fmt.Printf("Hexidecimal: %#X\n", integer)
	fmt.Printf("Leading zeros: %010d\n", integer)
	fmt.Printf("Left padding with spaces: % 10d\n", integer)
	fmt.Printf("Right padding: % -10d Padding right of number.\n", integer)
	fmt.Printf("Floating-point: %f\n", floatNum)
	fmt.Printf("Floating-point with precision 5: %.5f\n", floatNum)
	fmt.Printf("Floating-point with scientific notation: %e\n", floatNum)
	fmt.Printf("Floating-point for large exponents: %g\n", floatNum)
}
