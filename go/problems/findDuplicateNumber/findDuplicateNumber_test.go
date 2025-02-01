package findduplicatenumber

import (
	"testing"
)

func TestFindDuplicateNumber(t *testing.T) {
	tests := []struct {
		list []int
		want int
	}{
		{[]int{1,3,4,2,2}, 2},
		{[]int{3,1,3,4,2}, 3},
		{[]int{1,1}, 1},
		{[]int{1,1,2}, 1},
	}

	for _, tt := range tests {
		got := detectCycle(tt.list)
		if got != tt.want {
			t.Errorf("got %v want %v for %v", got, tt.want, tt.list)
		}
	}
}
