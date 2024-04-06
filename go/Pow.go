// https://go.dev/play/p/2SbfA1xRiNE
package main

import (
	"fmt"
)

func main() {
	fmt.Println(Pow(5, 3))
}

func Pow(a, b int) int {
	var answer int
	switch b {
	case 0:
		answer = 1
	case 1:
		answer = a
	default:
		x := a
		answer = a
		i := 2
		for ; i <= b; i++ {
			answer = answer * x
		}
	}
	return answer
}
