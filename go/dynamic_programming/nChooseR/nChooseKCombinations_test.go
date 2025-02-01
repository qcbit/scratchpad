package nchooser

import (
	"testing"
)

const ErrorString = `got %v expect %v`

func TestNChooseR(t *testing.T) {
	testCases := []struct {
		given map[string]int
		want  int
	}{
		{map[string]int{"n": 3, "r": 5}, 0},
		{map[string]int{"n": 0, "r": 0}, 1},
		{map[string]int{"n": 5, "r": 10}, 0},
		{map[string]int{"n": 8, "r": 3}, 56},
		{map[string]int{"n": 6, "r": 2}, 15},
		{map[string]int{"n": 100, "r": 100}, 1},
		{map[string]int{"n": 1000, "r": 0}, 1},
		{map[string]int{"n": 1, "r": 1}, 1},
		{map[string]int{"n": 13, "r": 5}, 1287},
		{map[string]int{"n": 200, "r": 100}, 407336795},
		{map[string]int{"n": 1000, "r": 500}, 159835829},
		{map[string]int{"n": 851, "r": 425}, 946589377},
		{map[string]int{"n": 646, "r": 323}, 870663212},
		{map[string]int{"n": 743, "r": 371}, 407435672},
		{map[string]int{"n": 993, "r": 496}, 504907690},
		{map[string]int{"n": 686, "r": 343}, 267158333},
		{map[string]int{"n": 739, "r": 871}, 0},
		{map[string]int{"n": 730, "r": 745}, 0},
		{map[string]int{"n": 776, "r": 523}, 857313578},
		{map[string]int{"n": 977, "r": 881}, 206093141},
		{map[string]int{"n": 870, "r": 785}, 645654742},
		{map[string]int{"n": 978, "r": 640}, 215525281},
	}

	t.Run("nChooserCombinations", func(t *testing.T) {
		for _, tt := range testCases {
			got := ncr(tt.given["n"], tt.given["r"])
			assertCorrect(t, got, tt.want)
		}
	})
}

func assertCorrect(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf(ErrorString, got, want)
	}
}
