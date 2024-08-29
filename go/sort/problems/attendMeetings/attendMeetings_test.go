package attendmeetings

import (
	"testing"
)

const ErrorString = `got %v expect %v given %v`

func TestAttendMeetings(t *testing.T) {
	testCases := []struct {
		intervals [][]int
		expected int
	}{
		// {[][]int{{1,5}, {4,8}}, 0},
		// {[][]int{{1,5}, {5,8}, {10,15}}, 1},
		// {[][]int{{13, 56},{1, 3},{4, 5},{9990, 10341},{8, 10},{67, 9990}}, 1},
		{[][]int{{23, 50},{56, 89},{98, 100000000},{10000, 12341243},{34, 44},{1, 10},{22, 23}}, 0},
		//			 {1, 10},{22, 23},{23, 50},{34, 44},{56, 89},{67, 9990},{98, 100000000},{10000, 12341243}
	}
	t.Run("can attend all meetings", func(t *testing.T) {
		for _, tt := range testCases {
			assertCorrect(t, tt.intervals, can_attend_all_meetings(tt.intervals), tt.expected)
		}
	})
}

func assertCorrect(t *testing.T, test [][]int, got, expect int) {
	t.Helper()
	if got != expect {
		t.Errorf(ErrorString, got, expect, test)
	}
}
