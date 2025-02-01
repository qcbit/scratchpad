// Problem
  
// Zigzag Level Order Traversal Of A Binary Tree
// Given a binary tree, return the zigzag level order traversal of the node values listing the odd levels from left to right and the even levels from right to left.

// Example One
// Example one

// Output:

// [
// [0],
// [2, 1],
// [3, 4]
// ]
// Example Two
// Example two

// Output:

// [
// [0],
// [1],
// [2],
// [3]
// ]
// Notes
// Root node is considered to be at the level 1.

// Constraints:

// 1 <= number of nodes in the given tree <= 20000
// 0 <= node value < number of nodes
// Node values are unique
package zigzaglevelordertraversal

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
	var levelFlag bool
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
		if levelFlag {
			for i, j := 0, len(level)-1; i < j; i, j = i+1, j-1 {
				level[i], level[j] = level[j], level[i]
			}
		}
		levelFlag = !levelFlag
		
		results = append(results, level)
	}
	return results
}