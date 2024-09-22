// Problem
  
// Symmetric Tree
// Given a binary tree, check whether it is a mirror of itself i.e. symmetric around its centre.

// Example One
// Tree

// Output:

// 1
// Example Two
// Tree

// Output:

// 0
// Example Three
// Tree

// Output:

// 0
// Notes
// Constraints:

// 0 <= number of nodes in the tree <= 5 * 104
// 0 <= value in a tree node <= 105
package symmetrictree

import "strconv"

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

type BinaryTreeNode struct {
    value int
    left *BinaryTreeNode
    right *BinaryTreeNode
}
// [1, 5, 5, 7, null, null, 7]
// [100, 90, 90, 70, 60, 60, 70]
// [78,null,67,null,55,null,63]
// [100, 33,33,22,22,44,44,5,null,null,5,7,null,null,7]
func check_if_symmetric(root *BinaryTreeNode) bool {
	if root == nil {
    return true
  }
	queue := NewQueue()
	queue.Enqueue(root)
	for len(*queue) > 0 {
		level := []string{}
		numNodesPerLevel := len(*queue)
		for q := 0; q < numNodesPerLevel; q++ {
			v, _ := queue.Dequeue()
			if v.(*BinaryTreeNode).left != nil {
				level = append(level, strconv.Itoa(v.(*BinaryTreeNode).left.value))
				queue.Enqueue(v.(*BinaryTreeNode).left)
			} else {
				level = append(level, "null")
			}
			if v.(*BinaryTreeNode).right != nil {
				level = append(level, strconv.Itoa(v.(*BinaryTreeNode).right.value))
				queue.Enqueue(v.(*BinaryTreeNode).right)
			} else {
				level = append(level, "null")
			}
		}
		if !checkSymmetry(level) {
			return false
		}
	}
	return true
}

func checkSymmetry(ary []string) bool {
	for i, j := 0, len(ary)-1; i < j; i, j = i+1, j-1 {
		if ary[i] != ary[j] {
			return false
		}
	}
	return true
}