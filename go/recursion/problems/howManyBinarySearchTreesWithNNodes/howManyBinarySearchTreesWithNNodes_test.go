package howmanybinarysearchtreeswithnnodes

import (
	"testing"
)

const ErrorString = `got %v expect %v given %v`

func TestNChooseKCombinations(t *testing.T) {
	testCases := []struct {
		n        int
		expected int64
	}{
		{n: 1, expected: 1},
		{n: 2, expected: 2},
		{n: 3, expected: 5},
		{n: 4, expected: 14},
		{n: 5, expected: 42},
		{n: 6, expected: 132},
		{n: 7, expected: 429},
		{n: 8, expected: 1430},
		{n: 9, expected: 4862},
		{n: 10, expected: 16796},
		{n: 11, expected: 58786},
		{n: 12, expected: 208012},
		{n: 13, expected: 742900},
		{n: 14, expected: 2674440},
		{n: 15, expected: 9694845},
		{n: 16, expected: 35357670},
	}

	t.Run("nChooseKCombinations", func(t *testing.T) {
		for _, tt := range testCases {
			actual := how_many_bsts(tt.n)
			assertCorrect(t, actual, tt.expected, tt.n)
		}
	})
}

func assertCorrect(t *testing.T, got, expected int64, given int) {
	t.Helper()
	if got != expected {
		t.Errorf(ErrorString, got, expected, given)
	}
}
