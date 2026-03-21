package nchoosekcombinations

import (
	"reflect"
	"testing"
)

const ErrorString = `got %v expect %v given %v`

func TestNChooseKCombinations(t *testing.T) {
	testCases := []struct {
		given    map[string]int
		expected [][]int
	}{
		{
			map[string]int{
				"n": 5,
				"k": 2,
			}, [][]int{{1, 2},
				{1, 3},
				{1, 4},
				{1, 5},
				{2, 3},
				{2, 4},
				{2, 5},
				{3, 4},
				{3, 5},
				{4, 5}}},
		{map[string]int{
			"n": 6,
			"k": 6,
		}, [][]int{{1, 2, 3, 4, 5, 6}}},
		{map[string]int{
			"n": 6,
			"k": 2,
		}, [][]int{{1, 2},
			{1, 3},
			{1, 4},
			{1, 5},
			{1, 6},
			{2, 3},
			{2, 4},
			{2, 5},
			{2, 6},
			{3, 4},
			{3, 5},
			{3, 6},
			{4, 5},
			{4, 6},
			{5, 6}}},
		{map[string]int{
			"n": 6,
			"k": 3,
		}, [][]int{{1, 2, 3},
			{1, 2, 4},
			{1, 2, 5},
			{1, 2, 6},
			{1, 3, 4},
			{1, 3, 5},
			{1, 3, 6},
			{1, 4, 5},
			{1, 4, 6},
			{1, 5, 6},
			{2, 3, 4},
			{2, 3, 5},
			{2, 3, 6},
			{2, 4, 5},
			{2, 4, 6},
			{2, 5, 6},
			{3, 4, 5},
			{3, 4, 6},
			{3, 5, 6},
			{4, 5, 6}}},
		{map[string]int{
			"n": 1,
			"k": 1,
		}, [][]int{{1}}},
		{map[string]int{
			"n": 6,
			"k": 5,
		}, [][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 6}, {1, 2, 3, 5, 6}, {1, 2, 4, 5, 6}, {1, 3, 4, 5, 6}, {2, 3, 4, 5, 6}}},
		{map[string]int{
			"n": 6,
			"k": 4,
		}, [][]int{{1, 2, 3, 4},
			{1, 2, 3, 5},
			{1, 2, 3, 6},
			{1, 2, 4, 5},
			{1, 2, 4, 6},
			{1, 2, 5, 6},
			{1, 3, 4, 5},
			{1, 3, 4, 6},
			{1, 3, 5, 6},
			{1, 4, 5, 6},
			{2, 3, 4, 5},
			{2, 3, 4, 6},
			{2, 3, 5, 6},
			{2, 4, 5, 6},
			{3, 4, 5, 6}}},
	}

	t.Run("nChooseKCombinations", func(t *testing.T) {
		for _, tt := range testCases {
			assertCorrect(t, find_combinations(tt.given["n"], tt.given["k"]), tt.expected, tt.given)
		}
	})
}

func assertCorrect(t *testing.T, got, expect [][]int, test map[string]int) {
	t.Helper()
	if !reflect.DeepEqual(got, expect) {
		t.Errorf(ErrorString, got, expect, test)
	}
}
