// Problem

// Reverse Level Order Traversal Of A Binary Tree
// Given a binary tree, return the bottom-up level order traversal of the node values listing each level from left to right.

// Example One
// Example one

// Output:

// [
// [3, 4],
// [1, 2],
// [0]
// ]
// Example Two
// Example two

// Output:

// [
// [3],
// [2],
// [1],
// [0]
// ]
// Notes
// Constraints:

// 1 <= number of nodes in the given tree <= 20000
// 0 <= node value < number of nodes
// Node values are unique
package reverlevelordertraversal

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
	left  *BinaryTreeNode
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

	for i := 0; i < len(results)/2; i++ {
		results[i], results[len(results)-1-i] = results[len(results)-1-i], results[i]
	}
	return results
}
