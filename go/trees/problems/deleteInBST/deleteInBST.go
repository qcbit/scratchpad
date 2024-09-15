// Problem

// Delete From BST
// Given a Binary Search Tree (BST) and a list of numbers, your task is to delete the nodes whose value matches any number from the given list.

// Example One
// Graph

// {
// "values_to_be_deleted": [5, 6]
// }
// Output:

// Graph

// Example Two
// Graph

// {
// "values_to_be_deleted": [-1, 0, 8, 9]
// }
// Output:

// Graph

// Notes
// You should ignore those numbers which do not match with any value of the BST nodes.
// If there are more than one output BSTs possible, you can return any of them.
// Constraints:

// 1 <= number of nodes in the BST <= 104
// -109 <= value of any node <= 109
// All the nodes in the BST have unique values.
// 1 <= length of the given list of numbers <= 104
package deleteinbst

type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func delete_from_bst(root *BinaryTreeNode, values_to_be_deleted []int) *BinaryTreeNode {
	for _, v := range values_to_be_deleted {
		root = delete(root, v)
	}
	return root
}

func delete(root *BinaryTreeNode, k int) *BinaryTreeNode {
	curr := root
	var prev *BinaryTreeNode

	// search
	for curr != nil {
		if k == curr.value {
			break
		} else if k < curr.value {
			prev = curr
			curr = curr.left
		} else {
			prev = curr
			curr = curr.right
		}
	}

	if curr == nil {
		return root
	}

	// leaf node
	if curr.left == nil && curr.right == nil {
		if prev == nil { // special case if just root
			return nil
		}
		if curr == prev.left { // delete left child that is a leaf
			prev.left = nil
		} else {
			prev.right = nil // delete right child that is a leaf
		}
	}

	// node to delete has a child
	var child *BinaryTreeNode
	if curr.left == nil && curr.right != nil {
		child = curr.right // identify right child
	}
	if curr.right == nil && curr.left != nil {
		child = curr.left // identify left child
	}
	if child != nil {
		if prev == nil { // special case of deleting the root
			root = child
			return root
		}
		// deletion of node, but keeping descendants
		if curr == prev.left {
			prev.left = child
		} else {
			prev.right = child
		}
	}

	// node to delete has children
	if curr.left != nil && curr.right != nil {
		succ := curr.right
		prev = curr
		// find successor
		for succ.left != nil {
			prev = succ
			succ = succ.left
		}
		// copy value of succ to curr
		curr.value = succ.value
		// delete
		if succ == prev.left {
			prev.left = succ.right
		} else {
			prev.right = succ.right
		}
	}

	return root
}
