package shape

import (
	"math"
	"testing"
)

func TestShape(t *testing.T) {
	checkResult := func(t *testing.T, actual, expected float64) {
		t.Helper()
		if actual != expected {
			t.Errorf("Actual: %.2f Expected: %.2f", actual, expected)
		}
	}

	t.Run("Test Perimeter of a Rectangle", func(t *testing.T) {
		rectangle := Rectangle{Width: 10.0, Height: 10.0}
		actual := rectangle.Perimeter()
		expected := float64(40.0)
		checkResult(t, actual, expected)
	})

	checkArea := func(t *testing.T, shape Shape, expectArea float64) {
		t.Helper()
		area := shape.Area()
		if area != expectArea {
			t.Errorf("Actual: %.2f Expected: %.2f", area, expectArea)
		}
	}

	t.Run("Test Area of a Rectangle", func(t *testing.T) {
		rectangle := Rectangle{Width: 10.0, Height: 10.0}
		checkArea(t, rectangle, float64(100.0))
	})

	t.Run("Test Area of a Circle", func(t *testing.T) {
		circle := Circle{X: 0, Y: 0, R: 5.0}
		checkArea(t, circle, float64(math.Pi*5.0*5.0))
	})

	areaTests := []struct {
		shape  Shape
		expect float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, expect: 72.0},
		{shape: Circle{X: 0, Y: 0, R: 10}, expect: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, expect: 36},
	}

	for _, tt := range areaTests {
		actual := tt.shape.Area()
		expected := tt.expect
		if actual != expected {
			t.Errorf("%#v, Actual: %.2f Expected: %.2f", tt.shape, actual, expected)
		}
	}
}
