// Problem

// Permute Array Of Unique Integers
// Given an array of unique numbers, return in any order all its permutations.

// Example
// {
// "arr": [1, 2, 3]
// }
// Output:

// [
// [1, 2, 3],
// [1, 3, 2],
// [2, 1, 3],
// [2, 3, 1],
// [3, 2, 1],
// [3, 1, 2]
// ]
// Notes
// Constraints:

// 1 <= size of the input array <= 9
// 0 <= any array element <= 99
package permuteArrayOfUniqueIntegers

func get_permutations(arr []int) [][]int {
	results := make([][]int, 0)
	helper(0, arr, &results)
	return results
}

func helper(curr int, arr []int, results *[][]int) {
	if curr == len(arr) {
		temp := make([]int, len(arr))
		copy(temp, arr)
		*results = append(*results, temp)
		return
	}
	
	for i := curr; i < len(arr); i++ {
		arr[i], arr[curr] = arr[curr], arr[i]
		helper(curr+1, arr, results)
		arr[i], arr[curr] = arr[curr], arr[i]
	}
}
