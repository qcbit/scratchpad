package permuteArrayOfUniqueIntegers

import (
	"reflect"
	"testing"
)

const ErrorString = `got %v expect %v given %v`

func TestPermuteArrayOfUniqueIntegers(t *testing.T) {
	testCases := []struct {
		array []int
		want  [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}},
		},
	}

	t.Run("Get Permutations", func(t *testing.T) {
		for _, tt := range testCases {
			assertCorrect(t, get_permutations(tt.array), tt.want, tt.array)
		}
	})
}

func assertCorrect(t *testing.T, got, expect [][]int, given []int) {
	t.Helper()
	if !reflect.DeepEqual(got, expect) {
		t.Errorf(ErrorString, got, expect, given)
	}
}
