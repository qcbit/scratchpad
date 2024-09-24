// Problem

// Is It A Tree
// Given an undirected graph, find out whether it is a tree.

// The undirected edges are given by two arrays edge_start and edge_end of equal size. Edges of the given graph connect nodes edge_start[i] and edge_end[i] for all valid i.

// Example One
// Graph 1

// {
// "node_count": 4,
// "edge_start": [0, 0, 0],
// "edge_end": [1, 2, 3]
// }
// Output:

// 1
// This graph is a tree because all the nodes are connected, and it does not have cycles.

// Example Two
// Graph 2

// {
// "node_count": 4,
// "edge_start": [0, 0],
// "edge_end": [1, 2]
// }
// Output:

// 0
// This graph is not a tree because node 3 is not connected to the other nodes.

// Example Three
// Graph 3

// {
// "node_count": 4,
// "edge_start": [0, 0, 1, 2],
// "edge_end": [3, 1, 2, 0]
// }
// Output:

// 0
// This graph is not a tree: nodes 0, 1 and 2 form a cycle.

// Example Four
// Graph 4

// {
// "node_count": 4,
// "edge_start": [0, 0, 0, 1],
// "edge_end": [1, 2, 3, 0]
// }
// Output:

// 0
// This graph is not a tree because the two edges that connect nodes 0 and 1 form a cycle.

// Notes
// A tree is an undirected connected acyclic graph.
// Constraints:

// 1 <= number of nodes <= 105
// 1 <= number of edges <= 105
// 0 <= node value < number of nodes
package isvalidtree

type Graph struct {
	adjList [][]int
	n       int
}

func NewGraph(nodeCount int, edgeStart, edgeEnd []int) *Graph {
	adjList := make([][]int, nodeCount)
	for i := 0; i < nodeCount; i++ {
		adjList[edgeStart[i]] = append(adjList[edgeStart[i]], edgeEnd[i])
		adjList[edgeEnd[i]] = append(adjList[edgeEnd[i]], edgeStart[i])
	}
	return &Graph{adjList: adjList, n: nodeCount}
}

type Queue []int

func is_it_a_tree(node_count int, edge_start []int, edge_end []int) bool {
	g := NewGraph(node_count, edge_start, edge_end)
	visited := make([]int, node_count)
	parent := make([]int, node_count)
	return false
}
