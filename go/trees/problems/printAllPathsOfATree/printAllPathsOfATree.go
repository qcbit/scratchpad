// Problem

// Print All Paths Of A Tree
// Given a binary tree, return all paths from root to leaf.

// Example One
// Example one

// Output:

// [
// [1, 2, 4],
// [1, 2, 5],
// [1, 3, 6],
// [1, 3, 7]
// ]
// There are four leafs in the given graph, so we have four paths: from the root to every leaf. Each path is a list of the values from the tree nodes on the path, and we have four lists. They can go in any order.

// Example Two
// Example two

// Output:

// [
// [10, 20, 40],
// [10, 20, 50],
// [10, 30]
// ]
// There are 3 paths in the tree.

// The leftmost path contains values: 10 -> 20 -> 40

// The rightmost path contains values: 10 -> 30

// The other path contains values: 10 -> 20 -> 50

// The order of the paths (order of the lists in the outer list) does not matter, so [[10, 30], [10, 20, 40], [10, 20, 50]] is another correct answer.

// Notes
// Return a list of integer lists, where each list is representing a path.
// The order of the paths (order of the lists in the outer list) does not matter.
// Constraints:

// 0 <= number of nodes in the given tree <= 105
// -109 <= value in a node <= 109
package printallpathsofatree

/*
For your reference:

	type BinaryTreeNode struct {
	    value int
	    left *BinaryTreeNode
	    right *BinaryTreeNode
	}
*/
func all_paths_of_a_binary_tree(root *BinaryTreeNode) [][]int {
	results := make([][]int, 0)
	dfs(root, []int{}, &results)
	return results
}

func dfs(node *BinaryTreeNode, path []int, results *[][]int) {
	if node == nil {
		return
	}

	if node.left == nil && node.right == nil {
		path = append(path, node.value)
		tmp := make([]int, len(path))
		copy(tmp, path)
		*results = append(*results, tmp)
		return
	}

	dfs(node.left, append(path, node.value), results)
	dfs(node.right, append(path, node.value), results)
}
