// Problem

// Single Value Tree
// Given a binary tree, find the number of unival subtrees. An unival (single value) tree is a tree that has the same value in every node.

// Example One
// Example one

// Output:

// 6
// The input tree has a total of 6 nodes. Each node is a root of a subtree. All those 6 subtrees are unival trees.

// Example Two
// Example two

// Output:

// 5
// The input tree has a total of 7 nodes, so there are 7 subtrees. Of those 7, all but two subtrees are unival. The two non-unival subtrees are:

// The one rooted in the root node and
// The one rooted in the root's right child.
// Notes
// Constraints:

// 0 <= number of nodes in the tree <= 105
// -109 <= node value <= 109// // // // // // // // // // // // // // // // // // //
package singlevaluetree

type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func find_single_value_trees(root *BinaryTreeNode) int {
	if root == nil { // empty tree, edge case
		return 0
	}
    uniVals, _ := getUniVals(root)
	return uniVals
}

func getUniVals(node *BinaryTreeNode) (int, bool) {
    // leaf node is always unival
	if node.left == nil && node.right == nil {
		return 1, true
	}

	isUniVal := true
	var lcount, rcount int
	if node.left != nil {
	    var lisUniVal bool
	    lcount, lisUniVal = getUniVals(node.left)
	    if !lisUniVal || node.left.value != node.value {
		    isUniVal = false
	    }
	}
	
	if node.right != nil {
	    var risUniVal bool
	    rcount, risUniVal = getUniVals(node.right)
	    if !risUniVal || node.right.value != node.value {
		    isUniVal = false
	    }
	}

	if !isUniVal {
		return lcount + rcount, false
	}
	return lcount + rcount + 1, true
}