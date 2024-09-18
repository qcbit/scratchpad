// Problem

// Vertical Order Traversal Of A Binary Tree
// Given a binary tree, find its vertical order traversal starting from the leftmost level to the rightmost level.

// Example
// Tree

// Output:

// [
// [4],
// [2],
// [1, 5, 6],
// [3, 8],
// [7]
// ]
// Tree

// Notes
// If two or more nodes lie in the same vertical level, then the one with earlier occurrence in the level-order traversal of the tree should come first in the output.
// If two or more nodes share the same horizontal and vertical level, then the order should be from left to right.
// Constraints:

// 0 <= number of nodes in the tree <= 2 * 104
// 0 <= value in a tree node <= 105
package verticalordertraversal

type Queue []interface{}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(element interface{}) {
	*q = append(*q, element)
}

func (q *Queue) Dequeue() interface{} {
	if len(*q) == 0 {
		return nil
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element
}

type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func vertical_order_traversal(root *BinaryTreeNode) [][]int {
	xMap := make(map[int][]int)
	if root == nil {
		return [][]int{}
	}

	var x int
	xMap[0] = append(xMap[x], root.value)
	q := NewQueue()
	q.Enqueue(root)
	for len(*q) > 0 {
		numNodesPerLevel := len(*q)
		for i := 0; i < numNodesPerLevel; i++ {
			v, _ := q.Dequeue().(*BinaryTreeNode)
			if v.left != nil {
				xx := x-1	
			}
		}
	}
	return xMap
}
