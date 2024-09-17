// Problem

// Height Of A Binary Tree
// Find the height of a given binary tree.

// Example
// Binary Tree

// Output:

// 4
// Notes
// The height of a binary tree is the number of nodes along the longest path from the root node down to the farthest leaf node.

// Constraints:

// 0 <= number of nodes <= 2 * 104
// -103 <= value of a binary tree node <= 103
package heightofabinarytree

import "math"

type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func height_of_binary_tree(root *BinaryTreeNode) int {
	return getHeightOfTree(root)
}

func getHeightOfTree(node *BinaryTreeNode) int {
	if node == nil {
		return 0
	}
	return int(math.Max(float64(getHeightOfTree(node.left)), float64(getHeightOfTree(node.right)))) + 1
}
