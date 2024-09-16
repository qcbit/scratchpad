// Problem

// Level Order Traversal Of A Tree
// Given a tree, list node values level by level from left to right.

// Example One
// Example one

// Output:

// [
// [1],
// [3, 4, 2],
// [5, 6]
// ]
// Example Two
// Example two

// Output:

// [
// [1],
// [2],
// [4],
// [3]
// ]
// Notes
// Constraints:

// 1 <= number of nodes <= 20000
// 1 <= value in a node <= number of nodes
// Node values are unique
// Root node's value is 1
package levelordertraversalnary

type Queue []interface{}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(node interface{}) {
	*q = append(*q, node)
}

func (q *Queue) Dequeue() (interface{}, bool) {
	if len(*q) == 0 {
		return nil, false
	}
	node := (*q)[0]
	(*q) = (*q)[1:]
	return node, true
}

type TreeNode struct {
	value    int
	children []*TreeNode
}

func level_order(root *TreeNode) [][]int {
	results := make([][]int, 0)

	q := NewQueue()
	q.Enqueue(root)
	for len(*q) > 0 {
		level := make([]int, 0)
		numNodesPerLevel := len(*q)
		for i := 0; i < numNodesPerLevel; i++ {
			node, _ := q.Dequeue()
			level = append(level, node.(*TreeNode).value)
			for _, c := range node.(*TreeNode).children {
				q.Enqueue(c)
			}
		}
		results = append(results, level)
	}

	return results
}
