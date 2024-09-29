package uniquepaths

import (
	"testing"
)

const ErrorString = `got %v expect %v`

func TestUniquePaths(t *testing.T) {
	testCases := []struct {
		given map[string]int
		want  int
	}{
		{map[string]int{"n": 1000, "m": 1000}, 965601742},
		{map[string]int{"n": 999, "m": 1000}, 482800871},
	}

	t.Run("Unique Paths", func(t *testing.T) {
		for _, tt := range testCases {
			got := unique_paths(tt.given["n"], tt.given["m"])
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
