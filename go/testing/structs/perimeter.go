package learningStructs

import "math"

func (rect Rectangle) Perimeter() float64 {
	return 2 * (rect.Length + rect.Width)
}

func (circ Circle) Perimeter() float64 {
	return 2 * math.Pi * circ.Radius
}
