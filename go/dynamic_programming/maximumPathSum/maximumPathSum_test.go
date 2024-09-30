package maximumpathsum

import "testing"

const ErrorString = `got %v want %v`

func TestMaximumPathSum(t *testing.T) {
	t.Helper()
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{4, 5, 8}, {3, 6, 4}, {2, 4, 7}}, 28},
		{[][]int{{1, -4, 3}, {4, -2, 2}}, 5},
		{[][]int{{-10000}}, -10000},
		{[][]int{{10000}}, 10000},
		{[][]int{{1, 5, 6, 4}, {2, 3, 4, 2}, {4, 13, -3, 6}, {2, 3, -4, 7}}, 32},
		{[][]int{{1, 2, 3, 4, 5, 6}, {7, 8, 9, 10, 11, 12}, {13, 14, 15, 16, 17, 18}, {19, 20, 21, 22, 23, 24}, {25, 26, 27, 28, 29, 30}, {31, 32, 33, 34, 35, 36}}, 266},
	}

	t.Run("Maximum Path Sum", func(t *testing.T) {
		for _, tt := range tests {
			got := maximum_path_sum(tt.grid)
			if got != tt.want {
				t.Errorf(ErrorString, got, tt.want)
			}
		}
	})

}
