package circulararrayloop

import (
	"testing"
)

func TestCircularArrayLoop(t *testing.T) {
	tests := []struct {
		list []int
		want bool
	}{
		{[]int{2,-1,1,2,2}, true},
		{[]int{-1, 2}, false},
		{[]int{-2,1,-1,-2,-2}, false},
	}

	for _, tt := range tests {
		got := detectCycle(tt.list)
		if got != tt.want {
			t.Errorf("got %v want %v for %v", got, tt.want, tt.list)
		}
	}
}
