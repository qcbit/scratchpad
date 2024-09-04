// Rounding floating point numbers
package main

import (
	"fmt"
	"math"
)

var valA float64 = 3.55554444

func main() {
	intVal := int(valA)
	fmt.Printf("Bad rounding by casting to int: %v\n", intVal)

	fRound := Round(valA)
	fmt.Printf("Rounding by custom function: %v\n", fRound)
}

func Round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}
