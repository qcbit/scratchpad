// Problem

// N Choose K Combinations
// Given two integers n and k, find all the possible unique combinations of k numbers in range 1 to n.

// Example One
// {
// "n": 5,
// "k": 2
// }
// Output:

// [
// [1, 2],
// [1, 3],
// [1, 4],
// [1, 5],
// [2, 3],
// [2, 4],
// [2, 5],
// [3, 4],
// [3, 5],
// [4, 5]
// ]
// Example Two
// {
// "n": 6,
// "k": 6
// }
// Output:

// [
// [1, 2, 3, 4, 5, 6]
// ]
// Notes
// The answer can be returned in any order.

// Constraints:

// 1 <= n <= 20
// 1 <= k <= n
package nchoosekcombinations


func find_combinations(n int, k int) [] [] int {
	results := make([][]int, 0)
	helper(n, k, 0, []int{}, &results)
	return results
}

func helper(n, k, ndx int, slate []int, results *[][]int) {
	if len(slate) == k {
		temp := make([]int, len(slate))
		copy(temp, slate)
		*results = append(*results, temp)
		return
	}

	for i := ndx; i < n; i++ {
		slate = append(slate, i+1)
		helper(n, k, i+1, slate, results)
		slate = slate[:len(slate)-1]
	}
}