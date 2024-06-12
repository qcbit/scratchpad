package learningStructs

import "testing"

func TestArea(t *testing.T) {

	areaTests := []struct {
		Name string
		shape Shape
    	hasArea  float64
	} {
		{"Rectangle", Rectangle{Width: 12, Length: 6}, 72.0},
    	{"Circle",Circle{Point{X: 0, Y: 0}, 10}, 314.1592653589793},
    	{"Triangle", Triangle{Base: 12,Height: 6},36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.Name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
			}
		})
		
	}

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