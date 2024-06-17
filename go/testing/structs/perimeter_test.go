package learningStructs

import "testing"

func TestPerimeter(t *testing.T) {

  assertCorrect := func(t testing.TB, got, want float64) {
    t.Helper()
    if got != want {
      t.Errorf("got %.2f want %.2f", got, want)
    }
  }

  t.Run("rectangle", func(t *testing.T) {
    rect := Rectangle{10.0, 10.0}
    got := rect.Perimeter()
    want := 40.0
    assertCorrect(t, got, want)
  })

  t.Run("circle", func(t *testing.T) {
    circ := Circle{Point{0, 0}, 10.0}
    got := circ.Perimeter()
    want := 62.83185307179586
    assertCorrect(t, got, want)
  })
}

