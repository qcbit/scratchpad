// Problem

// Is It A BST
// Given a binary tree, check if it is a binary search tree (BST). A valid BST does not have to be complete or balanced.

// Consider this definition of a BST:

// All nodes values of left subtree are less than or equal to parent node value.
// All nodes values of right subtree are greater than or equal to parent node value.
// Both left subtree and right subtree must be BSTs.
// NULL tree is a BST.
// Single node trees (including leaf nodes of any tree) are BSTs.
// Example One
// Example one

// Output:

// 0
// Left child value 200 is greater than the parent node value 100; violates the definition of BST.

// Example Two
// Example two

// Output:

// 1
// Notes
// Return true if the input tree is a BST or false otherwise.
// Constraints:

// 0 <= number of nodes <= 100000
// -109 <= values stored in the nodes <= 109
package isitabst

type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func is_bst(root *BinaryTreeNode) bool {
	isBst := true
	if root == nil {
		return true
	}
	if root.left != nil {
		isBst := checkIfBst(root.left, root.value, true, false)
		if !isBst {
			return false
		}
	}
	if root.right != nil {
		isBst := checkIfBst(root.right, root.value, false, true)
		if !isBst {
			return false
		}
	}
	return isBst
}

func checkIfBst(node *BinaryTreeNode, rootVal int, left, right bool) bool {
	isBst := true
	if node == nil {
		return true
	}
	if left && node.value > rootVal {
		return false
	}
	if right && node.value < rootVal {
		return false
	}
	if node.left != nil {
		if node.left.value > node.value {
			return false
		}
		isBst = checkIfBst(node.left, rootVal, left, right)
	}
	if node.right != nil {
		if node.right.value < node.value {
			return false
		}
		isBst = checkIfBst(node.right, rootVal, left, right)
	}

	if !isBst {
		return false
	}

	return isBst
}
