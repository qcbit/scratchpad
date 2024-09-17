// Problem

// Insert In BST
// Given an array of numbers, build a binary search tree(BST) by inserting the values sequentially inside an initially empty BST.

// Example
// {
// "values": [7, 5, 9]
// }
// Output:

// Graph

// After Inserting 7:
// Graph

// After inserting 5:
// Graph

// Finally after inserting 9:
// Graph

// Notes
// The values should be inserted in the left-to-right order described by the input array.
// While inserting any new node, the relative structure and the position of all the other nodes should not be altered. For example, if we have the following BST:
// Graph
// and are given to insert a new node with value 9, then the following will not be a valid output as the position of the nodes with value 5 and 7 are getting changed:
// Graph
// Constraints:

// 1 <= length of the input array <= 103
// -104 <= values in the array <= 104
// All of the given values are distinct.
package insertinbst

type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func build_a_bst(values []int) *BinaryTreeNode {
	if len(values) == 0 {
		return &BinaryTreeNode{}
	}
	root := &BinaryTreeNode{values[0], nil, nil}
	for i := 1; i < len(values); i++ {
		insert(root, values[i])
	}
	return root
}

func insert(node *BinaryTreeNode, k int) *BinaryTreeNode {
	newNode := &BinaryTreeNode{value: k}
	if node == nil {
		return newNode
	}
	var prev *BinaryTreeNode
	curr := node
	for curr != nil {
		if k == curr.value {
			return node
		}
		prev = curr
		if k < curr.value {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	if k < prev.value {
		prev.left = newNode
	} else {
		prev.right = newNode
	}

	return node
}
