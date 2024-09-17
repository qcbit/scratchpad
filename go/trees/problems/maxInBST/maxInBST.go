// Problem

// Maximum Valued Node In Binary Search Tree
// Find the node with maximum value in a given binary search tree and return its value.

// Example
// Binary-Search-Tree

// Output:

// 6
// Notes
// Constraints:

// 1 <= number of nodes <= 2 * 104
// -105 <= value of a binary search tree node <= 105
// All values in the binary search tree will be unique.
package maxinbst

type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func get_maximum_value(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	curr := root
	for curr.right != nil {
		curr = curr.right
	}
	return curr.value
}
