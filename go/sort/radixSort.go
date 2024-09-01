// Problem

// Implement Radix Sort
// Given a list of numbers, sort it using the Radix Sort algorithm.

// Example
// {
// "arr": [5, 8, 3, 9, 4, 1, 7]
// }
// Output:

// [1, 3, 4, 5, 7, 8, 9]
// Notes
// Constraints:

// 1 <= length of the given list <= 4 * 105
// 0 <= number in the list <= 109
package sort

import (
	"math"
)

// RadixSort
// Time Complexity: O(nk)
// Space Complexity: O(nk)
// Execution time during testing: .55 sec max
// Memory usage during testing: 49.31 MB
func RadixSort(arr []int) []int {
	for i := 0; i < 32; i++ {
		count := make(map[int][]int, 10)
		for i := 0; i < 10; i++ {
			count[i] = make([]int, 0)
		}

		div := int(math.Pow(10, float64(i)))
		for _, num := range arr {
			d := num / div % 10
			count[d] = append(count[d], num)
		}
		aux := make([]int, 0, len(arr))
		for i := 0; i < 10; i++ {
			aux = append(aux, count[i]...)
		}
		copy(arr, aux)
	}

	return arr
}

// RadixSort2 is an improved version of RadixSort
// Time Complexity: O(nk)
// Space Complexity: O(nk)
// Execution time during testing:  .21 sec max
// Memory usage during testing:  49.59 MB
func RadixSort2(arr []int) []int {
	for i := 0; i < 4; i++ {
		count := make([]int, 256)

		for _, num := range arr {
			count[num>>(8*i)&0xFF]++
		}

		for j := 1; j < 256; j++ {
			count[j] += count[j-1]
		}

		aux := make([]int, len(arr))
		for j := len(arr) - 1; j >= 0; j-- {
			aux[count[arr[j]>>(8*i)&0xFF]-1] = arr[j]
			count[arr[j]>>(8*i)&0xFF]--
		}
		copy(arr, aux)
	}

	return arr
}
