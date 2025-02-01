package graphs

type Graph struct {
	adjList [][]int
	v       int
}

func NewGraph(n int) *Graph {
	return &Graph{adjList: make([][]int, n), v: n}
}

func (g *Graph) HasEulerianCycle() bool {
	var odd int
	for v := range g.adjList {
		if len(g.adjList[v]) % 2 == 1 {
			odd++
		}
	}
	return odd == 0
}

func (g *Graph) HasEulerianPath() bool {
	var odd int
	for v := range g.adjList {
		if len(g.adjList[v]) % 2 == 1 {
			odd++
		}
	}
	return odd == 0 || odd == 2 
}

func (g *Graph) AddEdge(u, v int, bidir bool) {
	g.adjList[u] = append(g.adjList[u], v)
	if bidir {
		g.adjList[v] = append(g.adjList[v], u)
	}
}
