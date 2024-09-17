// Problem

// Preorder Traversal Of A Binary Tree
// Given a binary tree, return node values in the preorder traversal order.

// Example One
// Example one

// Output:

// [0, 1, 3, 4, 2]
// Example Two
// Example two

// Output:

// [0, 1, 2, 3]
// Notes
// The preorder traversal processes all the nodes of a binary tree by first visiting the root, then recursively visiting its left and right subtrees respectively.

// Constraints:

// 1 <= number of nodes in the given tree <= 20000
// 0 <= node value < number of nodes
// Node values are unique
package preordertraversal

type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

// Preorder Traversal
// Time Complexity: O(n)
// Space Complexity: O(n)
func preorder(root *BinaryTreeNode) []int {
	result := make([]int, 0)
    dfs(root, &result)
    return result
}

func dfs(node *BinaryTreeNode, result *[]int) {
    if node == nil {
        return
    }
    *result = append(*result, node.value)
    dfs(node.left, result)
    dfs(node.right, result)
}