package reversepolishnotation

import "testing"

const ErrorString = `got %v want %v`

func TestAdder(t *testing.T) {
	t.Helper()
	calc := NewCalculator(20)
	tests := []struct {
		rpnString string
		want      float64
	}{
		{"4 3 - 5 *", 5},
		{"3 5 * 4 -", 11},
		{"15 3 /", 5},
		{"-21 3 /", -7},
		{"2 5 - 5 *", -15},
		{"3 -10 +", -7},
	}

	t.Run("Reverse Polish Notation Tests", func(t *testing.T) {
		for _, tt := range tests {
			ans := calc.calculate(tt.rpnString)
			if ans != tt.want {
				t.Errorf(ErrorString, ans, tt.want)
			}
		}
	})
}
