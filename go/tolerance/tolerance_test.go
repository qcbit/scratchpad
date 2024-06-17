package main

import (
	"reflect"
	"testing"
)

func TestTolerance(T *testing.T) {
	tests := []struct {
		tolerance int
		values    []int
		expected  [][]int
	}{
		{
			tolerance: 5,
			values:    []int{1, 2, 4, 10, 11, 15, 22, 23, 40, 50, 100, 101, 120},
			expected:  [][]int{{1, 4}, {10, 15}, {22, 23}, {40}, {50}, {100, 101}, {120}},
		}, {
			tolerance: 5,
			values:    []int{1, 2, 4, 10, 11, 15, 22, 23, 27, 35, 40, 46, 50},
			expected:  [][]int{{1, 4}, {10, 15}, {22, 27}, {35, 40}, {46, 50}},
		},
	}

	assertCorrect := func(t *testing.T, tolerance int, values []int, expected [][]int) {
		T.Helper()
		if !reflect.DeepEqual(findToleranceSet(tolerance, values), expected) {
			t.Errorf("Expected %v, but got %v", expected, findToleranceSet(tolerance, values))
		}
	}

	for test := range tests {
		assertCorrect(T, tests[test].tolerance, tests[test].values, tests[test].expected)
	}
}
