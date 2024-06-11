package learningStructs

import "testing"

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		checkArea(t, rect, 100.0)
	})

	t.Run("circle", func(t *testing.T) {
		circ := Circle{Point{0.0, 0.0}, 10.0}
		want := 314.1592653589793
		checkArea(t, circ, want)
	})
}