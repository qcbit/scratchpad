// Problem

// Diameter Of A Binary Tree
// Given a binary tree, find its diameter.

// Example One
// Example one

// Output:

// 3
// Example Two
// Example two

// Output:

// 4
// Notes
// Diameter of a binary tree is the length of the longest path between any two nodes of the tree.
// Length between any two nodes is equal to the number of edges traversed to reach one node from the other.
// Constraints:

// 1 <= number of nodes in the given tree <= 105
// 0 <= node value < number of nodes
// Node values are unique
package diameter

import "math"

type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

// {
// "root": [0,
// 1, null,
// 2, 3,
// 4, null, null, 5]
// }
// got: 3
// want:4
func binary_tree_diameter(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	var diameter int
	getDiameter(root, &diameter)
	return diameter
}

func getDiameter(node *BinaryTreeNode, diameter *int) int {
	if node.left == nil && node.right == nil {
		return 0
	}

	var lmax, rmax int
	if node.left != nil {
		lmax = getDiameter(node.left, diameter) + 1
	}
	if node.right != nil {
		rmax = getDiameter(node.right, diameter) + 1
	}

	*diameter = int(math.Max(float64(*diameter), float64(lmax + rmax)))
	return int(math.Max(float64(lmax), float64(rmax)))
}
