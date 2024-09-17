// Problem

// Level Order Traversal Of A Binary Tree
// Given a binary tree, list the node values level by level from left to right.

// Example One
// Example one

// Output:

// [
// [0],
// [1],
// [2],
// [4],
// [3]
// ]
// Example Two
// Example two

// Output:

// [
// [2],
// [5, 4],
// [0, 1, 3, 6]
// ]
// Notes
// Constraints:

// 1 <= number of nodes in the given tree <= 20000
// 0 <= node value < number of nodes
// Node values are unique
package levelordertraversal

type Queue []interface{}

func NewQueue() *Queue {
    return &Queue{}
}

func (q *Queue) Enqueue(element interface{}) {
    *q = append(*q, element)
}

func (q *Queue) Dequeue() (interface{}, bool) {
    if len(*q) == 0 {
        return nil, false
    }
    element := (*q)[0]
    *q = (*q)[1:]
    return element, true
}

//For your reference:

type BinaryTreeNode struct {
		value int
		left *BinaryTreeNode
		right *BinaryTreeNode
}

func level_order_traversal(root *BinaryTreeNode) [][]int {
	results := make([][]int, 0)

	queue := NewQueue()
	queue.Enqueue(root)
	for len(*queue) > 0 {
		level := []int{}
		numNodesPerLevel := len(*queue)
		for q := 0; q < numNodesPerLevel; q++ {
			v, _ := queue.Dequeue()
			level = append(level, v.(*BinaryTreeNode).value)
			if v.(*BinaryTreeNode).left != nil {
				queue.Enqueue(v.(*BinaryTreeNode).left)
			}
			if v.(*BinaryTreeNode).right != nil {
				queue.Enqueue(v.(*BinaryTreeNode).right)
			}
		}
		
		results = append(results, level)
	}
	return results
}