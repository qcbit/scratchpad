package countconnectedcomponents

import "testing"

const ErrorStr = `got %v want %v`

func TestCountConnectedComponents(t *testing.T) {
	t.Helper()
	tests := []struct {
		n     int
		edges [][]int
		want int
	}{
		{5, [][]int{{0, 1}, {1, 2}, {0, 2}, {3, 4}}, 2},
		{4, [][]int{{0, 1}, {0, 3}, {0, 2}, {2, 1}, {2, 3}}, 1},
	}
	t.Run("Count Connected Components in a Graph", func(t *testing.T) {
		for _, tt := range tests {
			got := number_of_connected_components(tt.n, tt.edges)
			if got != tt.want {
				t.Errorf(ErrorStr, got, tt.want)
			}
		}
	})
}
