// Problem

// Postorder Traversal Of A Binary Tree
// Given a binary tree, find its postorder traversal.

// Example One
// Example one

// Output:

// [3, 4, 1, 2, 0]
// Example Two
// Example two

// Output:

// [3, 2, 1, 0]
// Notes
// The postorder traversal visits all the nodes of a binary tree by recursively visiting the left subtree, then the right subtree and finally visiting the root.

// Constraints:

// 1 <= number of nodes in the tree <= 20000
// 0 <= node value < number of nodes
// No two nodes have the same value
package postordertraversal

type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func postorder(root *BinaryTreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

func dfs(node *BinaryTreeNode, result *[]int) {
	if node == nil {
		return
	}
	dfs(node.left, result)
	dfs(node.right, result)
	*result = append(*result, node.value)
}
