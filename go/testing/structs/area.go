package learningStructs

import "math"

func (rect Rectangle) Area() float64 {
	return rect.Length * rect.Width
}

func (circ Circle) Area() float64 {
	return circ.Radius * circ.Radius * math.Pi
}

func (tri Triangle) Area() float64 {
	return .5 * tri.Base * tri.Height
}
