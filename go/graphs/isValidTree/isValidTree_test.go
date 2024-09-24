package isvalidtree

import "testing"

const ErrorStr = `got %v want %v`

func TestIsValidTree(t *testing.T) {
	t.Helper()
	tests := []struct {
		nodeCount int
		edgeStart  []int
		edgeEnd    []int
		want       int
	}{
		{4, []int{0, 0, 0}, []int{1, 2, 3}, 1},
		{4, []int{0, 0}, []int{1, 2}, 0},
		{4, []int{0, 0, 1, 2}, []int{3, 1, 2, 0}, 0},
		{4, []int{0, 0, 0, 1}, []int{1, 2, 3, 0}, 0},
	}
	t.Run("Is Valid Tree", func(t *testing.T) {
		for _, tt := range tests {
			got := is_it_a_tree(tt.nodeCount, tt.edgeStart, tt.edgeEnd)
			if got != tt.want {
				t.Errorf(ErrorStr, got, tt.want)
			}
		}
	})
}
