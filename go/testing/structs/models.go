package learningStructs

type Shape interface {
  Area() float64
}

type Rectangle struct {
  Length float64
  Width  float64
}

type Point struct {
  X float64
  Y float64
}

type Circle struct {
  Point  Point
  Radius float64
}

type Triangle struct {
  Base float64
  Height float64
}