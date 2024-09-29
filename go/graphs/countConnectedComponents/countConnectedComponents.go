// Problem

// Count Connected Components In An Undirected Graph
// Given an undirected graph, find its total number of connected components.

// Example One
// Graph

// {
// "n": 5,
// "edges": [[0 ,1], [1, 2], [0, 2], [3, 4]]
// }
// Output:

// 2
// Example Two
// Graph

// {
// "n": 4,
// "edges": [[0 , 1], [0 , 3], [0 , 2], [2 , 1], [2 , 3]]
// }
// Output:

// 1
// Notes
// Constraints:

// 1 <= number of nodes <= 105
// 0 <= number of edges <= 105
// 0 <= node value < number of nodes
package countconnectedcomponents

type Graph struct {
	adjList [][]int
	n       int
}

func NewGraph(n int, edges [][]int) *Graph {
	adjList := make([][]int, n)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}
	return &Graph{adjList: adjList, n: n}
}

type Queue []int

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(x int) {
	*q = append(*q, x)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(*q) == 0 {
		return -1, false
	}
	x := (*q)[0]
	*q = (*q)[1:]
	return x, true
}

func number_of_connected_components(n int, edges [][]int) int {
	g := NewGraph(n, edges)
	var components int
	visited := make([]bool, n)

	for v := 0; v < n; v++ {
		if !visited[v] {
			components++
			bfs(v, visited, g.adjList)
		}
	}
	return components
}

func bfs(s int, visited []bool, adjList [][]int) {
	visited[s] = true
	q := NewQueue()
	q.Enqueue(s)
	for len(*q) > 0 {
		if u, ok := q.Dequeue(); ok {
			for _, v := range adjList[u] {
				if !visited[v] {
					visited[v] = true
					q.Enqueue(v)
				}
			}
		}
	}
}
