package kthsmallest

import "testing"

func TestKthSmallest(t *testing.T) {
	testCases := []struct {
		input []int
		k     int
		want  int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 2},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 3},
	}

	t.Run("find the kth largest", func(t *testing.T) {
		for _, tt := range testCases {
			got := findKthSmallest(tt.input, tt.k)
			assertCorrect(t, got, tt.want, tt.input)
		}
	})
}

func assertCorrect(t *testing.T, got, want int, test []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v given %v", got, want, test)
	}
}
