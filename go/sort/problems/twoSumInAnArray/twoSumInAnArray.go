// Problem

// 2 Sum In An Array
// Given an array and a target number, find the indices of the two values from the array that sum up to the given target number.

// Example One
// {
// "numbers": [5, 3, 10, 45, 1],
// "target": 6
// }
// Output:

// [0, 4]
// Sum of the elements at index 0 and 4 is 6.

// Example Two
// {
// "numbers": [4, 1, 5, 0, -1],
// "target": 10
// }
// Output:

// [-1, -1]
// Notes
// The function returns an array of size two where the two elements specify the input array indices whose values sum up to the given target number.
// In case when no answer exists, return an array of size two with both values equal to -1, i.e., [-1, -1].
// In case when multiple answers exist, you may return any of them.
// The order of the returned indices does not matter.
// A single index cannot be used twice.
// Constraints:

// 2 <= array size <= 105
// -105 <= array elements <= 105
// -105 <= target number <= 105
// Array can contain duplicate elements.
package twoSumInAnArray

import (
	sort "github.com/qcbit/scratchpad/go/sort"
)

func two_sum(numbers []int, target int) []int {
	sorted_array := sort.MergeSort(numbers)
	i, j := 0, len(sorted_array) - 1
	for i < j {
		if sorted_array[i] + sorted_array[j] == target {
			return []int{i,j}
		} else if sorted_array[i] + sorted_array[j] < target {
			i++
		} else {
			j--
		}
	}

	return []int{-1,-1}
}
