// Operating complex numbers
package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	a := complex(2, 3)

	fmt.Printf("Real part: %f\n", real(a))
	fmt.Printf("Complex part: %f\n", imag(a))

	b := complex(6, 4)
	fmt.Printf("a - b = %v\n", a-b)
	fmt.Printf("a + b = %v\n", a+b)
	fmt.Printf("a * b = %v\n", a*b)
	fmt.Printf("a / b = %v\n", a/b)

	conjugate := cmplx.Conj(a)
	fmt.Println("Complex number a's conjugate : ", conjugate)

	cos := cmplx.Cos(b)
	fmt.Println("Cosine of b : ", cos)
}
