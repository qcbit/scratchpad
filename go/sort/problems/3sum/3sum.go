// Problem

// 3 Sum
// Given an integer array arr of size n, find all magic triplets in it.

// Magic triplet is a group of three numbers whose sum is zero.

// Note that magic triplets may or may not be made of consecutive numbers in arr.

// Example
// {
// "arr": [10, 3, -4, 1, -6, 9]
// }
// Output:

// ["10,-4,-6", "3,-4,1"]
// Notes
// Function must return an array of strings. Each string (if any) in the array must represent a unique magic triplet and strictly follow this format: "1,2,-3" (no whitespace, one comma between numbers).
// Order of the strings in the array is insignificant. Order of the integers in any string is also insignificant. For example, if ["1,2,-3", "1,-1,0"] is a correct answer, then ["0,1,-1", "1,-3,2"] is also a correct answer.
// Triplets that only differ by order of numbers are considered duplicates, and duplicates must not be returned. For example, if "1,2,-3" is a part of an answer, then "1,-3,2", "-3,2,1" or any permutation of the same numbers may not appear in the same answer (though any one of them may appear instead of "1,2,-3").
// Constraints:

// 1 <= n <= 2000
// -1000 <= any element of arr <= 1000
// arr may contain duplicate numbers
// arr is not necessarily sorted
package threeSum

import (
	"sort"
	"fmt"
)

func find_zero_sum(arr []int) []string {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	resultString := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		j, k := i+1, len(arr)-1
		for j < k {
			sum := arr[i] + arr[j] + arr[k]
			if sum == 0 {
				resultString = append(resultString, fmt.Sprintf("%d,%d,%d", arr[i], arr[j], arr[k]))
				j++
				k--
				for j < k && arr[j] == arr[j-1] {
					j++
				}
				for j < k && arr[k] == arr[k+1] {
					k--
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return resultString
}
