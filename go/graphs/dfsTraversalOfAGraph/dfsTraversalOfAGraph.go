// Problem

// DFS Traversal Of A Graph
// Given an undirected graph, perform a Depth-First Search Traversal on it.

// Example
// Graph

// {
// "n": 6,
// "edges": [
// [0, 1],
// [0, 2],
// [1, 4],
// [3, 5]
// ]
// }
// Output:

// [0, 1, 4, 2, 3, 5]
// Below are some other valid outputs if the DFS traversal starts from the vertex 0:

// [0, 2, 1, 4, 5, 3]
// [0, 2, 1, 4, 3, 5]
// [0, 1, 4, 2, 5, 3]
// DFS starting from any other node will also be considered valid.

// Notes
// There are n nodes in the graph and each node has a distinct value from 0 to n - 1.
// Edges are represented in the form of a two-dimensional list where each inner list consists of two values [u, v]. This represents an undirected edge from node u to node v.
// The list won't contain duplicate edges. That is, if [u, v] is present, then there will be no other occurrence of [u, v] or [v, u].
// If multiple DFS traversals exist, you may return any.
// Constraints:

// 1 <= n <= 104
// 0 <= number of edges <= 104
// 0 <= value of each vertex <= n - 1
// No self loop is present.
package dfstraversalofagraph

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

type Stack struct {
	stack []int
}

func NewStack() *Stack {
	return &Stack{stack: make([]int, 0)}
}

func (s *Stack) Push(x int) {
	s.stack = append(s.stack, x)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.stack) == 0 {
		return 0, false
	}
	x := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return x, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

func dfs_traversal(n int, edges [][]int) []int {
	g := NewGraph(n, edges)
	result := make([]int, 0)
	visited := make([]bool, n)
	dfs(n, g.adjList, 0, visited, &result)
	return result
}

func dfs(n int, adjList [][]int, s int, visited []bool, result *[]int) {
	for i := s; i < n; i++ {
		if !visited[i] {
			visited[i] = true
			stack := NewStack()
			stack.Push(i)
			for !stack.IsEmpty() {
				if u, ok := stack.Pop(); ok {
					*result = append(*result, u)
					for _, v := range adjList[u] {
						if !visited[v] {
							visited[v] = true
							stack.Push(v)
						}
					}
				}
			}
		}
	}
}