package power

import (
	"testing"
)

const ErrorString = `got %v expect %v given %v and %v`

func TestPower(t *testing.T) {
	testCases := []struct {
		a    int64
		b    int64
		want int
	}{
		{10000, 10000000, 144659229},
		{2, 32, 294967268},
		{31, 100, 908629681},
		{123, 123, 921450052},
	}

	t.Run("Generate Binary Strings Of Length N", func(t *testing.T) {
		for _, tt := range testCases {
			assertCorrect(t, calculate_power(tt.a, tt.b), tt.want, tt.a, tt.b)
		}
	})
}

func assertCorrect(t *testing.T, got, expect int, a, b int64) {
	t.Helper()
	t.Errorf(ErrorString, got, expect, a, b)
}
