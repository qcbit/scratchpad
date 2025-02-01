package kthlargest

import "testing"

func TestKthLargest(t *testing.T) {
	testCases := []struct {
		input []int
		k     int
		want  int
	}{
		{[]int{3, 2, 1, 5, 6, 1}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}

	t.Run("find the kth largest", func(t *testing.T) {
		for _, tt := range testCases {
			got := findKthLargest(tt.input, tt.k)
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
