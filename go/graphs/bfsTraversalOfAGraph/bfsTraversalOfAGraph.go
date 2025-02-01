// Problem

// BFS Traversal Of A Graph
// Given an undirected graph, perform a Breadth-First Search Traversal on it.

// Example
// Graph

// {
// "n": 6,
// "edges": [
// [0, 1],
// [0, 2],
// [0, 4],
// [2, 3]
// ]
// }
// Output:

// [0, 1, 2, 4, 3, 5]
// Below are some other following valid outputs if the BFS traversal starts from the vertex 0:

// [0, 1, 4, 2, 3, 5]
// [0, 2, 1, 4, 3, 5]
// [0, 2, 4, 1, 3, 5]
// [0, 4, 1, 2, 3, 5]
// [0, 4, 2, 1, 3, 5]
// BFS starting from another node will also be considered valid.

// Notes
// There are n nodes in the graph and each node has a distinct value from 0 to n-1.
// Edges are represented in the form of a two-dimensional list where each inner list consists of two values [u, v]. This represents an undirected edge from node u to node v.
// The list won't contain duplicate edges. That is, if [u, v] is present, then there will be no other occurrence of [u, v] or [v, u].
// If multiple BFS traversals exist, you may return any.
// Constraints:

// 1 <= n <= 104
// 0 <= number of edges <= 104
// 0 <= value of each vertex <= n - 1
// No self loop is present.
package bfstraversalofagraph

type Graph struct {
	adjList [][]int
	n       int
}

func NewGraph(edges [][]int, n int) *Graph {
	adjList := make([][]int, n)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}
	return &Graph{adjList: adjList, n: n}
}

func bfs_traversal(n int, edges [][]int) []int {
	g := NewGraph(edges, n)
	result := make([]int, 0)
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visited[i] {
			visited[i] = true
			queue := []int{i}
			for len(queue) > 0 {
				u := queue[0]
				result = append(result, u)
				queue = queue[1:]
				for _, v := range g.adjList[u] {
					if !visited[v] {
						visited[v] = true
						queue = append(queue, v)
					}
				}
			}
		}
	}
	return result
}