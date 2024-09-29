// Problem

// Check If Eulerian Cycle Exists
// Check if there exists any eulerian cycle in a given undirected connected graph. The Euler cycle is a path in the graph that visits every edge exactly once and starts and finishes at the same vertex.

// Example One
// Graph

// {
// "n": 5,
// "edges": [
// [0, 1],
// [0, 2],
// [1, 3],
// [3, 0],
// [3, 2],
// [4, 3],
// [4, 0]
// ]
// }
// Output:

// 1
// For example, the graph has an Eulerian Cycle, [2, 0, 1, 3, 0, 4, 3, 2].

// Example Two
// Graph

// {
// "n": 6,
// "edges": [
// [0, 4],
// [0, 5],
// [1, 2],
// [2, 3],
// [3, 1],
// [4, 3],
// ]
// }
// Output:

// 0
// Notes
// The graph has n vertices, with each vertex having a distinct value from 0 to n - 1.
// Edges are given as a list of lists where each inner list has exactly two elements. Each list [X, Y] represents an undirected edge from X to Y.
// The list won't contain any duplicate edges i.e. if [X, Y] is present, then there will be no other occurrence of [X, Y] or [Y, X].
// Constraints:

// 1 <= n <= 103
// 0 <= value of each vertex <= n - 1
// 0 <= number of edges <= (n * (n - 1)) / 2
// The graph won't contain self loops.
package checkifeuleriancycleexists

type Graph struct {
    adjList [][]int
    n   int
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

func (q *Queue) Dequeue() (bool, int) {
    if len(*q) == 0 {
        return false, -1
    }
    x := (*q)[0]
    *q = (*q)[1:]
    return true, x
}


func check_if_eulerian_cycle_exists(n int, edges [] [] int) bool {
    g := NewGraph(n, edges)
    visited := make([]bool, n)
    if n == 1 {
        return true
    }
    for i := 0; i < n; i++ {
        if !visited[i] {
            if eulerianCycleExists := bfs(i, visited, g.adjList); eulerianCycleExists {
                return true
            }
        }
    }
	return false
}

type Odd struct {
    nodes []int
}

func bfs(s int, visited []bool, adjList [][]int) bool {
    var even int
    var odd Odd
    var onePath int
    q := NewQueue()
    q.Enqueue(s)
    visited[s] = true
    for len(*q) > 0 {
        _, u := q.Dequeue()
        for _, v := range adjList[u] {
            if !visited[v] {
                visited[v] = true
                q.Enqueue(v)
            }
        }
        if len(adjList[u]) == 1 {
            onePath++
        }
        if len(adjList[u]) % 2 == 0 {
            even++
        } else if len(adjList[u]) > 2 && len(adjList[u]) % 2 == 1 {
                odd.nodes = append(odd.nodes, u)
        }
    }
    if even >= 2 && len(odd.nodes) == 0 && onePath == 0{
        return true
    } else if len(odd.nodes) == 2 && onePath == 0{
        return true
    }
    return false
}
