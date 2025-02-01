// Problem

// Inorder Traversal Of A Binary Tree
// Given a binary tree, return the inorder traversal of its node values.

// Example One
// Example one

// Output:

// [3, 1, 4, 0, 2]
// Example Two
// Example two

// Output:

// [1, 3, 2, 0]
// Notes
// The inorder traversal of a binary tree first visits the left subtree, then the root and finally the right subtree.

// Constraints:

// 1 <= number of nodes in the given tree <= 20000
// 0 <= node value < number of nodes
// Node values are unique

package inordertraversal

type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func inorder(root *BinaryTreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

func dfs(node *BinaryTreeNode, result *[]int) {
	if node == nil {
		return
	}
	dfs(node.left, result)
	*result = append(*result, node.value)
	dfs(node.right, result)
}
