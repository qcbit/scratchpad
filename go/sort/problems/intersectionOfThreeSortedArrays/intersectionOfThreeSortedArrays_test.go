package intersectionofthreesortedarrays

import (
	"testing"
	"reflect"
)

const ErrorString = `got %v expect %v given %v`

func TestAttendMeetings(t *testing.T) {
	testCases := []struct {
		arrays [][]int
		expect [] int
	}{
		{[][]int{{2, 5, 10}, {2, 3, 4, 10}, {2, 4, 10}}, []int{2, 10}},
		{[][]int{{1, 2, 3}, {}, {2, 2}}, []int{-1}},
		{[][]int{{1, 2, 2, 2, 9}, {1, 1, 2, 2}, {1, 1, 1, 2, 2, 2}}, []int{1, 2, 2}},
		{[][]int{{3, 30, 300}, {3, 30}, {30, 300}}, []int{30}},
	}
	t.Run("intersection", func(t *testing.T) {
		for _, tt := range testCases {
			assertCorrect(t, tt.arrays, find_intersection(tt.arrays[0], tt.arrays[1], tt.arrays[2]), tt.expect)
		}
	})
}

func assertCorrect(t *testing.T, test [][]int, got, expect []int) {
	t.Helper()
	if !reflect.DeepEqual(got, expect) {
		t.Errorf(ErrorString, got, expect, test)
	}
}
