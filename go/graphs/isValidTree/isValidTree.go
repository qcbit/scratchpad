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
    for i := 0; i < len(edgeStart); i++ {
        adjList[edgeStart[i]] = append(adjList[edgeStart[i]], edgeEnd[i])
        adjList[edgeEnd[i]] = append(adjList[edgeEnd[i]], edgeStart[i])
    }
    return &Graph{adjList: adjList, n: nodeCount}
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

func is_it_a_tree(node_count int, edge_start []int, edge_end []int) bool {
    g := NewGraph(node_count, edge_start, edge_end)
    visited := make([]bool, node_count)
    parent := make([]int, node_count)
    var components int
    for v := 0; v < node_count; v++ {
        if !visited[v] {
            components++
            // disconnected components means not a valid tree
            if components > 1 {
                return false
            }
            if crossEdgeDetected := bfs(v, g.adjList, visited, parent); crossEdgeDetected {
                return false
            }
        }
    }

    return true
}

// bfs runs the BFS algorithm and checks if is/is not a tree
func bfs(s int, adjList [][]int, visited []bool, parent []int) bool {
    q := NewQueue()
    q.Enqueue(s)
    visited[s] = true
    for len(*q) > 0 {
        u, _ := q.Dequeue()
        for _, v := range adjList[u] {
            if !visited[v] {
                visited[v] = true
                parent[v] = u
                q.Enqueue(v)
            } else {
                if v != parent[u] {
                    return true // cross edge
                }
                if v == u {
                    return true // self loop
                }
            }
        }
    }
    return false
}

