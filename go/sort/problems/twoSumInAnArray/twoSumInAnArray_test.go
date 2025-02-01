package twoSumInAnArray

import (
	"reflect"
	"testing"
)

const ErrorString = `got %v expected %v for target %v given %v`

func TestTwoSumInASortedArray(t *testing.T) {
	testCases := []struct {
		numbers  []int
		target   int
		expected [][]int
	}{
		{[]int{3, 2, 1, 10, 5}, 7, [][]int{{1, 4}, {4, 1}}},
		{[]int{2, 5, 1, 3, 10}, 9, [][]int{{-1, -1}}},
		{[]int{0, 0, 0, 0, 0}, 0, [][]int{{0, 1}, {1, 0}, {0, 2}, {2, 0}, {0, 3}, {3, 0}, {0, 4}, {4, 0}}},
		{[]int{0, 0, 0, 0, 0}, 1, [][]int{{-1, -1}}},
		{[]int{-11, 1, -20, 0, 12}, 1, [][]int{{0, 4}, {4, 0}, {1, 3}, {3, 1}}},
	}
	t.Run("two_sum", func(t *testing.T) {
		for _, tt := range testCases {
			assertCorrect(t, two_sum(tt.numbers, tt.target), tt.expected)
		}
	})
}

func assertCorrect(t *testing.T, got []int, expected [][]int) {
	t.Helper()
	found := false
	for _, x := range expected {
		if reflect.DeepEqual(got, x) {
			found = true
		}
	}
	if !found {
		t.Errorf("got %v expected one of these %v", got, expected)
	}
}
