// Problem

// Search A Node In Binary Search Tree
// Find whether a node with a given value is present in a given binary search tree or not.

// Example One
// binary search tree

// {
// "value": 4
// }
// Output:

// 1
// Example Two
// binary tree

// {
// "value": 5
// }
// Output:

// 0
// Notes
// Constraints:

// 0 <= number of nodes <= 2 * 104
// -105 <= value of a binary tree node <= 105
// -105 <= given value <= 105
// All the nodes in the binary search tree have distinct values
package searchnodeinbst

type BinaryTreeNode struct {
	value int
	left *BinaryTreeNode
	right *BinaryTreeNode
}

func search_node_in_bst(root *BinaryTreeNode, value int) bool {

	for root != nil {
		if root.value == value {
			return true
		} else if value < root.value {
			root = root.left
		} else {
			root = root.right
		}
	}
	return false
}
