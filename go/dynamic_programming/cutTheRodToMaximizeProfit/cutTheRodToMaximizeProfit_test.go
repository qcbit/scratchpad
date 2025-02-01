package cuttherodtomaximizeprofit

import (
	"testing"
)

const ErrorString = `got %v want %v`

func TestCutTheRodToMaximizeProfit(t *testing.T) {
	t.Helper()
	tests := []struct {
		price []int
		want  int
	}{
		{[]int{1, 5, 8, 9}, 10},
		{[]int{3, 7, 2, 6, 6}, 17},
		{[]int{3, 3, 3, 3, 3}, 15},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{5, 4, 3, 2, 1}, 25},
		{[]int{3, 12, 1, 1, 1, 3}, 36},
		{[]int{7, 8, 9, 9, 2, 8}, 42},
		{[]int{2, 3, 4, 7, 5, 1, 10}, 14},
		{[]int{5, 10, 2, 1, 8, 13, 6, 7}, 40},
		{[]int{2, 2, 2, 3, 3, 3, 4, 4, 4}, 18},
		{[]int{10, 1, 100, 10, 10}, 120},
	}

	t.Run("Cut the Rod to Maximize Profits", func(t *testing.T) {
		for _, tt := range tests {
			got := get_maximum_profit(tt.price)
			if got != tt.want {
				t.Errorf(ErrorString, got, tt.want)
			}
		}
	})

}
