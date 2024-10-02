package minimumcoins

import "testing"

const ErrorString = `got %v want %v given coins %v given target value %v`

func TestMinimumCoins(t *testing.T) {
	t.Helper()
	tests := []struct {
		coins []int
		value int
		want  int
	}{
		{[]int{1, 3, 5}, 9, 3},
		{[]int{1, 8, 9, 2, 5}, 15, 3},
		{[]int{1}, 1, 1},
		{[]int{1}, 100, 100},
		{[]int{1, 3}, 116, 40},
		{[]int{22, 14, 1, 18}, 889, 43},
		{[]int{3, 2, 1, 7, 5, 6, 4}, 999, 143},
		{[]int{1, 5, 3, 7}, 888, 128},
		{[]int{12, 14, 8, 6, 1, 4, 16, 2, 10, 18}, 909, 52},
		{[]int{8, 10, 3, 7, 9, 1, 2}, 455, 46},
	}

	t.Run("Minimum Cost Climbing Stairs", func(t *testing.T) {
		for _, tt := range tests {
			got := minimum_coins(tt.coins, tt.value)
			if got != tt.want {
				t.Errorf(ErrorString, got, tt.want, tt.coins, tt.value)
			}
		}
	})

}
