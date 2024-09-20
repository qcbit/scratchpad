package graphs

type AdjacencyMatrix struct {
	adjMtrx [][]bool
	n       int
}

func NewAdjMatrix(n int) AdjacencyMatrix {
	adjMtrx := make([][]bool, n)
	for i := 0; i < n; i++ {
		adjMtrx[i] = make([]bool, n)
	}
	return AdjacencyMatrix{
		adjMtrx: adjMtrx,
	}
}

func (am *AdjacencyMatrix) AddEdge(u, v int) {
	am.adjMtrx[u][v] = true
	am.adjMtrx[v][u] = true
}

func (am *AdjacencyMatrix) GetNeighbors(u, v int) [][]int {
	neighbors := make([][]int, 0)
	if u+1 < am.n {
		if (*am).adjMtrx[u+1][v] {
			neighbors = append(neighbors, []int{u + 1, v})
		}
	}
	if u-1 >= 0 {
		if (*am).adjMtrx[u-1][v] {
			neighbors = append(neighbors, []int{u - 1, v})
		}
	}
	if v+1 < am.n {
		if (*am).adjMtrx[u][v+1] {
			neighbors = append(neighbors, []int{u, v + 1})
		}
	}
	if v-1 >= 0 {
		if (*am).adjMtrx[u][v-1] {
			neighbors = append(neighbors, []int{u, v - 1})
		}
	}
	return neighbors
}
