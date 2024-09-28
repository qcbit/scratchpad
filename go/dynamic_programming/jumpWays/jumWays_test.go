package jumpways

import "testing"

const ErrorString = `got %v want %v`

func TestFibonnaci(t *testing.T) {
	t.Helper()
	tests := []struct {
		n    int
		want int64
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{7, 21},
		{10, 89},
		{15, 987},
		{22, 28657},
		{45, 1836311903},
		{60, 2504730781961},
		{70, 308061521170129},
	}

	t.Run("fibonacci", func(t *testing.T) {
		for _, tt := range tests {
			got := jump_ways(tt.n)
			if got != tt.want {
				t.Errorf(ErrorString, got, tt.want)
			}
		}
	})
}
