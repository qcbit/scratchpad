package singlenumber

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
	}

	t.Run("Find single number", func(t *testing.T) {
		for _, tt := range tests {
			got := find_single_number(tt.nums)
			if got != tt.want {
				t.Errorf("got %v want %v for %v", got, tt.want, tt.nums)
			}
		}
	})
}
