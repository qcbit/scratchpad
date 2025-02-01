package twoSumInASortedArray

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
		{[]int{1, 2, 3, 5, 10}, 7, [][]int{{1, 3}, {3, 1}}},
		{[]int{1, 2, 3, 5, 10}, 9, [][]int{{-1, -1}}},
		{[]int{0, 0, 0, 0, 0}, 0, [][]int{{0, 1}, {1, 0}, {0, 2}, {2, 0}, {0, 3}, {3, 0}, {0, 4}, {4, 0}}},
		{[]int{0, 0, 0, 0, 0}, 1, [][]int{{-1, -1}}},
		{[]int{-20, -11, 0, 1, 12}, 1, [][]int{{1, 4}, {4, 1}, {3, 2}, {2, 3}}},
	}
	t.Run("BruteForce", func(t *testing.T) {
		for _, tt := range testCases {
			assertCorrect(t, BruteForce(tt.numbers, tt.target), tt.expected)
		}
	})

	t.Run("BruteForceTwoPointers", func(t *testing.T) {
		for _, tt := range testCases {
			assertCorrect(t, BruteForceTwoPointers(tt.numbers, tt.target), tt.expected)
		}
	})

	t.Run("BinarySearch", func(t *testing.T) {
		for _, tt := range testCases {
			assertCorrect(t, BinarySearch(tt.numbers, tt.target), tt.expected)
		}
	})

	t.Run("Transform", func(t *testing.T) {
		for _, tt := range testCases {
			assertCorrect(t, Transform(tt.numbers, tt.target), tt.expected)
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
