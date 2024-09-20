// Problem

// Convert Edge List To Adjacency List
// Convert the given edge list to the adjacency list of an undirected connected graph.
// An adjacency list represents a graph as a list of lists. The list index represents a vertex, and each element in its inner list represents the other vertices that form an edge with the vertex.

// Example
// {
// "n": 5,
// "edges": [
// [0, 1],
// [1, 4],
// [1, 2],
// [1, 3],
// [3, 4]
// ]
// }
// Output:

// [
// [1],
// [0, 2, 3, 4],
// [1],
// [1, 4],
// [1, 3]
// ]
// Notes
// There are n nodes in the graph, and each node has a distinct value from 0 to n - 1.
// Edges are given as a list of pairs. Each pair [u, v] represents an undirected edge between node u and node v.
// The list won't contain duplicate edges. That is, if [u, v] is present, then there will be no other occurrence of [u, v] or [v, u].
// Every inner list of the output list should hold the nodes in ascending order.
// Constraints:

// 1 <= n <= 103
// 0 <= number of edges <= (n * (n - 1)) / 2
// 0 <= value of each node <= n - 1
// The graph won't contain self loops.
package edgetoadjacencylistconversion

import "sort"

type Graph struct {
	adjList [][]int
	v       int
}

func NewGraph(n int) *Graph {
	return &Graph{adjList: make([][]int, n), v: n}
}

func (g *Graph) AddEdge(u, v int, bidir bool) {
	g.adjList[u] = append(g.adjList[u], v)
	if bidir {
		g.adjList[v] = append(g.adjList[v], u)
	}
}

func convert_edge_list_to_adjacency_list(n int, edges [][]int) [][]int {
	g := NewGraph(n)
	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1], true)
	}
	for i := 0; i < len(g.adjList); i++ {
		sort.Ints(g.adjList[i])
	}
	return g.adjList
}
